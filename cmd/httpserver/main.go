package main

import (
	"context"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/gibalmeida/mailservermngr/internal/adapter/http/api"
	"github.com/gibalmeida/mailservermngr/internal/adapter/repository"
	"github.com/gibalmeida/mailservermngr/internal/app/usecase"

	"github.com/gibalmeida/mailservermngr/internal/app/config"
	"github.com/gibalmeida/mailservermngr/internal/domain"
	"github.com/gibalmeida/mailservermngr/pkg/jwx"
)

func main() {
	cfg := config.LoadConfig()

	authRepo := repository.NewMemAuthRepository()
	// While we are not using a persistent store, we need to create a user in the memory store.
	// For this we obtain the username and password provided in the flag parameters and create this user in the repository.
	err := authRepo.CreateUser(context.TODO(), domain.User{Username: cfg.AdminUsername, Password: cfg.AdminPassword})
	if err != nil {
		panic("Cannot create admin user")
	}

	mailServerRepo, err := repository.NewMysqlMailServerRepository(cfg.DatabaseURI)
	// repo := repository.NewMemMailServerRepository()

	if err != nil {
		log.Fatalln("error instatiating the repository:", err)
	}
	// Descomentar a linha abaixo para um banco de dados MySQL
	defer mailServerRepo.Close()

	mailServerUseCase := usecase.NewMailServerUseCase(mailServerRepo)

	// This is how you set up a basic Echo router
	e := echo.New()

	// Create a JWS authenticator. This allows us to issue tokens, and also
	// implements a validator to check their validity.
	jwsAuth, err := jwx.NewJWSAuthenticator(cfg.PrivateKey)
	if err != nil {
		log.Fatalln("error creating authenticator:", err)
	}

	// Create middleware for validating tokens.
	mw, err := api.CreateMiddleware(jwsAuth)
	if err != nil {
		log.Fatalln("error creating middleware:", err)
	}

	// Log all requests
	e.Use(echomiddleware.Logger())

	e.Use(mw...)

	authUseCase := usecase.NewAuthUseCase(authRepo, jwsAuth)

	// Create an instance of our handler which satisfies the generated interface
	mailServerServer := api.NewServer(authUseCase, mailServerUseCase)

	// We now register our mailServer above as the handler for the interface
	api.RegisterHandlersWithBaseURL(e, mailServerServer, "/api/v1")

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", cfg.HttpPort)))
}

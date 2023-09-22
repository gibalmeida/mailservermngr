package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gibalmeida/mailservermngr/pkg/env"
	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort      string
	DatabaseURI   string
	AdminUsername string
	AdminPassword string
	PrivateKey    []byte
}

func LoadConfig() *Config {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	dbURI := fmt.Sprintf("%s:%s@(%s:%d)/%s",
		env.GetAsString("DB_USER", "user"),
		env.GetAsString("DB_PASSWORD", "password"),
		env.GetAsString("DB_HOST", "localhost"),
		env.GetAsInt("DB_PORT", 3306),
		env.GetAsString("DB_NAME", "dbname"),
	)

	port := flag.String("port", "8080", "Port for test HTTP server")

	adminUsername := flag.String("username", "admin", "Admin Username")
	adminPassword := flag.String("password", "admin", "Admin Password")

	jwtPrivateKeyFilePath := flag.String("privkey", "./ecprivatekey.pem", "Path to the file containing the ECDSA private key (PEM format). You can generate the key file using the following command:\nopenssl ecparam -name prime256v1 -genkey -noout -out ecprivatekey.pem\n")

	flag.Parse()

	privKey, err := os.ReadFile(*jwtPrivateKeyFilePath)
	if err != nil {
		log.Fatalln("error reading the private key file:", err)
	}

	return &Config{
		HttpPort:      *port,
		DatabaseURI:   dbURI,
		AdminUsername: *adminUsername,
		AdminPassword: *adminPassword,
		PrivateKey:    privKey,
	}
}

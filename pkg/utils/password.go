package utils

import (
	"github.com/GehirnInc/crypt/md5_crypt"
)

func HashPassword(clearTextPassword string) (string, error) {

	crypter := md5_crypt.New()
	salt := []byte{}

	return crypter.Generate([]byte(clearTextPassword), salt)

}

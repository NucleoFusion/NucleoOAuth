package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return string(hashed), err
	}

	return string(hashed), nil
}

package usecases

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password []byte) ([]byte, error) {
	newPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return newPassword, nil
}

package security

import "golang.org/x/crypto/bcrypt"

func Hash(key string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
}

func CheckPassword(hash, key string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(key))
}

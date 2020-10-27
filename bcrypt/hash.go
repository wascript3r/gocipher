package bcrypt

import "golang.org/x/crypto/bcrypt"

func Compute(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func Validate(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

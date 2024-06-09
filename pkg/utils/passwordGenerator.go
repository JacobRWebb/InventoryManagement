package utils

import "golang.org/x/crypto/bcrypt"

func NormalizePassword(password string) []byte {
	return []byte(password)
}

func GeneratePasswordHash(password string) string {
	normalizedPassword := NormalizePassword(password)

	hash, err := bcrypt.GenerateFromPassword(normalizedPassword, bcrypt.DefaultCost)

	if err != nil {
		return err.Error()
	}

	return string(hash)
}

func ComparePasswords(hashedPassword string, password string) bool {
	// TODO make more secure - Include salts?
	byteHash := NormalizePassword(hashedPassword)
	bytePassword := NormalizePassword(password)

	if err := bcrypt.CompareHashAndPassword(byteHash, bytePassword); err != nil {
		return false
	}

	return true
}

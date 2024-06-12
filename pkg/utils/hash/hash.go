package hash

type PasswordHash interface {
	ComparePasswordAndHash(password string, encodedHash string) (bool, error)
	GenerateFromPassword(password string) (string, error)
}

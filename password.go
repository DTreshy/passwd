package passwd

import "golang.org/x/crypto/bcrypt"

// Hash takes password and returnes hash based on it
func Hash(pwd string) (string, error) {
	return HashWithCost(pwd, 12)
}

// HashWithCost takes password and returnes hash based on it with cost as provided
func HashWithCost(pwd string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	return string(bytes), err
}

// CheckPasswordHash takes password and hash. Returns true if it matches, flase otherwise
func Check(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

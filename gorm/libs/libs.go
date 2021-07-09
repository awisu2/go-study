package libs

import (
	"crypto/rand"
)

const RANDOM_LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func CreateRandomString(digit uint32, letters string) (string, error) {
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	s, length := "", len(letters)
	for _, v := range b {
		s += string(letters[int(v) % length])
	}

	return s, nil
}


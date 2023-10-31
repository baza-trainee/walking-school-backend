package service

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(password, salt string) string {
	sum := sha256.Sum256([]byte(password + salt))

	return fmt.Sprintf("%x", sum)
}

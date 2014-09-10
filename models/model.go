package models

import (
	"crypto/sha512"
	"fmt"
	"io"
)

type Model interface {
	All() interface{}
	FindById() interface{}
	Create(params map[string]string) interface{}
}

func passwordHash(password string) string {
	passwordHash := sha512.New()
	io.WriteString(passwordHash, password)

	return fmt.Sprintf("%x", passwordHash.Sum(nil))
}

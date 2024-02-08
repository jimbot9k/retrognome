package utils

import (
	"crypto/sha512"
	"math/rand"
)

func HashPassword(password string, salt string) string {
	saltedPassword := password + salt
	hash := sha512.New()
	hash.Write([]byte(saltedPassword))
	return string(hash.Sum(nil))
}

func RandomString(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

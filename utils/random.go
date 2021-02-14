package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

//RandomCode function for make random hex code
func RandomCode(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

//GenerateCode helper for generate code
func GenerateCode(length int) string {
	codeverify := ""
	for i := 0; i < length; i++ {
		random, _ := RandomCode(2)
		codeverify = fmt.Sprintf("%s%v", codeverify, random)
	}
	return codeverify
}

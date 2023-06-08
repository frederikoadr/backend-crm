package hashing

import (
	"crypto/sha256"
	"fmt"
)

func GenerateHash(data string) string {
	// Membuat objek hashing dari algoritma SHA-256
	hash := sha256.New()

	// Mengupdate hashing dengan data yang ingin di-hashing
	hash.Write([]byte(data))

	// Mengambil nilai hashing sebagai array byte
	hashBytes := hash.Sum(nil)

	// Mengubah array byte menjadi representasi heksadesimal
	hashString := fmt.Sprintf("%x", hashBytes)

	return hashString
}

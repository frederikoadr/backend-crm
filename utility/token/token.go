package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateJWT(id, secretKey string) (string, error) {
	// Inisialisasi klaim-klaim yang ingin Anda sertakan dalam token
	claims := jwt.MapClaims{
		"sub": id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	// Tandatangani token dengan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		// Penanganan kesalahan
		return "", err
	}
	return signedToken, nil
}
func VerfiyJWT(receivedToken, secret string) (string, error) {

	// Verifikasi token dengan kunci rahasia
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		// Penanganan kesalahan
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Token valid, akses klaim-klaim yang ada
		return fmt.Sprintf("%s", claims["sub"]), nil
	} else {
		// Token tidak valid
	}
	return "", errors.New("Invalid token")
}

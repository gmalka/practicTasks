package authrepository

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

type Auth struct {
	secret []byte
	key    []byte
}

func NewAuth(secret, key []byte) Auth {
	return Auth{secret: secret, key: key}
}

func (a Auth) NewCryptedMessage(message []byte) (string, error) {
	aesBlock, err := aes.NewCipher(a.key)
	if err != nil {
		return "", nil
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return "", nil
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, message, bytes)

	return base64.StdEncoding.EncodeToString(cipheredText), nil
}

func (a Auth) EncryptedMessage(message string) (string, error) {
	ciphered, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", nil
	}
	aesBlock, err := aes.NewCipher(a.key)
	if err != nil {
		return "", nil
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return "", nil
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, bytes)
	if err != nil {
		return "", nil
	}
	return string(originalText), nil
}

func (a Auth) NewToken(username, tokenType string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["type"] = tokenType
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, _ := token.SignedString(a.secret)

	return tokenString
}

func (a Auth) CheckToken(token, wantedType string) (AuthData, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return a.secret, nil
	})
	if err != nil {
		return AuthData{}, err
	}
	if t == nil || t.Valid == false {
		return AuthData{}, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return AuthData{}, errors.New("invalid token")
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return AuthData{}, errors.New("the token has already expired")
	}

	return AuthData{
		Username: claims["username"].(string),
		TokenType: claims["type"].(string),
	}, nil
}

type AuthData struct {
	Username string		`json:"username"`
	TokenType string	`json:"type"`
}
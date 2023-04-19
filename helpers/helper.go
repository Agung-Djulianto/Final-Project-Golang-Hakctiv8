package helpers

import (
	"Project-Akhir/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateID() string {
	id := uuid.New()
	return id.String()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CekPasswordHash(password, Hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(Hash), []byte(password))

	return err == nil
}

func GenerateToken(userID string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte("AGung___89898"))

	return tokenString, err
}

func VerifyToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, models.ErrorInvalidToken
		}
		return []byte("AGung___89898"), nil
	})
	return jwtToken, err
}

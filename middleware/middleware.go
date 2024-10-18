package middleware

import (
	"os"
	"time"

	"go-authorization/model/entity"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func GenerateToken(user entity.MasterUserModel, roleName string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.MapClaims{
		"id":              user.ID,
		"role":            roleName,
		"working_area_id": user.WorkingAreaID,
		"isarea":          user.Isarea,
		"exp":             expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

package jwtToken

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sunzhqr/sso/internal/domain/models"
	"time"
)

func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["app_id"] = app.ID
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

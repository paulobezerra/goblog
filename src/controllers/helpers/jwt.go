package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/paulobezerra/goblog/src/models"
	"github.com/paulobezerra/goblog/src/utils"
)

var JWTHash []byte = []byte("ec5dad9f155810adba7c300a3270094a88bf04b0")

func GenerateJWT(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})

	tokenString, err := token.SignedString(JWTHash)

	utils.CheckErr(err)

	return tokenString
}

func GetUserByJWT(hash string) (*models.User, error) {
	token, err := jwt.Parse(hash, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTHash, nil
	})

	var user models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userJson, _ := json.Marshal(claims["user"])
		json.Unmarshal(userJson, &user)
		dbUser := models.GetUser(strconv.Itoa(user.Id))
		return dbUser, nil
	}

	return nil, err
}

package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"server.com/pkg/types"
	authTypes "server.com/pkg/types/auth"
	"server.com/pkg/utils/logger"
)

type GenerateTokenPayload struct {
	Username   string            `json:"username"`
	UserID     uint              `json:"userId"`
	Exp        int64             `json:"exp"`
	Permisions []types.Permision `json:"pemisions"`
	GrantType  string            `json:"grantType"`
}

func GenerateToken(payload GenerateTokenPayload) (string, error) {
	ttlInMns, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_IN_MINUTES"))
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	timeExpire := time.Duration(ttlInMns)
	permisions := make(map[string]bool)
	for _, ps := range payload.Permisions {
		permisions[string(ps)] = true
	}

	// token age
	age := time.Now().Add(time.Minute * timeExpire).Unix()
	if payload.GrantType == "refresh-token" {
		age = age * 2
	}

	// make token paload
	claims := jwt.MapClaims{
		"username": payload.Username,
		"userId":   payload.UserID,
		// "exp":        age,
		"permisions": permisions,
		"grantType":  payload.GrantType,
	}

	// Set private token credentials:
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenStr, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		logger.Log(err.Error())
		return tokenStr, err
	}

	return tokenStr, nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*authTypes.TokenMetadata, error) {
	// TryExample()
	data := NewTryCatch(
		func() (*authTypes.TokenMetadata, error) {
			// logger.Info("==== new try")
			token := c.Locals("jwt").(*jwt.Token)
			// Setting and checking token and credentials.
			claims, ok := token.Claims.(jwt.MapClaims)
			// logger.Info("==== res token token", token, " => ", token.Valid, " => ", ok)
			if ok && token.Valid {

				// get all details.
				userID := uint(claims["userId"].(float64))
				username := claims["username"].(string)
				grantType := claims["grantType"].(string)
				roleID := uint(claims["roleId"].(float64))
				// exp := int64(claims["exp"].(float64))

				// check token type
				path := string(c.Request().URI().PathOriginal())
				// logger.Debug(path)
				// logger.Debug(grantType)
				if (path != "/api/token/refresh_token" && grantType != "token") || (path == "/api/token/refresh_token" && grantType != "refresh-token") {
					logger.Debug(">>>> grattype wrong")
					return nil, errors.New("invalid grant type")
				}

				// permisions := make(map[types.Permision]bool)
				// for key, value := range claims["permisions"].(map[string]interface{}) {
				// 	permisions[types.Permision(key)] = value.(bool)
				// }

				tokenMetaData := &authTypes.TokenMetadata{
					UserID:   userID,
					Username: username,
					// Exp:        exp,
					// Permisions: permisions,
					GrantType: grantType,
					RoleID:    roleID,
				}
				// logger.Info("==== res token")
				return tokenMetaData, nil
			}
			return nil, errors.New("invalid token")
		}).Catch(func(try *TryCatch[*authTypes.TokenMetadata]) {

		// do sth with error
		// logger.Log(try.panicErr)
	}).data
	if data == nil {
		logger.Info("==== res token nil")
		return nil, errors.New("invalid token")
	}
	//logger.Log(number, panicErr, err)
	return data, nil
}

package middlewares

import (
	"errors"
	"go-task-app/internal/config"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("不正なトークンです。")

func decodeJWT(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println("署名方法が違います。")
			return nil, ErrInvalidToken
		}

		return []byte(config.AuthSecret), nil
	})

	if err != nil {
		log.Println(err)
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &claims, nil
	}

	return nil, ErrInvalidToken
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "アクセストークンが付与されていません。"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := decodeJWT(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		userId := int((*claims)["userId"].(float64))
		expiresInUnix := int64((*claims)["expiresIn"].(float64))
		expiresIn := time.Unix(expiresInUnix, 0)

		expired := time.Now().After(expiresIn)

		if expired {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": errors.New("アクセストークンの期限が切れています。"),
			})
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}

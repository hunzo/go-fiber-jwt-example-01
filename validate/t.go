package validate

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const JWT_SECRET = "secret"

func GetAccessToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		Audience:  "test_user",
	})

	ss, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return "", err
	}

	return ss, nil
}

func Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if !strings.Contains(auth, "Bearer ") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
		}

		if auth == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpect signing method: %v", t.Header["alg"])
			}
			return []byte(JWT_SECRET), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			aud := claims["aud"]
			c.Locals("aud", aud)
		}

		fmt.Println(tokenString)
		return c.Next()
	}
}

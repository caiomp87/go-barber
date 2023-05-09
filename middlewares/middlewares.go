package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	interceptorServices := getInterceptorService()
	interceptorValue := formatInterceptorValues(c)

	fmt.Println(interceptorValue)

	if !isPublic(interceptorServices, interceptorValue) {
		token := c.Get("authorization")
		if token == "" {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "token not provided",
			})
			return nil
		}
	}

	c.Next()
	return nil
}

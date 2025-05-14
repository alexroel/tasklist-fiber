package middlewares

import (
	"net/http"
	"tasklist/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Middleware para proteger rutas
func Protected(c *fiber.Ctx) error {
	// Obtener el token del encabezado Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token no proporcionado",
		})
	}

	// Parsear el token
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		// Validar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(http.StatusUnauthorized, "Método de firma inválido")
		}
		return config.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token inválido",
		})
	}

	// Obtener el userID de los claims y guardarlo en el contexto
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userID, ok := claims["user_id"].(string); ok {
			c.Locals("userID", userID)
		} else if userIDFloat, ok := claims["user_id"].(float64); ok {
			// Si el user_id es numérico
			c.Locals("userID", int(userIDFloat))
		} else {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "user_id no encontrado en el token",
			})
		}
	} else {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Claims inválidos",
		})
	}

	// Continuar con la solicitud
	return c.Next()
}
package main

import (
	"tasklist/config"
	"tasklist/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	// Configurar las rutas
	routers.SetupRoutes(app)

	// Iniciar el servidor
	app.Listen(":3000")
}
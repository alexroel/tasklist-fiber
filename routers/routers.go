package routers

import (
	"tasklist/handlers"
	"tasklist/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Rutas de autenticación
	authRoutes := app.Group("/auth")
	authRoutes.Post("/login", handlers.Login)
	authRoutes.Post("/register", handlers.RegisterUser)
	authRoutes.Post("/logout", handlers.Logout)

	// Rutas protegidas de usuarios
	userRoutes := app.Group("/users", middlewares.Protected)
	userRoutes.Get("/", handlers.GetUsers)
	userRoutes.Get("/:id", handlers.GetUser)
	userRoutes.Post("/", handlers.CreateUser)
	userRoutes.Put("/:id", handlers.UpdateUser)
	userRoutes.Delete("/:id", handlers.DeleteUser)

	// Rutas protegidas de tareas
	taskRoutes := app.Group("/tasks", middlewares.Protected)
	taskRoutes.Get("/", handlers.GetTasks)
	taskRoutes.Get("/:id", handlers.GetTask)
	taskRoutes.Post("/", handlers.CreateTask)
	taskRoutes.Put("/:id", handlers.UpdateTask)
	taskRoutes.Delete("/:id", handlers.DeleteTask)
}
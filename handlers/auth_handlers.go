package handlers

import (
	"net/http"
	"tasklist/config"
	"tasklist/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Buscar al usuario por correo electrónico
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Login maneja el inicio de sesión de los usuarios
func Login(c *fiber.Ctx) error {
	// Definir la estructura de entrada
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parsear el cuerpo de la solicitud
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	// Buscar al usuario por correo electrónico
	user, err := GetUserByEmail(input.Email)
	if err != nil {	
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email o contraseña incorrectos",
		})
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email o contraseña incorrectos",
		})
	}

	// Generar el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	})

	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo generar el token",
		})
	}

	// Respuesta exitosa con el token
	return c.JSON(fiber.Map{
		"message": "Inicio de sesión exitoso",
		"token":   tokenString,
	})
}

// RegisterHandler maneja el registro de nuevos usuarios
func RegisterUser(c *fiber.Ctx) error {
	// Definir la estructura de entrada
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parsear el cuerpo de la solicitud
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	// Verificar si el correo ya está registrado
	existingUser, err := GetUserByEmail(input.Email)
	if err == nil && existingUser != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"error": "El correo electrónico ya está registrado",
		})
	}

	// Encriptar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al encriptar la contraseña",
		})
	}

	// Crear el nuevo usuario
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al registrar el usuario",
		})
	}

	// Respuesta exitosa
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Usuario registrado exitosamente",
		"user":    user,
	})
}

// LogoutHandler maneja el cierre de sesión de los usuarios
func Logout(c *fiber.Ctx) error {
	// Elimina el token del cliente (si se usa en cookies)
	c.ClearCookie("Authorization")

	// Respuesta exitosa
	return c.JSON(fiber.Map{
		"message": "Cierre de sesión exitoso",
	})
}
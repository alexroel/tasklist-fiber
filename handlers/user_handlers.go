package handlers

import (
	"net/http"
	"tasklist/config"
	"tasklist/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Obtener todos los usuarios
func GetUsers(c *fiber.Ctx) error {
    var users []models.User
    result := config.DB.Find(&users)
    if result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al obtener los usuarios",
        })
    }
    return c.JSON(users)
}

// Obtener un usuario por ID
func GetUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    result := config.DB.First(&user, id)
    if result.Error != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Usuario no encontrado",
        })
    }
    return c.JSON(user)
}

// Crear un nuevo usuario
func CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Datos inválidos",
        })
    }

    // Encripta la contraseña
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al encriptar la contraseña",
        })
    }
    user.Password = string(hashedPassword)

    result := config.DB.Create(&user)
    if result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al crear el usuario",
        })
    }
    return c.Status(http.StatusCreated).JSON(user)
}

// Actualizar un usuario
func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    result := config.DB.First(&user, id)
    if result.Error != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Usuario no encontrado",
        })
    }

    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Datos inválidos",
        })
    }

    config.DB.Save(&user)
    return c.JSON(user)
}

// Eliminar un usuario
func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    result := config.DB.First(&user, id)
    if result.Error != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Usuario no encontrado",
        })
    }

    config.DB.Delete(&user)
    return c.Status(http.StatusNoContent).Send(nil)
}
package handlers

import (
	"tasklist/config"
	"tasklist/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Buscar la tarea por ID
func GetTaskByID(id string) (*models.Task, error) {
	var task models.Task
	result := config.DB.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

// GetTasks recupera las tareas de la base de datos solo del suario autenticado
func GetTasks(c *fiber.Ctx) error {
    userID := c.Locals("userID")
    if userID == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Usuario no autenticado"})
    }

    var tasks []models.Task
    if err := config.DB.Preload("User").Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve tasks"})
    }

    return c.JSON(tasks)
}

// GetTask retrieves a single task by ID
func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	
	task, err := GetTaskByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarea no encontrada"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al recuperar la tarea"})
	}

	return c.JSON(task)

}

// CreateTask creates a new task in the database
func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := config.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}

// UpdateTask updates an existing task in the database
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")

	// Buscar la tarea por ID
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve task"})
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := config.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}
	return c.JSON(task)
}

// DeleteTask deletes a task by ID from the database
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Task{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete task"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

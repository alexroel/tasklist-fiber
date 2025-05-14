package config

import (
	"log"
	"os"
	"tasklist/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    // Cargar variables de entorno
    if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
	

    // Construir la cadena de conexión
    dsn := os.Getenv("DB_DSN")

    // Conectar a la base de datos
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error conectando a la base de datos: %v", err)
    }

    // Migrar los modelos
    err = database.AutoMigrate(&models.User{}, &models.Task{})
    if err != nil {
        log.Fatalf("Error en la migración de modelos: %v", err)
    }

    DB = database
    log.Println("Conexión a la base de datos exitosa")
}
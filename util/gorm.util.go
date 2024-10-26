package util

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB(db *gorm.DB) (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	fmt.Println(dsn)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error obteniendo la conexión de base de datos: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)  // Máximo número de conexiones abiertas
	sqlDB.SetMaxIdleConns(10)   // Máximo número de conexiones inactivas
	sqlDB.SetConnMaxLifetime(0) // Tiempo de vida máximo de una conexión

	fmt.Println("Conexión a la base de datos exitosa")

	return db, nil
}

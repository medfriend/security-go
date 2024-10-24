package util

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(db *gorm.DB) {

	dsn := "host=localhost user=elliot password=H0sp1T4l dbname=med-friend port=7500 sslmode=disable TimeZone=UTC"

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
}

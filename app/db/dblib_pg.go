package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbPuerto := os.Getenv("DB_PORT")
	dbUsuario := os.Getenv("DB_USER")
	dbPw := os.Getenv("DB_PASSWORD")
	dbNombre := os.Getenv("DB_NAME")

	fmt.Println("Conectando con " + dbHost + ":" + dbPuerto)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUsuario, dbPw, dbHost, dbPuerto, dbNombre)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error en conexion a la base de datos")
		panic(err)
	} else {
		fmt.Println("Conexion con base de datos exitosa!")
	}
}

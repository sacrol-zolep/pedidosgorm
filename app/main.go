package main

import (
	"app_pedidos/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dataBase *gorm.DB

func main() {
	loadEnv()
	dbHost := os.Getenv("DB_HOST")
	dbPuerto := os.Getenv("DB_PORT")
	dbUsuario := os.Getenv("DB_USER")
	dbPw := os.Getenv("DB_PASSWORD")
	dbNombre := os.Getenv("DB_NAME")

	fmt.Println("Conectando con " + dbHost + ":" + dbPuerto)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUsuario, dbPw, dbHost, dbPuerto, dbNombre)
	dataBase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error en conexion a la base de datos")
		panic(err)
	} else {
		fmt.Println("Concexion con base de datos exitosa!")
	}

	dataBase.AutoMigrate(&models.Cliente{})
	dataBase.AutoMigrate(&models.Contacto{})
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error encontrando .env file")
	}
}

package main

import (
	"app_pedidos/db"
	"app_pedidos/models"
	"context"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	loadEnv()
	db.Connect()
	db.Database.AutoMigrate(&models.Cliente{})
	db.Database.AutoMigrate(&models.Contacto{})
	NuevoCliente()
	NuevoContacto()
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error encontrando .env file")
	}
}

func NuevoCliente() {
	cliente := models.Cliente{
		Nombre1:       "Juan",
		Apellido1:     "Perez",
		TipoDocumento: "DNI",
		Documento:     "2012",
	}

	ctx := context.Background()
	err := gorm.G[models.Cliente](db.Database).Create(ctx, &cliente)

	if err != nil {
		panic(err)
	} else {
		log.Println("Cliente creado correctamente")
	}

}

func NuevoContacto() {
	contacto := models.Contacto{
		TipoDato:   "correo",
		Dato:       "jaun@test.com",
		ClienteDoc: "2012",
	}

	ctx := context.Background()
	err := gorm.G[models.Contacto](db.Database).Create(ctx, &contacto)

	if err != nil {
		panic(err)
	} else {
		log.Println("Contacto creado correctamente")
	}

}

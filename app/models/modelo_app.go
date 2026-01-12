package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nombre1       string    // Cadena obligatoria
	Nombre2       *string   // Cadena que puede ser vacia
	Apellido1     string    // Cadena obligatoria
	Apellido2     *string   // Cadena que puede ser vacia
	TipoDocumento string    // Cadena obligatoria
	CreatedAt     time.Time // Gestionada automaticamente por GORM al crear registro
	UpdatedAt     time.Time // Gestionada automaticamente por GORM al actualizar registro
	Enabled       bool      // Indicador de registro activo
	Contactos     []Contacto
}

type Contacto struct {
	gorm.Model
	TipoDato  string    // Tipo de registro, puede ser correo, telefono, etc
	Dato      string    // El valor registrado
	CreatedAt time.Time // Gestionada automaticamente por GORM al crear registro
	UpdatedAt time.Time // Gestionada automaticamente por GORM al actualizar registro
	Enabled   bool      // Indicador de registro activo
	ClienteID uint      //llave foranea IDCliente
}

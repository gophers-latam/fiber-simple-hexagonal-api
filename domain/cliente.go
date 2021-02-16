package domain

import (
	"github.com/zeroidentidad/fiber-hex-api/dto"
	"github.com/zeroidentidad/fiber-hex-api/errors"
)

type Cliente struct {
	ID              string `db:"cliente_id"`
	Nombre          string
	Ciudad          string
	CodigoPostal    string `db:"codigo_postal"`
	FechaNacimiento string `db:"fecha_nacimiento"`
	Estatus         string
}

func (c Cliente) estatusAsText() string {
	estatusAsText := "active"
	if c.Estatus == "0" {
		estatusAsText = "inactive"
	}
	return estatusAsText
}

func (c Cliente) ToDtoResponse() dto.ResponseCliente {
	return dto.ResponseCliente{
		ID:              c.ID,
		Nombre:          c.Nombre,
		Ciudad:          c.Ciudad,
		CodigoPostal:    c.CodigoPostal,
		FechaNacimiento: c.FechaNacimiento,
		Estatus:         c.estatusAsText(),
	}
}

type StorageCliente interface {
	FindAll(string) ([]Cliente, *errors.AppError)
	ById(string) (*Cliente, *errors.AppError)
}

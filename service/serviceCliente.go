package service

import (
	"github.com/zeroidentidad/fiber-hex-api/domain"
	"github.com/zeroidentidad/fiber-hex-api/dto"
	"github.com/zeroidentidad/fiber-hex-api/errors"
)

type ServiceCliente interface {
	GetAll(string) ([]dto.ResponseCliente, *errors.AppError)
	GetById(string) (*dto.ResponseCliente, *errors.AppError)
}

type DefaultServiceCliente struct {
	repo domain.StorageCliente
}

func (s DefaultServiceCliente) GetAll(estatus string) ([]dto.ResponseCliente, *errors.AppError) {
	if estatus != "" {
		estatus = ternary(estatus == "active", "1", "0").(string)
	} else {
		estatus = ""
	}

	clientes, err := s.repo.FindAll(estatus)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ResponseCliente, 0)
	for _, c := range clientes {
		res = append(res, c.ToDtoResponse())
	}

	return res, err
}

func (s DefaultServiceCliente) GetById(id string) (*dto.ResponseCliente, *errors.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	res := dto.ResponseCliente{
		ID:              c.ID,
		Nombre:          c.Nombre,
		Ciudad:          c.Ciudad,
		CodigoPostal:    c.CodigoPostal,
		FechaNacimiento: c.FechaNacimiento,
		Estatus:         c.Estatus,
	}

	return &res, nil
}

func NewServiceCliente(repo domain.StorageCliente) DefaultServiceCliente {
	return DefaultServiceCliente{
		repo,
	}
}

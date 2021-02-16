package service

import (
	"github.com/zeroidentidad/fiber-hex-api/domain"
	"github.com/zeroidentidad/fiber-hex-api/dto"
	"github.com/zeroidentidad/fiber-hex-api/errors"
)

type ServiceCuenta interface {
	PostNew(dto.RequestCuenta) (*dto.ResponseCuenta, *errors.AppError)
	PostNewTransaccion(dto.RequestTransaccion) (*dto.ResponseTransaccion, *errors.AppError)
}

type DefaultServiceCuenta struct {
	repo domain.StorageCuenta
}

func (s DefaultServiceCuenta) PostNew(req dto.RequestCuenta) (*dto.ResponseCuenta, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	c := domain.NewCuenta(req.ClienteID, req.TipoCuenta, req.Cantidad)

	cuenta, err := s.repo.Save(c)
	if err != nil {
		return nil, err
	}
	response := cuenta.ToDtoResponse()

	return &response, nil
}

func (s DefaultServiceCuenta) PostNewTransaccion(req dto.RequestTransaccion) (*dto.ResponseTransaccion, *errors.AppError) {
	// validación peticion entrante
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// validación del servidor para verificar el saldo disponible en la cuenta
	if req.EsTipoTransaccionRetiro() {
		cuenta, err := s.repo.FindBy(req.CuentaID)
		if err != nil {
			return nil, err
		}
		if !cuenta.PuedeRetirar(req.Cantidad) {
			return nil, errors.NewValidationError("Saldo insuficiente en la cuenta")
		}
	}
	// si todo bien, hacer objeto de dominio y guardar transacción
	t := domain.NewTransaccion(req.CuentaID, req.TipoTransaccion, req.Cantidad)

	transaccion, appError := s.repo.SaveTransaccion(t)
	if appError != nil {
		return nil, appError
	}

	response := transaccion.ToDtoResponse()

	return &response, nil
}

func NewServiceCuenta(repo domain.StorageCuenta) DefaultServiceCuenta {
	return DefaultServiceCuenta{
		repo,
	}
}

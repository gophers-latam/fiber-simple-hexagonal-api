package dto

import (
	"strings"

	"github.com/zeroidentidad/fiber-hex-api/errors"
)

type RequestCuenta struct {
	ClienteID  string  `json:"cliente_id"`
	TipoCuenta string  `json:"tipo_cuenta"`
	Cantidad   float64 `json:"cantidad"`
}

func (r RequestCuenta) Validate() *errors.AppError {
	if r.Cantidad < 5000 {
		return errors.NewValidationError("Para abrir una nueva cuenta, debe depositar al menos 5000.00")
	}
	if strings.ToLower(r.TipoCuenta) != "ahorro" && strings.ToLower(r.TipoCuenta) != "chequera" {
		return errors.NewValidationError("El tipo de cuenta debe ser de cheques o de ahorro")
	}
	return nil
}

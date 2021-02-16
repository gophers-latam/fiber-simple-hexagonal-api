package dto

import "github.com/zeroidentidad/fiber-hex-api/errors"

const RETIRO = "retiro"
const DEPOSITO = "deposito"

type RequestTransaccion struct {
	CuentaID         string  `json:"cuenta_id"`
	Cantidad         float64 `json:"cantidad"`
	TipoTransaccion  string  `json:"tipo_transaccion"`
	FechaTransaccion string  `json:"fecha_transaccion"`
	ClienteID        string  `json:"-"`
}

func (r RequestTransaccion) EsTipoTransaccionRetiro() bool {
	return r.TipoTransaccion == RETIRO
}

func (r RequestTransaccion) EsTipoTransaccionDeposito() bool {
	return r.TipoTransaccion == DEPOSITO
}

func (r RequestTransaccion) Validate() *errors.AppError {
	if !r.EsTipoTransaccionRetiro() && !r.EsTipoTransaccionDeposito() {
		return errors.NewValidationError("El tipo de transacción solo puede ser depósito o retiro")
	}
	if r.Cantidad < 0 {
		return errors.NewValidationError("La cantidad no puede ser menor a cero")
	}
	return nil
}

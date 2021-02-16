package domain

import (
	"time"

	"github.com/zeroidentidad/fiber-hex-api/dto"
	"github.com/zeroidentidad/fiber-hex-api/errors"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Cuenta struct {
	ID            string `db:"cuenta_id"`
	ClienteID     string `db:"cliente_id"`
	FechaApertura string `db:"fecha_apertura"`
	TipoCuenta    string `db:"tipo_cuenta"`
	Cantidad      float64
	Estatus       string
}

func (c Cuenta) ToDtoResponse() dto.ResponseCuenta {
	return dto.ResponseCuenta{ID: c.ID}
}

type StorageCuenta interface {
	Save(Cuenta) (*Cuenta, *errors.AppError)
	SaveTransaccion(Transaccion) (*Transaccion, *errors.AppError)
	FindBy(string) (*Cuenta, *errors.AppError)
}

func (c Cuenta) PuedeRetirar(cantidad float64) bool {
	if c.Cantidad < cantidad {
		return false
	}
	return true
}

func NewCuenta(clienteId, tipoCuenta string, cantidad float64) Cuenta {
	return Cuenta{
		ClienteID:     clienteId,
		FechaApertura: time.Now().Format(dbTSLayout),
		TipoCuenta:    tipoCuenta,
		Cantidad:      cantidad,
		Estatus:       "1",
	}
}

func NewTransaccion(cuentaId, tipoTransaccion string, cantidad float64) Transaccion {
	return Transaccion{
		CuentaID:         cuentaId,
		Cantidad:         cantidad,
		TipoTransaccion:  tipoTransaccion,
		FechaTransaccion: time.Now().Format(dbTSLayout),
	}
}

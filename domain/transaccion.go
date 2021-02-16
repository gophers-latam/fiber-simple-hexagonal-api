package domain

import "github.com/zeroidentidad/fiber-hex-api/dto"

const RETIRO = "retiro"

type Transaccion struct {
	ID               string  `db:"transaccion_id"`
	CuentaID         string  `db:"cuenta_id"`
	Cantidad         float64 `db:"cantidad"`
	TipoTransaccion  string  `db:"tipo_transaccion"`
	FechaTransaccion string  `db:"fecha_transaccion"`
}

func (t Transaccion) EsRetiro() bool {
	if t.TipoTransaccion == RETIRO {
		return true
	}
	return false
}

func (t Transaccion) ToDtoResponse() dto.ResponseTransaccion {
	return dto.ResponseTransaccion{
		ID:               t.ID,
		CuentaID:         t.CuentaID,
		Cantidad:         t.Cantidad,
		TipoTransaccion:  t.TipoTransaccion,
		FechaTransaccion: t.FechaTransaccion,
	}
}

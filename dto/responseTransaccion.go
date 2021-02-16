package dto

type ResponseTransaccion struct {
	ID               string  `json:"transaccion_id"`
	CuentaID         string  `json:"cuenta_id"`
	Cantidad         float64 `json:"nuevo_balance"`
	TipoTransaccion  string  `json:"tipo_transaccion"`
	FechaTransaccion string  `json:"fecha_transaccion"`
}

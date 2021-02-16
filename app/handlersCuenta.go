package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeroidentidad/fiber-hex-api/dto"
	"github.com/zeroidentidad/fiber-hex-api/service"
)

type HandlersCuenta struct {
	service service.ServiceCuenta
}

func (hc *HandlersCuenta) postNewCuenta(c *fiber.Ctx) error {
	cliente_id := c.Params("id")

	body := new(dto.RequestCuenta)
	if err := ParseBody(c, body); err != nil {
		return err
	}

	body.ClienteID = cliente_id
	cuenta, err := hc.service.PostNew(*body)
	return resJSON(cuenta, err, c)
}

func (hc *HandlersCuenta) postNewTransaccion(c *fiber.Ctx) error {
	cliente_id := c.Params("id")
	cuenta_id := c.Params("id_cuenta")

	body := new(dto.RequestTransaccion)
	if err := ParseBody(c, body); err != nil {
		return err
	}

	body.ClienteID = cliente_id
	body.CuentaID = cuenta_id
	transaccion, err := hc.service.PostNewTransaccion(*body)
	return resJSON(transaccion, err, c)
}

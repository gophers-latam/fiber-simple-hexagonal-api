package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zeroidentidad/fiber-hex-api/service"
)

type HandlersCliente struct {
	service service.ServiceCliente
}

func (hc *HandlersCliente) getAllClientes(c *fiber.Ctx) error {
	estatus := c.Query("estatus")
	clientes, err := hc.service.GetAll(estatus)
	return resJSON(clientes, err, c)
}

func (hc *HandlersCliente) getCliente(c *fiber.Ctx) error {
	id := c.Params("id")
	cliente, err := hc.service.GetById(id)
	return resJSON(cliente, err, c)
}

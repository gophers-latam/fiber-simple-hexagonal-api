package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zeroidentidad/fiber-hex-api/domain"
	"github.com/zeroidentidad/fiber-hex-api/service"
)

func Start() {
	configLoad()
	router := fiber.New()
	router.Use(logger.New())

	dbClient := dbClient()
	storageDbCliente := domain.NewStorageDbCliente(dbClient)
	storageDbCuenta := domain.NewStorageDbCuenta(dbClient)

	hclientes := HandlersCliente{service.NewServiceCliente(storageDbCliente)}
	hcuentas := HandlersCuenta{service.NewServiceCuenta(storageDbCuenta)}

	router.Get("/clientes", hclientes.getAllClientes)
	router.Get("/clientes/:id", hclientes.getCliente)
	router.Post("/clientes/:id/cuenta", hcuentas.postNewCuenta)
	router.Post("/clientes/:id/cuenta/:id_cuenta", hcuentas.postNewTransaccion)

	ADDR := os.Getenv("SERVER_ADDRESS")
	PORT := os.Getenv("SERVER_PORT")
	serve(ADDR, PORT, router)
}

package app

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/zeroidentidad/fiber-hex-api/logger"
)

func serve(addr string, port string, router *fiber.App) {
	go func() {
		logger.Info(fmt.Sprintf("Starting server on %s:%s ...", addr, port))
		err := router.Listen(fmt.Sprintf("%s:%s", addr, port))
		logger.Fatal(err.Error())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

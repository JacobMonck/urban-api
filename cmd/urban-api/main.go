package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()

	if err := app.Listen(":3000"); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve web server.")
	}

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGTERM)

	<-shutdownChannel

	log.Info().Msg("Gracefully shutting down server.")
	if err := app.Shutdown(); err != nil {
		log.Error().Err(err).Msg("Failed to shutdown web server.")
	}
}

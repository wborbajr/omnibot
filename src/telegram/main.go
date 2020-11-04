package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	bplogger "github.com/wborbajr/omnibot/src/utils"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func botHealthCheck(c *fiber.Ctx) error {
	return c.SendString("Telegram Bor v1.0.0")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", botHealthCheck)
}

func setupApp() {
	app := fiber.New(fiber.Config{
		Concurrency:          256 * 1024,
		CompressedFileSuffix: ".fiber.gz",
	})

	app.Use(limiter.New(limiter.Config{
		Duration: 10 * time.Second,
		Max:      14,
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2000",
		TimeZone:   "America/Sao_Paulo",
		Output:     os.Stdout,
	}))

	setupRoutes(app)

	log.Printf("[Telegram BOT v1.0.0] up and running: http://127.0.0.1:%s", "3001")
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

func main() {
	// Strting the Application
	bplogger.GeneralLogger.Println("Starting..")

	setupApp()
}


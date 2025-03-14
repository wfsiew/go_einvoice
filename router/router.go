package router

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/einvoice", logger.New())
    SetupInvoiceRoutes(api)
    SetupCNRoutes(api)
}
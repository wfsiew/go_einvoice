package router

import (
    "einvoice/controller"
    "github.com/gofiber/fiber/v2"
)

func SetupInvoiceRoutes(router fiber.Router) {
	inv := router.Group("/invoice")
	inv.Get("/events", controller.GetInvoiceEvent)
	inv.Get("/xml", controller.GetInvoiceXml)
}
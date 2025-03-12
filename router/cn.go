package router

import (
    "einvoice/controller"
    "github.com/gofiber/fiber/v2"
)

func SetupCNRoutes(router fiber.Router) {
	inv := router.Group("/cn")
	inv.Get("/events", controller.GetCNEvent)
	inv.Get("/xml", controller.GetCNXml)
}
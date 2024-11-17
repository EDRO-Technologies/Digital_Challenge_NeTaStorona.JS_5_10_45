package routes

import (
	"github.com/gofiber/fiber/v2"
	"other-side/api/controllers"
	mw "other-side/api/middleware"
	C "other-side/constants"
)

func SetupProductsRoutes(router fiber.Router) {
	router.Post("/prompt", mw.RateLimit(C.Tier7, 0), controllers.PostPrompt)

	router.Post("/documentPrompt", mw.RateLimit(C.Tier7, 0), controllers.DocumentPrompt)
}

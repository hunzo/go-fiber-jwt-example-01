package routes

import (
	"server/handlers"
	"server/validate"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(r *fiber.App) {
	public := r.Group("/")
	public.Get("/login", handlers.Login)

	private := r.Group("/api", validate.Protect())
	private.Get("/profile", handlers.Profile)

}

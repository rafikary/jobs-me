package routes

import (
	"jobsme/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Web routes
	app.Get("/", handlers.GetDashboard)

	// API routes
	api := app.Group("/api")

	// Job application endpoints
	jobs := api.Group("/jobs")
	jobs.Get("/", handlers.GetAllJobs)
	jobs.Get("/:id", handlers.GetJobByID)
	jobs.Post("/", handlers.CreateJob)
	jobs.Put("/:id", handlers.UpdateJob)
	jobs.Delete("/:id", handlers.DeleteJob)

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}

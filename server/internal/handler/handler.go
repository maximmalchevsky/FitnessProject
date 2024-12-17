package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"

	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "server/docs"
)

type Handler struct {
	db *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Router() *fiber.App {
	f := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	f.Use(logger.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	f.Get("/swagger/*", fiberSwagger.WrapHandler)
	f.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	})

	f.Get("/training/id/:id", h.TrainingGetByID)
	f.Get("/training", h.TrainingGetAll)
	f.Post("/training", h.TrainingCreate)
	f.Put("/training", h.TrainingUpdate)
	f.Delete("/training/id/:id", h.TrainingDelete)
	f.Put("/training/id/:id/confirm", h.ConfirmTraining)
	f.Put("/training/id/:id/cancel", h.CancelTraining)

	f.Get("/pass", h.PassGetAll)
	f.Get("/pass/id/:id", h.PassGetByID)
	f.Post("/pass", h.PassCreate)
	f.Delete("/pass/id/:id", h.PassDelete)

	return f
}

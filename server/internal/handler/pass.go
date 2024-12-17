package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/models"
	"server/internal/repository/postgres"
)

// PassGetByID
// @Tags         pass
// @Summary      Получение абонемента по ID
// @Description  Получение информации о абонементе по его уникальному идентификатору
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID абонемента"
// @Success      200  {object}  models.Pass "Данные абонемента по заданному ID"
// @Failure      400  {object}  models.ErrorResponse "Некорректный идентификатор абонемента"
// @Failure      500  {object}  models.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /pass/id/{id} [get]
func (h *Handler) PassGetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	exists, err := postgres.PassExistsID(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "pass doesn't exist"})
	}

	pass, err := postgres.PassGetByID(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(pass)
}

// PassGetAll
// @Tags         pass
// @Summary      Получение списка всех абонементов
// @Description  Получение списка всех доступных абонементов в системе
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Pass "Список всех абонементов"
// @Failure      500  {object}  models.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /pass [get]
func (h *Handler) PassGetAll(c *fiber.Ctx) error {
	passes, err := postgres.PassGetAll(h.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(passes)
}

// PassCreate
// @Tags         pass
// @Summary      Создание нового абонемента
// @Description  Создание нового абонемента с указанием имени, телефона, типа и продолжительности (в месяцах).
// @Accept       json
// @Produce      json
// @Param        pass  body      models.CreatePass  true  "Данные абонемента"
// @Success      200   {object}  models.Pass        "Успешное создание абонемента"
// @Failure      400   {object}  models.ErrorResponse "Некорректный запрос"
// @Failure      500   {object}  models.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /pass [post]
func (h *Handler) PassCreate(c *fiber.Ctx) error {
	var req models.CreatePass
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	pass := &models.Pass{
		Name:     req.Name,
		Phone:    req.Phone,
		Type:     req.Type,
		Duration: req.Duration,
	}

	res, err := postgres.PassCreate(h.db, pass)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// PassDelete
// @Tags         pass
// @Summary      Удаление абонемента
// @Description  Удаление абонемента по идентификатору.
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID абонемента"
// @Success      200  {object}  map[string]string "Успешное удаление абонемента"
// @Failure      400  {object}  models.ErrorResponse   "Некорректный идентификатор абонемента"
// @Failure      500  {object}  models.ErrorResponse   "Внутренняя ошибка сервера"
// @Router       /pass/id/{id} [delete]
func (h *Handler) PassDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	exists, err := postgres.PassExistsID(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "pass doesn't exist"})
	}

	err = postgres.PassDelete(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

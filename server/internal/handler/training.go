package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server/internal/models"
	"server/internal/repository/postgres"
	"strconv"
)

// TrainingGetByID
// @Tags         training
// @Summary      Получение тренировки по ID
// @Description  Получение информации о тренировке на основе переданного идентификатора
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID тренировки"
// @Success      200  {object}  models.Training "Информация о тренировке"
// @Failure      400  {object}  models.ErrorResponse "Некорректный идентификатор тренировки"
// @Failure      500  {object}  models.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /training/id/{id} [get]
func (h *Handler) TrainingGetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	exists, err := postgres.TrainingExistsID(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "training doesn't exist"})
	}

	training, err := postgres.TrainingGetByID(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(training)
}

// TrainingCreate
// @Tags         training
// @Summary      Создание новой тренировки
// @Description  Добавление новой записи о тренировке в базу данных
// @Accept       json
// @Produce      json
// @Param        training  body      models.CreateTraining  true  "Данные для создания тренировки"
// @Success      200       {object}  models.Training        "Успешное создание тренировки"
// @Failure      400       {object}  models.ErrorResponse   "Некорректные данные запроса"
// @Failure      500       {object}  models.ErrorResponse   "Внутренняя ошибка сервера"
// @Router       /training [post]
func (h *Handler) TrainingCreate(c *fiber.Ctx) error {
	var request models.CreateTraining
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	log.Println("TrainingExists")
	exists, err := postgres.TrainingExists(h.db, request.Name, request.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "training already exists"})
	}

	training := &models.Training{
		Name:  request.Name,
		Phone: request.Phone,
	}

	log.Println("TrainingCreate")
	res, err := postgres.TrainingCreate(h.db, training)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// TrainingUpdate
// @Tags         training
// @Summary      Обновление существующей тренировки
// @Description  Обновление данных тренировки с проверкой на существование записи в базе данных
// @Accept       json
// @Produce      json
// @Param        training  body      models.UpdateTraining  true  "Данные для обновления тренировки"
// @Success      200       {object}  models.Training        "Успешное обновление тренировки"
// @Failure      400       {object}  models.ErrorResponse   "Некорректные данные запроса"
// @Failure      404       {object}  models.ErrorResponse   "Тренировка не найдена"
// @Failure      500       {object}  models.ErrorResponse   "Внутренняя ошибка сервера"
// @Router       /training [put]
func (h *Handler) TrainingUpdate(c *fiber.Ctx) error {
	var request models.UpdateTraining
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	exists, err := postgres.TrainingExistsID(h.db, request.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "training doesn't exist"})
	}

	training := &models.Training{
		Name:  request.Name,
		Phone: request.Phone,
		ID:    request.ID,
	}
	res, err := postgres.TrainingUpdate(h.db, training)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// TrainingDelete
// @Tags         training
// @Summary      Удаление тренировки
// @Description  Удаление существующей записи тренировки по её идентификатору
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID тренировки"
// @Success      200  {object}  map[string]string  "Успешное удаление тренировки"
// @Failure      400  {object}  models.ErrorResponse   "Некорректный идентификатор тренировки"
// @Failure      500  {object}  models.ErrorResponse   "Внутренняя ошибка сервера"
// @Router       /training/id/{id} [delete]
func (h *Handler) TrainingDelete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	exists, err := postgres.TrainingExistsID(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "training doesn't exist"})
	}

	err = postgres.TrainingDelete(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

// TrainingGetAll
// @Tags         training
// @Summary      Получение всех тренировок
// @Description  Получение списка всех существующих тренировок из базы данных
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Training  "Список всех тренировок"
// @Failure      500  {object}  models.ErrorResponse  "Внутренняя ошибка сервера"
// @Router       /training [get]
func (h *Handler) TrainingGetAll(c *fiber.Ctx) error {
	trainings, err := postgres.TrainingGetAll(h.db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(trainings)
}

// ConfirmTraining
// @Tags         training
// @Summary      Подтверждение тренировки
// @Description  Обновление статуса тренировки на "подтверждено"
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID тренировки"
// @Success      200  {object}  map[string]string  "Тренировка успешно подтверждена"
// @Failure      400  {object}  models.ErrorResponse "Некорректный идентификатор тренировки"
// @Failure      500  {object}  models.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /training/id/{id}/confirm [put]
func (h *Handler) ConfirmTraining(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	exists, err := postgres.TrainingExistsID(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "training doesn't exist"})
	}

	err = postgres.ConfirmationUpdate(h.db, "1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

// CancelTraining
// @Tags         training
// @Summary      Отмена тренировки
// @Description  Обновление статуса тренировки на "отменено"
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID тренировки"
// @Success      200  {object}  map[string]string  "Тренировка успешно отменена"
// @Failure      400  {object}  models.ErrorResponse "Некорректный идентификатор тренировки"
// @Failure      500  {object}  models.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /training/id/{id}/cancel [put]
func (h *Handler) CancelTraining(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	exists, err := postgres.TrainingExistsID(h.db, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "training doesn't exist"})
	}

	err = postgres.ConfirmationUpdate(h.db, "0", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

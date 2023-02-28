package handler

import (
	"GOLANG-TODOS/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type todoHandler struct {
	todoSrv service.TodoService
}

func NewTodoHandler(todoSrv service.TodoService) todoHandler {
	return todoHandler{todoSrv: todoSrv}
}

func (h *todoHandler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.todoSrv.GetTodos()
	if err != nil {
		return handleError(c, err)
	}
	return c.JSON(todos)
}

func (h *todoHandler) GetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return handleError(c, err)
	}

	todo, err := h.todoSrv.GetTodo(id)
	if err != nil {
		return handleError(c, err)
	}
	return c.JSON(todo)
}

func (h *todoHandler) NewTodo(c *fiber.Ctx) error {
	todo := service.TodoRequest{}

	err := c.BodyParser(&todo)
	if err != nil {
		return handleError(c, err)
	}

	errors := validateStruct(todo)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newTodo, err := h.todoSrv.NewTodo(todo)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(newTodo)
}

func (h *todoHandler) UpdateTodo(c *fiber.Ctx) error {
	todo := service.TodoUpdateRequest{}
	err := c.BodyParser(&todo)
	if err != nil {
		return handleError(c, err)
	}

	errors := validateStruct(todo)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	saveTodo, err := h.todoSrv.UpdateTodo(todo)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(saveTodo)
}

func (h *todoHandler) DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return handleError(c, err)
	}

	err = h.todoSrv.DeleteTodo(uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})
}

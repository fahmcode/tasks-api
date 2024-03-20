package controllers

import (
	"tasksapi/models"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Get Todos - GET /api/todos
func GetAllTodos(ctx *fiber.Ctx) {
	todos := []models.Todo{}
	collection := mgm.Coll(&models.Todo{})

	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})
}

// GetTodoById - GET /api/todos/:id
func GetTodoById(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// CreateTodo - POST /api/todos
func CreateTodo(ctx *fiber.Ctx) {
	params := new(struct {
		Title       string
		Description string
	})
	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": "Title or Description not specified.",
		})
		return
	}

	todo := models.CreateTodo(params.Title, params.Description)
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// UpdateTodo - PATCH /api/todos/:id
func UpdateTodo(ctx *fiber.Ctx){
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	todo.Done = !todo.Done
	err = collection.Update(todo)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}

// DeleteTodo - DELETE /api/todos/:id
func DeleteTodo(ctx *fiber.Ctx){
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	err = collection.Delete(todo)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
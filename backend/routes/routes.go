package routes

import (
	"todo-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/api/todos", controllers.GetTodos)
    app.Post("/api/todos", controllers.CreateTodo)
    app.Patch("/api/todos/:id", controllers.UpdateTodo)
    app.Delete("/api/todos/:id", controllers.DeleteTodo)
}

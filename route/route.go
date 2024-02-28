package route

import (
	"github.com/devsmranjan/golang-fiber-basic-todo-app/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}

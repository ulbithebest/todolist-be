package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ulbithebest/todolist-be/controller"
)

func SetupTaskRoutes(app *fiber.App) {
	app.Post("/register", controller.RegisterUser)
	app.Post("/login", controller.LoginUser)

	app.Use(controller.Authenticate)
	app.Get("/getme", controller.GetMe)
	app.Get("/tasks", controller.GetAllTask)
	app.Get("/task/get/:id", controller.GetTaskById)
	app.Post("/task/insert", controller.InsertTask)
	app.Put("/task/update/:id", controller.UpdateTask)
	app.Delete("/task/delete/:id", controller.DeleteTask)

	app.Post("/logout", controller.LogoutUser)
}

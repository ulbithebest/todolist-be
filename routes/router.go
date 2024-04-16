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
	app.Get("/task/get/:id_task", controller.GetTaskById)
	app.Post("/task/insert", controller.InsertTask)
	app.Put("/task/update/:id_task", controller.UpdateTask)
	app.Delete("/task/delete/:id_task", controller.DeleteTask)

	app.Get("/roles", controller.GetAllRole)
	app.Get("/role/get/:id_role", controller.GetRoleById)
	app.Post("/role/insert", controller.InsertRole)
	app.Put("/role/update/:id_role", controller.UpdateRole)
	app.Delete("/role/delete/:id_role", controller.DeleteRole)

	app.Post("/logout", controller.LogoutUser)
}

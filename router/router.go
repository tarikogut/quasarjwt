package router

import (
	"estpguiapi/controllers"
	"estpguiapi/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", controllers.Status)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", middlewares.Protected(), controllers.GetUser)
	user.Post("/", middlewares.Protected(), controllers.CreateUser)
	user.Patch("/:id", middlewares.Protected(), controllers.UpdateUser)
	user.Delete("/:id", middlewares.Protected(), controllers.DeleteUser)
	estp := api.Group("/estp")
	estp.Post("/add", middlewares.Protected(), controllers.EstpAdd)
	estp.Post("/update/:id", middlewares.Protected(), controllers.EstpUpdate)
	estp.Post("/delete/:id", middlewares.Protected(), controllers.EstpDelete)
	estp.Get("/list", middlewares.Protected(), controllers.EstpList)
	estp.Get("/details/:id", middlewares.Protected(), controllers.EstpDetails)
	//estp.Post("/:id/:service/:action?", middlewares.Protected(), controllers.EstpControllers)
}

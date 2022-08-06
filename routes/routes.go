package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rekib0023/go-crud-apis/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	GetUserRoutes(api)
	GetProductRoutes(api)
}

func GetUserRoutes(api fiber.Router) {
	group := api.Group("/users")
	userRoutes(group)
}

func GetProductRoutes(api fiber.Router) {
	group := api.Group("/products")
	productRoutes(group)
}

func userRoutes(api fiber.Router) {
	api.Post("/", controllers.CreateUser)
	api.Get("/", controllers.GetUsers)
	api.Get("/:id", controllers.GetUser)
	api.Put("/:id", controllers.UpdateUser)
	api.Delete("/:id", controllers.DeleteUser)
}

func productRoutes(api fiber.Router) {
	api.Post("/", controllers.CreateProduct)
	api.Get("/", controllers.GetProducts)
	api.Get("/:id", controllers.GetProduct)
	api.Put("/:id", controllers.UpdateProduct)
	api.Delete("/:id", controllers.DeleteProduct)
}

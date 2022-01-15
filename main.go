package main

import (
	"log"

	"github.com/Chaksack/centri-backend-web/controllers"
	"github.com/Chaksack/centri-backend-web/database"
	"github.com/Chaksack/centri-backend-web/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/api", welcome)
	//controllers endpoints
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)
	app.Get("api/staff", controllers.Staff)
	app.Post("api/Logout", controllers.Logout)
	//user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Category
	app.Post("/api/categorys", routes.CreateCategory)
	app.Get("/api/categorys", routes.GetCategorys)
	app.Get("/api/categorys/:id", routes.GetCategory)
	app.Put("/api/categorys/:id", routes.UpdateCategory)
	app.Delete("/api/categorys/:id", routes.DeleteCategory)
	// Product
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	// Invoice
	app.Post("/api/invoices", routes.CreateInvoice)
	app.Get("/api/invoices", routes.GetInvoices)
	app.Get("/api/invoices/:id", routes.GetInvoice)

}

func main() {
	database.ConnectDb()
	app := fiber.New()
	//new addition
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

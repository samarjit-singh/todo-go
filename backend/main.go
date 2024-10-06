package main

import (
	"log"
	"todo-go/config"
	"todo-go/database"
	"todo-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	// folloing : is the pointer and address concept
	// var x int = 5;
	// var p *int = &x;

	// fmt.Println("prints the address", p)
	// fmt.Println("prints the actual value", *p)

    // Load environment variables
    config.LoadEnv()

    // Initialize MongoDB connection
    database.ConnectDB()

    // Setup routes
    routes.SetupRoutes(app)

    // Start the server
    PORT := config.GetPort()
    log.Fatal(app.Listen(":" + PORT))
}

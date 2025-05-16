package main

import (
	"log"
	"os"
	"todo-api/db"
	"todo-api/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	defer db.CloseDB()

	app := fiber.New()

	// Маршруты
	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks", handlers.GetAllTasks)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}

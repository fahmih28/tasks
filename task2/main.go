package main

import "github.com/gofiber/fiber/v2"

func main() {
	server := fiber.New()
	server.Get("/api/v1/books/summary", getSummary)
	server.Listen(":8383")
}

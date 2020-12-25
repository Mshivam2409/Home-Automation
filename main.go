package main

import (
	"github.com/Mshivam2409/Home-Automation/auth"
	"github.com/Mshivam2409/Home-Automation/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New()
	database.IntializeMongoDB()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Home Automation Ssytem is up!")
	})
	app.Post("/signup", auth.SignUp)
	app.Post("/signin", auth.SignIn)
	app.Listen(":5000")
}

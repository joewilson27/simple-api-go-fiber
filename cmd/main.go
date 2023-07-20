package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"simple-go-fiber/pkg/books"
	"simple-go-fiber/pkg/common/config"
	"simple-go-fiber/pkg/common/db"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	
	app := fiber.New()
	db := db.Init(c.DBUrl) // load the string connection to database

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})

	books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}




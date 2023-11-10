package main

import (
	"fdzz/infrastructure/inmemory"
	"fdzz/usecase"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	inmem := inmemory.NewDB()
	uc := usecase.NewBookStore(inmem)

	app.Post("/books", nil)

	app.Get("/books", func(c *fiber.Ctx) error {
		res, err := uc.Read(c.Context())
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(res)
	})

	app.Put("/books", nil)

	app.Delete("/books", nil)

	log.Fatal(app.Listen(":3000"))
}

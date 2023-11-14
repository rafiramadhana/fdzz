package main

import (
	"fdzz/infrastructure/sqlite"
	"fdzz/usecase"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// im := inmemory.NewDB()
	sq := sqlite.NewDB()
	uc := usecase.NewBookStore(sq)

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

package main

import (
	"strconv"

	"github.com/ali2210/postgres_client/postgresdb"
	"github.com/gofiber/fiber/v2"
)

type Operator struct {
	X, Y int64
}

func main() {

	app_web := fiber.New()

	add := Operator{X: 2, Y: 3}

	app_web.Get("/", func(c *fiber.Ctx) error {

		go func() {
			postgresdb.ORM()
		}()

		return c.SendString(strconv.Itoa(int(add.X) + int(add.Y)))
	})

	err := app_web.Listen(":3000")
	if err != nil {
		return
	}

}

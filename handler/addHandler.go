package handler

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/usysrc/ausleihe/data"
	"github.com/usysrc/ausleihe/db"
)

// add items to the db
func ItemHandler(c *fiber.Ctx) error {
	fmt.Print(string(c.Body()))
	var newItem data.Item
	if err := c.BodyParser(&newItem); err != nil {
		c.Status(500)
		return err
	}

	_, err := db.Postgres.Exec(context.Background(), "INSERT into items (name) VALUES ($1)", newItem.Name)
	if err != nil {
		c.Status(500)
		return err
	}
	return IndexHandler(c)
}

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/usysrc/ausleihe/data"
	"github.com/usysrc/ausleihe/db"
)

// view items
func IndexHandler(c *fiber.Ctx) error {
	rows, err := db.Postgres.Query(context.Background(), "SELECT id, name FROM items")
	if err != nil {
		return err
	}
	defer rows.Close()

	var items []data.Item
	for rows.Next() {
		var item data.Item
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			return err
		}
		items = append(items, item)
	}

	return c.Render("index", fiber.Map{
		"Items": items,
	})
}

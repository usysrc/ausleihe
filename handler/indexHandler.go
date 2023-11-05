package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/usysrc/ausleihe/data"
	"github.com/usysrc/ausleihe/db"
)

// view items
func IndexHandler(c *fiber.Ctx) error {
	desiredStart := time.Now()
	desiredEnd := time.Now().AddDate(0, 0, 7)

	// get all items that are not reserved for the time frame
	rows, err := db.Postgres.Query(context.Background(), `
    SELECT * FROM items
    WHERE id NOT IN (
      SELECT DISTINCT item_id FROM reservations
      WHERE (start_datetime, end_datetime) OVERLAPS ($1, $2)
    )`,
		desiredStart, desiredEnd)
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
		"Items":    items,
		"Now":      desiredStart.Format("2006-01-02T15:04"),
		"NextWeek": desiredEnd.Format("2006-01-02T15:04"),
	})
}

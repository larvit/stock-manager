package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/larvit/stock-manager/api/src/db"
)

// PostAddStockItems godoc
// @Summary Add stock balances
// @ID add-balances
// @Accept  json
// @Produce  json
// @Param body body []db.StockItem true "Array of stock items to add balances to"
// @Success 200 {string} JSON string "OK"
// @Failure 400 {object} []ResJSONError
// @Failure 500 {object} []ResJSONError
// @Router /add-stock-items [post]
func (h Handlers) PostAddStockItems(c *fiber.Ctx) error {
	input := new([]db.StockItem)

	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON([]ResJSONError{
			{Error: err.Error()},
		})
	}

	for _, stockItem := range *input {
		h.Db.AddStockItem(stockItem)
	}

	return c.Status(200).JSON("OK")
}

// PostSubtractStockItems godoc
// @Summary Subtract stock balances
// @ID subtract-balances
// @Accept  json
// @Produce  json
// @Param body body []StockItemAdjuster true "Array of stock items to subtract balance in"
// @Success 200 {string} JSON string "OK"
// @Failure 400 {object} []ResJSONError
// @Failure 500 {object} []ResJSONError
// @Router /subtract-stock-items [post]
func (h Handlers) PostSubtractStockItems(c *fiber.Ctx) error {
	input := new([]StockItemAdjuster)

	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON([]ResJSONError{
			{Error: err.Error()},
		})
	}

	for _, stockItem := range *input {
		h.Db.SubtractStockItem(stockItem.ID, stockItem.Amount)
	}

	return c.Status(200).JSON("OK")
}

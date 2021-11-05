package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type BalancesQuery struct {
	Ids []string `query:"ids[]"`
}

// GetStockBalances godoc
// @Summary Get stock balances
// @Description Gets a list of all stock and their balances
// @ID get-stock-balances
// @Accept  json
// @Produce  json
// @Param ids[] query string false "Array of ids"
// @Success 200 {object} map[string]db.StockItem
// @Failure 500 {object} []ResJSONError
// @Router /stock-balances [get]
func (h Handlers) GetStockBalances(c *fiber.Ctx) error {
	query := new(BalancesQuery)

	err := c.QueryParser(query)
	if err != nil {
		return c.Status(400).JSON([]ResJSONError{
			{Error: err.Error()},
		})
	}

	return c.JSON(h.Db.GetStockItemsByIds(query.Ids))
}

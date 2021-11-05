package db

import (
	"errors"
	"strconv"
)

// Will update the name of existing stock items with the same ID
func (db Db) AddStockItem(stockItem StockItem) {
	if val, ok := db.Stock[stockItem.ID]; ok {
		stockItem.Balance += val.Balance

		// If no new name is provided, use the old name
		if val.Name == "" {
			val.Name = stockItem.Name
		}

		db.Stock[stockItem.ID] = stockItem
	} else {
		db.Stock[stockItem.ID] = stockItem
	}
}

func (db Db) SubtractStockItem(id string, amount uint64) error {
	if val, ok := db.Stock[id]; ok {
		if val.Balance > amount {
			val.Balance -= amount
			db.Stock[id] = val
		} else if val.Balance == amount {
			delete(db.Stock, id)
		} else {
			return errors.New("Can not deduct higher amount than existing balance. Current balace: " + strconv.FormatUint(val.Balance, 10) + ". Amount to deduct: " + strconv.FormatUint(amount, 10))
		}
	} else {
		return errors.New("No stock item exists with id: \"" + id + "\"")
	}

	return nil
}

func (db Db) GetAllStockItems() map[string]StockItem {
	return db.Stock
}

func (db Db) GetStockItemsByIds(ids []string) map[string]StockItem {
	var slice = make(map[string]StockItem)

	for _, id := range ids {
		stockItem, ok := db.Stock[id]
		if ok {
			slice[id] = stockItem
		}
	}

	return slice
}

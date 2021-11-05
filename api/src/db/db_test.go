package db

import (
	"testing"

	"github.com/larvit/stock-manager/api/src/utils"
)

func TestAddStockItem(t *testing.T) {
	dbInstance := Db{
		Log:   utils.GetLog(),
		Stock: make(map[string]StockItem),
	}

	dbInstance.AddStockItem(StockItem{
		ID:      "xxx",
		Name:    "A thing",
		Balance: 3,
	})

	// Adding again should only increase balance
	dbInstance.AddStockItem(StockItem{
		ID:      "xxx",
		Name:    "A thing",
		Balance: 2,
	})

	savedStock := dbInstance.GetAllStockItems()

	if len(savedStock) != 1 {
		t.Errorf("Expected 1 saved stock item, found %d", len(savedStock))
	}

	if savedStock["xxx"].Balance != 5 {
		t.Errorf("Expected total stock balance for xxx to be 5, but is %d", savedStock["xxx"].Balance)
	}
}

func TestAddMultipleStockItems(t *testing.T) {
	dbInstance := Db{
		Log:   utils.GetLog(),
		Stock: make(map[string]StockItem),
	}

	dbInstance.AddStockItem(StockItem{
		ID:      "xxx",
		Name:    "A thing",
		Balance: 3,
	})

	dbInstance.AddStockItem(StockItem{
		ID:      "xxy",
		Name:    "Another thing",
		Balance: 2,
	})

	savedStock := dbInstance.GetAllStockItems()

	if len(savedStock) != 2 {
		t.Errorf("Expected 2 saved stock item, found %d", len(savedStock))
	}

	if savedStock["xxx"].Balance != 3 {
		t.Errorf("Expected total stock balance for xxx to be 3, but is %d", savedStock["xxx"].Balance)
	}
}

func TestSubtractStockItem(t *testing.T) {
	dbInstance := Db{
		Log:   utils.GetLog(),
		Stock: make(map[string]StockItem),
	}

	dbInstance.AddStockItem(StockItem{
		ID:      "xxx",
		Name:    "A thing",
		Balance: 3,
	})

	dbInstance.AddStockItem(StockItem{
		ID:      "xxy",
		Name:    "Another thing",
		Balance: 2,
	})

	savedStock := dbInstance.GetAllStockItems()

	if len(savedStock) != 2 {
		t.Errorf("Expected 2 saved stock item, found %d", len(savedStock))
	}

	if savedStock["xxx"].Balance != 3 {
		t.Errorf("Expected total stock balance for xxx to be 3, but is %d", savedStock["xxx"].Balance)
	}
}

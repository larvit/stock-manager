package db

import "go.uber.org/zap"

// Db struct is the "parent" struct for all database functions
type Db struct {
	Log   *zap.SugaredLogger
	Stock map[string]StockItem // Our in-memory database, index by stock item ID
}

// A stock item
type StockItem struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance uint64 `json:"balance"`
}

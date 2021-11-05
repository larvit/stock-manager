package handlers

import (
	"github.com/larvit/stock-manager/api/src/db"
	"go.uber.org/zap"
)

// Handlers is the overall struct for all http request handlers
type Handlers struct {
	Db  db.Db
	Log *zap.SugaredLogger
}

// ResJSONError is an error field that is used in JSON error responses
type ResJSONError struct {
	Error string `json:"error"`
	Field string `json:"field,omitempty"`
}

// Used when adjusting the balance of a StockItem
type StockItemAdjuster struct {
	Amount uint64 `json:"amount"`
	ID     string `json:"id"`
}

package helper

import (
	"time"
)

type StoreTransactionType struct {
	Type_Name   string
	Description string
	Created_at  time.Time
	Created_by  int
}

type UpdateTransactionType struct {
	Type_Name   string
	Description string
	Updated_at  time.Time
	Updated_by  int
}

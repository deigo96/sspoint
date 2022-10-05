package helper

import "time"

type TransactionType struct {
	Type_Name        string `json:"parent_id" form:"typeName"`
	Type_Description string `json:"child_id" form:"typeDescription"`
}

type IdTransactionType struct {
	Id_Transaction_type int `json:"idTransaction" form:"typeId"`
}

func current_time(format string) string {
	t := time.Now()
	current_time := t.Format(format)
	return current_time
}

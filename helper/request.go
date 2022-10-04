package helper

import "time"

type RegisterReferral struct {
	Parent_id    int `json:"parent_id" form:"parent_id"`
	Child_id     int `json:"child_id" form:"child_id"`
	To_Parent_id int `json:"to_parent_id" form:"to_parent_id"`
	// Created_by string `json:"created_at" form:"created_by"`
}

func current_time(format string) string {
	t := time.Now()
	current_time := t.Format(format)
	return current_time
}

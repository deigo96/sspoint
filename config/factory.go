package config

import (
	"pointHistory-service/helper"
	"pointHistory-service/model"
)

func RepositoryFactory(dbCon *DatabaseConnection) helper.PointHistoryList {
	var Repository helper.PointHistoryList

	if dbCon.Driver == PostgreSQL {
		Repository = model.NewPointHistoryService(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return Repository
}

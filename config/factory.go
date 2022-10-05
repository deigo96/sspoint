package config

import (
	"transactionPoint-service/helper"
	"transactionPoint-service/model"
)

func RepositoryFactory(dbCon *DatabaseConnection) helper.TransactionPointList {
	var Repository helper.TransactionPointList

	if dbCon.Driver == PostgreSQL {
		Repository = model.NewStoreReward(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return Repository
}

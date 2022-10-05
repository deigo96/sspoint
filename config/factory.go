package config

import (
	"transactionType-service/helper"
	"transactionType-service/model"
)

func RepositoryFactory(dbCon *DatabaseConnection) helper.TransactionTypeList {
	var Repository helper.TransactionTypeList

	if dbCon.Driver == PostgreSQL {
		Repository = model.NewStoreReward(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return Repository
}

package config

import (
	"reward-list-service/helper"
	"reward-list-service/model"
)

func RepositoryFactory(dbCon *DatabaseConnection) helper.RewardList {
	var Repository helper.RewardList

	if dbCon.Driver == PostgreSQL {
		Repository = model.NewStoreReward(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return Repository
}

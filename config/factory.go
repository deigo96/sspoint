package config

import (
	"referralUser-service/helper"
	"referralUser-service/model"
)

func RepositoryFactory(dbCon *DatabaseConnection) helper.AuthService {
	var Repository helper.AuthService

	if dbCon.Driver == PostgreSQL {
		Repository = model.NewRegisterReferral(dbCon.PostgreSQL)
	} else {
		panic("Database driver not supported")
	}

	return Repository
}

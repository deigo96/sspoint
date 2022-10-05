package model

import (
	"errors"
	"time"
	"transactionType-service/helper"

	"gorm.io/gorm"
)

type TransactionTypeRepo struct {
	db *gorm.DB
}

func NewStoreReward(db *gorm.DB) helper.TransactionTypeList {
	return &TransactionTypeRepo{
		db: db,
	}
}

func (t *TransactionTypeRepo) StoreTransactionTypeList(id int, s helper.TransactionType) error {
	record := helper.StoreTransactionType{
		Type_Name:   s.Type_Name,
		Description: s.Type_Description,
		Created_by:  id,
		Created_at:  time.Now(),
	}

	res := t.db.Table("transaction_type").Create(&record)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *TransactionTypeRepo) UpdateTransactionTypeList(id int, IdTransactionType int, s helper.TransactionType) error {
	var idr helper.IdTransactionType

	checkId := r.db.Table("transaction_type").Where("id = ?", IdTransactionType).Find(&idr)

	if checkId.Error != nil {
		return checkId.Error
	}

	if checkId.RowsAffected == 0 {
		return errors.New("Not found")
	}

	record := helper.UpdateTransactionType{
		Type_Name:   s.Type_Name,
		Description: s.Type_Description,
		Updated_by:  id,
		Updated_at:  time.Now(),
	}

	res := r.db.Table("transaction_type").Where("id = ?", IdTransactionType).Updates(&record)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

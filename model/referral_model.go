package model

import (
	"errors"
	"time"
	"transactionPoint-service/helper"

	"gorm.io/gorm"
)

type TransactionPointRepo struct {
	db *gorm.DB
}

func NewStoreReward(db *gorm.DB) helper.TransactionPointList {
	return &TransactionPointRepo{
		db: db,
	}
}

func (t *TransactionPointRepo) StoreTransactionPointList(id int, s helper.TransactionPoint) error {
	record := helper.StoreTransactionPoint{
		Trx_Type_Id:  s.Trx_Type_Id,
		Reward_Id:    s.Reward_Id,
		Ss_Point:     s.Point,
		Own_Pct:      s.Own_Pct,
		Referral_Pct: s.Referral_Pct,
		Branch_Pct:   s.Branch_Pct,
		Created_by:   id,
		Created_at:   time.Now(),
	}

	recordHis := helper.StoreTransactionPointHistory{
		Trx_Type_Id:       s.Trx_Type_Id,
		Reward_Id:         s.Reward_Id,
		Ss_Point:          s.Point,
		Own_User_Pct:      s.Own_Pct,
		Referral_User_Pct: s.Referral_Pct,
		Branch_Pct:        s.Branch_Pct,
		Created_by:        id,
		Created_at:        time.Now(),
	}

	res := t.db.Table("transaction_point").Create(&record)
	if res.Error != nil {
		return res.Error
	}

	resHis := t.db.Table("transaction_point_hist").Create(&recordHis)
	if resHis.Error != nil {
		return res.Error
	}

	return nil
}

func (t *TransactionPointRepo) UpdateTransactionPointList(id int, IdTransactionPoint int, s helper.TransactionPoint) error {
	var idr helper.IdTransactionPoint

	checkId := t.db.Table("transaction_point").Where("id = ?", IdTransactionPoint).Find(&idr)

	if checkId.Error != nil {
		return checkId.Error
	}

	if checkId.RowsAffected == 0 {
		return errors.New("Not found")
	}

	record := helper.UpdateTransactionPoint{
		Trx_Type_Id:  s.Trx_Type_Id,
		Reward_Id:    s.Reward_Id,
		Ss_Point:     s.Point,
		Own_Pct:      s.Own_Pct,
		Referral_Pct: s.Referral_Pct,
		Branch_Pct:   s.Branch_Pct,
		Updated_by:   id,
		Updated_at:   time.Now(),
	}

	recordHis := helper.StoreTransactionPointHistory{
		Trx_Type_Id:       s.Trx_Type_Id,
		Reward_Id:         s.Reward_Id,
		Ss_Point:          s.Point,
		Own_User_Pct:      s.Own_Pct,
		Referral_User_Pct: s.Referral_Pct,
		Branch_Pct:        s.Branch_Pct,
		Created_by:        id,
		Created_at:        time.Now(),
	}

	res := t.db.Table("transaction_point").Where("id = ?", IdTransactionPoint).Updates(&record)
	if res.Error != nil {
		return res.Error
	}

	resHis := t.db.Table("transaction_point_hist").Create(&recordHis)
	if resHis.Error != nil {
		return res.Error
	}

	return nil
}

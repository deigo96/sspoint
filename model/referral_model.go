package model

import (
	"pointHistory-service/helper"

	"gorm.io/gorm"
)

type PointRepo struct {
	db *gorm.DB
}

func NewPointHistoryService(db *gorm.DB) helper.PointHistoryList {
	return &PointRepo{
		db: db,
	}
}

type User struct {
	Username string `json:"name"`
}

func a() int {
	return 1
}

func (p *PointRepo) GetPointHistoryList(id int, t helper.TrxPointRequest) (d []helper.Data) {
	// fromTrxDate, _ := time.Parse("2006-01-01", t.FromTrxDate)
	// toTrxDate, _ := time.Parse("2006-01-01", t.ToTrxDate)

	res := p.db.Table("referral_transaction").Select("referral_transaction.trx_date, transaction_type.type_name, referral_transaction.ss_point_before, referral_transaction.ss_point_trx, referral_transaction.ss_point_after").Joins("left join transaction_point on referral_transaction.trx_point_id = transaction_point.id").Joins("left join transaction_type on transaction_point.trx_type_id = transaction_type.id").Where("referral_transaction.trx_date >= ? AND referral_transaction.trx_date < ?", t.FromTrxDate, t.ToTrxDate).Find(&d)
	if res.Error != nil {
		return d
	}
	return d
}

func (p *PointRepo) StorePointHistoryList(s helper.StoreTrxPointReq) error {
	res := p.db.Table("referral_transaction").Create(&s)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

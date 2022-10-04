package model

import (
	"errors"
	"reward-list-service/helper"
	"time"

	"gorm.io/gorm"
)

type RewardRepo struct {
	db *gorm.DB
}

func NewStoreReward(db *gorm.DB) helper.RewardList {
	return &RewardRepo{
		db: db,
	}
}

func (r *RewardRepo) GetRewardList() (d []helper.Data) {
	// var req []helper.GetRewardRequst
	res := r.db.Table("reward_master").Find(&d)
	if res.Error != nil {
		return d
	}
	return d
}

func (r *RewardRepo) StoreRewardList(id int, s helper.StoreRewardRequest) error {
	record := helper.StoreData{
		Reward_Name: s.Reward_Name,
		Img_Name:    s.Reward_Image,
		Description: s.Reward_Description,
		Created_By:  id,
		Created_At:  time.Now(),
	}

	res := r.db.Table("reward_master").Create(&record)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *RewardRepo) UpdateRewardList(id int, idReward int, s helper.StoreRewardRequest) error {
	var idr helper.Idreward

	checkId := r.db.Table("reward_master").Where("id = ?", idReward).Find(&idr)

	if checkId.Error != nil {
		return checkId.Error
	}

	if checkId.RowsAffected == 0 {
		return errors.New("Not found")
	}

	record := helper.UpdateData{
		Reward_Name: s.Reward_Name,
		Img_Name:    s.Reward_Image,
		Description: s.Reward_Description,
		Updated_By:  id,
		Updated_At:  time.Now(),
	}

	res := r.db.Table("reward_master").Where("id = ?", idReward).Updates(&record)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

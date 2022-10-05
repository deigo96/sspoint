package helper

import "time"

type TransactionPoint struct {
	Trx_Type_Id  int `json:"trxTypeId" form:"trxTypeId"`
	Reward_Id    int `json:"rewardId" form:"rewardId"`
	Point        int `json:"point" form:"point"`
	Own_Pct      int `json:"ownPct" form:"ownPct"`
	Referral_Pct int `json:"referralPct" form:"referralPct"`
	Branch_Pct   int `json:"branchPct" form:"branchPct"`
}

type IdTransactionPoint struct {
	Id_Transaction_Point int `json:"idTransaction" form:"trxPointId"`
}

func current_time(format string) string {
	t := time.Now()
	current_time := t.Format(format)
	return current_time
}

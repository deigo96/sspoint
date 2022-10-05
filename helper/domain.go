package helper

import (
	"time"
)

type StoreTransactionPoint struct {
	Trx_Type_Id  int
	Reward_Id    int
	Ss_Point     int
	Own_Pct      int
	Referral_Pct int
	Branch_Pct   int
	Created_at   time.Time
	Created_by   int
}

type UpdateTransactionPoint struct {
	Trx_Type_Id  int
	Reward_Id    int
	Ss_Point     int
	Own_Pct      int
	Referral_Pct int
	Branch_Pct   int
	Created_at   time.Time
	Updated_at   time.Time
	Updated_by   int
}

// type StoreTransactionPointHistory struct {
// 	Trx_Type_Id  int
// 	Reward_Id    int
// 	Ss_Point     int
// 	Own_Pct      int
// 	Referral_Pct int
// 	Branch_Pct   int
// 	Created_at   time.Time
// 	Created_by   int
// }

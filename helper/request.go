package helper

import "time"

type TrxPointRequest struct {
	FromTrxDate string `json:"fromTrxDate" form:"fromTrxDate" query:"fromTrxDate"`
	ToTrxDate   string `json:"toTrxDate" form:"toTrxDate" query:"toTrxDate"`
}

type TrxPointData struct {
	Trx_date        time.Time
	Type_name       string
	Ss_point_before int
	Ss_point_trx    int
	Ss_point_after  int
}

type StoreTrxPointReq struct {
	Trx_date         string `json:"trxDate" form:"trxDate"`
	User_id          int    `json:"userId" form:"userId"`
	Is_branch        int    `json:"isBranch" form:"isBranch"`
	Trx_point_id     int    `json:"trxPointId" form:"trxPointId"`
	Reference_trx_id int    `json:"referenceTrxId" form:"referenceTrxId"`
	Ss_point_before  int    `json:"ssPointBefore" form:"ssPointBefore"`
	Ss_point_trx     int    `json:"ssPointTrx" form:"ssPointTrx"`
	Ss_point_after   int    `json:"ssPointAfter" form:"ssPointAfter"`
}

type UserId struct {
	User_id int `json:"userId"`
}

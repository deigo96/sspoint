package helper

import "time"

type GetData struct {
	TotalAllData int
	Summary      Point
	Result       []Data
}

type Point struct {
	TotalPointBefore       int
	TotalPointAddition     int
	TotalPointSubstraction int
	TotalPointAfter        int
}
type Data struct {
	Trx_date        time.Time
	Type_name       string
	Ss_point_before int
	Ss_point_trx    int
	Ss_point_after  int
}

func AllPointHisList(g GetData) GetData {
	return GetData{
		TotalAllData: g.TotalAllData,
		Summary:      g.Summary,
		Result:       AllPointHisSlice(g.Result),
	}
}

func AllPointHisResult(d Data) Data {
	return Data{
		Trx_date:        d.Trx_date,
		Type_name:       d.Type_name,
		Ss_point_before: d.Ss_point_before,
		Ss_point_trx:    d.Ss_point_trx,
		Ss_point_after:  d.Ss_point_after,
	}
}

func AllPointHisSlice(data []Data) []Data {
	var resultArray []Data
	for key := range data {
		resultArray = append(resultArray, AllPointHisResult(data[key]))
	}
	return resultArray
}

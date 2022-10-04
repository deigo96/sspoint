package helper

import (
	"time"
)

type GetData struct {
	TotalAllData int    `json:"TotalAllData"`
	Result       []Info `json:"Result"`
}

type Info struct {
	Upline   []Upline   `json:"Upline"`
	Downline []Downline `json:"Downline"`
}

type Upline struct {
	TotalAllData int      `json:"TotalAllData"`
	Results      []Result `json:"Result"`
}

type Downline struct {
	TotalAllData int      `json:"TotalAllData"`
	Results      []Result `json:"Result"`
}

type Result struct {
	Parent_id int    `json:"userId"`
	Name      string `json:"name"`
}

type ReqId struct {
	Parent_id int
	Child_id  int
}

type Update struct {
	Parent_id    int `json:"parent_id"`
	Child_id     int `json:"child_id"`
	To_Parent_id int `json:"to_parent_id"`
}

type Data struct {
	Parent_id int `json:"parent_id"`
	Child_id  int `json:"child_id"`
}

type Postdata struct {
	Parent_id  int
	Child_id   int
	Created_at time.Time
	Created_by int
}

type Posthistory struct {
	Parent_id  int
	Child_id   int
	Created_at time.Time
	Created_by int
}

type UpdateData struct {
	Parent_id  int
	Child_id   int
	Updated_at time.Time
	Updated_by int
}

func AllUserReferral(g GetData) GetData {
	return GetData{
		TotalAllData: g.TotalAllData,
		Result:       AllUserInfoSlice(g.Result),
	}
}

func AllUserInfo(i Info) Info {
	return Info{
		Upline:   AllUserUplineSlice(i.Upline),
		Downline: AllUserDownlineSlice(i.Downline),
	}
}

func AllUserInfoSlice(i []Info) []Info {
	var infoArray []Info
	for key := range i {
		infoArray = append(infoArray, AllUserInfo(i[key]))
	}
	return infoArray
}

func AllUserUpline(i Upline) Upline {
	return Upline{
		TotalAllData: i.TotalAllData,
		Results:      AllUserReferralSlice(i.Results),
	}
}

func AllUserUplineSlice(i []Upline) []Upline {
	var uplineArray []Upline
	for key := range i {
		uplineArray = append(uplineArray, AllUserUpline(i[key]))
	}

	return uplineArray
}

func AllUserDownline(i Downline) Downline {
	return Downline{
		TotalAllData: i.TotalAllData,
		Results:      AllUserReferralSlice(i.Results),
	}
}

func AllUserDownlineSlice(i []Downline) []Downline {
	var downArray []Downline
	for key := range i {
		downArray = append(downArray, AllUserDownline(i[key]))
	}

	return downArray
}

func AllUserResult(r Result) Result {
	return Result{
		Parent_id: r.Parent_id,
		Name:      r.Name,
	}
}

func AllUserReferralSlice(data []Result) []Result {
	var resultArray []Result
	for key := range data {
		resultArray = append(resultArray, AllUserResult(data[key]))
	}

	return resultArray
}

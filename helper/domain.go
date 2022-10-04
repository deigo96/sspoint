package helper

import "time"

type GetData struct {
	TotalAllData int
	Result       []Data
}

type Data struct {
	Id          int    `json:"RewardID"`
	Reward_Name string `json:"RewardName"`
	Description string `json:"RewardDescription"`
	Img_Name    string `json:"RewardImageUrl"`
	// RewardPoint       int    `json:""`
}

type StoreData struct {
	Reward_Name string
	Img_Name    string
	Description string
	Created_By  int
	Created_At  time.Time
}

type UpdateData struct {
	Reward_Name string
	Img_Name    string
	Description string
	Updated_By  int
	Updated_At  time.Time
}

func AllRewardList(g GetData) GetData {
	return GetData{
		TotalAllData: g.TotalAllData,
		Result:       AllRewardSlice(g.Result),
	}
}

func AllRewardResult(d Data) Data {
	return Data{
		Id:          d.Id,
		Reward_Name: d.Reward_Name,
		Img_Name:    d.Img_Name,
		Description: d.Description,
	}
}

func AllRewardSlice(data []Data) []Data {
	var resultArray []Data
	for key := range data {
		resultArray = append(resultArray, AllRewardResult(data[key]))
	}
	return resultArray
}

package helper

type StoreRewardRequest struct {
	Reward_Name        string `json:"rewardName" form:"rewardName"`
	Reward_Image       string `json:"rewardImage" form:"rewardImage"`
	Reward_Description string `json:"rewardDescription" form:"rewardDescription"`
}

type GetRewardRequst struct {
	Id                 int
	Reward_Name        string
	Reward_Image       string
	Reward_Description string
}

type Idreward struct {
	Id_Reward int `json:"idReward" form:"idReward"`
}

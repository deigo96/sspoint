package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elgs/gojq"
)

type RewardList interface {
	GetRewardList() (d []Data)
	StoreRewardList(id int, s StoreRewardRequest) error
	UpdateRewardList(id int, idReward int, s StoreRewardRequest) error
}

type RewardService interface {
	GetReward() (g GetData, err error)
	StoreReward(id int, s StoreRewardRequest) (*StoreRewardRequest, error)
	UpdateReward(id int, idReward int, s StoreRewardRequest) (*Idreward, error)
}

type rewardService struct {
	reward RewardList
}

func NewRewardService(reward RewardList) RewardService {
	return &rewardService{
		reward: reward,
	}
}

func (r *rewardService) GetReward() (res GetData, err error) {
	result := []Data{}
	data := r.reward.GetRewardList()
	for _, i := range data {
		result = append(result, i)
	}

	res.TotalAllData = len(result)
	res.Result = result

	return res, nil
}

func (r *rewardService) StoreReward(id int, s StoreRewardRequest) (*StoreRewardRequest, error) {
	err := r.reward.StoreRewardList(id, s)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *rewardService) UpdateReward(id int, idReward int, s StoreRewardRequest) (*Idreward, error) {
	err := r.reward.UpdateRewardList(id, idReward, s)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func GetIdToken(token string) (id int, err error) {
	url := fmt.Sprintf("http://192.168.97.121:8000/seen/user/profile")
	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	str := string(body)
	p := strings.TrimSuffix(str, "\n")

	parser, err := gojq.NewStringQuery(p)
	if err != nil {
		fmt.Println(parser)
	}
	d, _ := parser.QueryToInt64("message.data.id")
	id = int(d)

	return id, nil

}

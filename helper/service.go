package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elgs/gojq"
)

type PointHistoryList interface {
	GetPointHistoryList(id int, t TrxPointRequest) (d []Data)
	StorePointHistoryList(s StoreTrxPointReq) error
}

type PointHistoryService interface {
	GetPointHistoryService(id int, t TrxPointRequest) (g GetData, err error)
	StorePointHistoryService(s StoreTrxPointReq) error
}

type pointHistoryService struct {
	point PointHistoryList
}

func NewRewardService(point PointHistoryList) PointHistoryService {
	return &pointHistoryService{
		point: point,
	}
}

func (p *pointHistoryService) GetPointHistoryService(id int, t TrxPointRequest) (res GetData, err error) {
	result := []Data{}
	// summary := map[int]interface{}{}
	data := p.point.GetPointHistoryList(id, t)
	totalBefore := 0
	totalAfter := 0
	for _, i := range data {
		result = append(result, i)
		totalAfter = i.Ss_point_after + totalAfter
		totalBefore = i.Ss_point_before + totalBefore
	}

	s := Point{
		TotalPointBefore:       totalBefore,
		TotalPointAddition:     102,
		TotalPointSubstraction: 10,
		TotalPointAfter:        totalAfter,
	}

	res.Summary = s
	res.TotalAllData = len(result)
	res.Result = result

	return res, nil
}

func (p *pointHistoryService) StorePointHistoryService(s StoreTrxPointReq) error {
	err := p.point.StorePointHistoryList(s)
	if err != nil {
		return err
	}

	return nil
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

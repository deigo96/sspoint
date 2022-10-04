package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elgs/gojq"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	GetReferral(id int, token string) (d []Result)
	Register(Parent_id int, Child_id int, idToken int) error
	UpdateReferral(Parent_id int, Child_id int, To_Parent_id int, idToken int) error
}

type RefferralService interface {
	GetRef(id int, token string) (g GetData, err error)
	RegisterRef(Parent_id int, Child_id int, idToken int) (*Data, error)
	UpdateRef(Parent_id int, Child_id int, To_Parent_id int, idToken int) (*Update, error)
}

type referralService struct {
	referral AuthService
}

func NewRereferralService(referral AuthService) RefferralService {
	return &referralService{
		referral: referral,
	}
}

type Echo interface {
	echo.Context
}

func (r *referralService) GetRef(id int, token string) (res GetData, err error) {
	info := []Info{}
	upline := []Upline{}
	downline := []Downline{}
	domain := []Result{}

	data := r.referral.GetReferral(id, token)
	dataParent, _ := GetIdToken(token)

	for _, d := range data {
		domain = append(domain, d)
	}

	u := Upline{
		TotalAllData: dataParent.TotalAllData,
		Results:      dataParent.Results,
	}
	d := Downline{
		TotalAllData: len(domain),
		Results:      domain,
	}
	upline = append(upline, u)
	downline = append(downline, d)
	i := Info{
		Upline:   upline,
		Downline: downline,
	}

	info = append(info, i)

	res.TotalAllData = u.TotalAllData + d.TotalAllData
	res.Result = info

	return res, nil
}

func (r *referralService) RegisterRef(Parent_id int, Child_id int, idToken int) (*Data, error) {
	err := r.referral.Register(Parent_id, Child_id, idToken)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *referralService) UpdateRef(Parent_id int, Child_id int, To_Parent_id int, idToken int) (*Update, error) {
	err := r.referral.UpdateReferral(Parent_id, Child_id, To_Parent_id, idToken)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetIdToken(token string) (res Upline, err error) {
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
	name, _ := parser.QueryToString("message.data.username")

	result := []Result{}

	r := Result{
		Parent_id: int(d),
		Name:      name,
	}

	result = append(result, r)

	res.TotalAllData = len(result)
	res.Results = result

	return res, nil
}

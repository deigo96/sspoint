package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"referralUser-service/helper"
	"strings"
	"time"

	"github.com/elgs/gojq"
	"gorm.io/gorm"
)

type ReferralRepo struct {
	db *gorm.DB
}

func NewRegisterReferral(db *gorm.DB) helper.AuthService {
	return &ReferralRepo{
		db: db,
	}
}

type User struct {
	Username string `json:"name"`
}

func a() int {
	return 1
}

func (r *ReferralRepo) GetReferral(id int, token string) (d []helper.Result) {
	var resp []helper.ReqId
	res := r.db.Table("referral_tree").Where("created_by = ?", id).Find(&resp)

	if res.Error != nil {
		return d
	}

	for _, arg := range resp {
		url := fmt.Sprintf("http://192.168.97.121:8000/seen/user/detail/%d", arg.Child_id)
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
		name, _ := parser.QueryToString("message.data.username")
		idUser, _ := parser.QueryToInt64("message.data.id")
		// fmt.Println(idUser)
		data := helper.Result{
			Parent_id: int(idUser),
			Name:      name,
		}
		d = append(d, data)
	}

	return d

}

func (r *ReferralRepo) Register(Parent_id, Child_id, idToken int) error {
	record := helper.Postdata{
		Parent_id:  Parent_id,
		Child_id:   Child_id,
		Created_at: time.Now(),
		Created_by: idToken,
	}
	his := helper.Posthistory{
		Parent_id:  record.Parent_id,
		Child_id:   record.Child_id,
		Created_at: time.Now(),
		Created_by: record.Created_by,
	}

	res := r.db.Table("referral_tree").Create(&record)
	resHis := r.db.Table("referral_tree_hist").Create(&his)
	if res.Error != nil {
		return res.Error
	}
	if resHis.Error != nil {
		return res.Error
	}

	return nil
}

func (r *ReferralRepo) UpdateReferral(Parent_id, Child_id, To_Parent_id int, idToken int) error {
	record := helper.UpdateData{
		Parent_id:  To_Parent_id,
		Child_id:   Child_id,
		Updated_at: time.Now(),
		Updated_by: idToken,
	}
	his := helper.Posthistory{
		Parent_id:  record.Parent_id,
		Child_id:   record.Child_id,
		Created_at: time.Now(),
		Created_by: idToken,
	}

	res := r.db.Table("referral_tree").Where("parent_id = ? AND child_id = ?", Parent_id, Child_id).Updates(&record)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("No record found")
	}

	r.db.Table("referral_tree_hist").Create(&his)
	return nil
}

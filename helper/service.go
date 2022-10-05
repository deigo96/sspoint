package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elgs/gojq"
)

type TransactionPointList interface {
	StoreTransactionPointList(id int, s TransactionPoint) error
	UpdateTransactionPointList(id int, idTransaction int, s TransactionPoint) error
}

type TransactionPointService interface {
	StoreTransactionPointService(id int, s TransactionPoint) (*TransactionPoint, error)
	UpdateTransactionPointService(id int, idTransaction int, s TransactionPoint) (*IdTransactionPoint, error)
}

type transactionService struct {
	transaction TransactionPointList
}

func NewTransactionPointService(transaction TransactionPointList) TransactionPointService {
	return &transactionService{
		transaction: transaction,
	}
}

func (t *transactionService) StoreTransactionPointService(id int, s TransactionPoint) (*TransactionPoint, error) {
	err := t.transaction.StoreTransactionPointList(id, s)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *transactionService) UpdateTransactionPointService(id int, idTransactionPoint int, s TransactionPoint) (*IdTransactionPoint, error) {
	err := t.transaction.UpdateTransactionPointList(id, idTransactionPoint, s)
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

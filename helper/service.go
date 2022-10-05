package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elgs/gojq"
)

type TransactionTypeList interface {
	StoreTransactionTypeList(id int, s TransactionType) error
	UpdateTransactionTypeList(id int, idTransaction int, s TransactionType) error
}

type TransactionTypeService interface {
	StoreTransactionTypeService(id int, s TransactionType) (*TransactionType, error)
	UpdateTransactionTypeService(id int, idTransaction int, s TransactionType) (*IdTransactionType, error)
}

type transactionService struct {
	transaction TransactionTypeList
}

func NewTransactionTypeService(transaction TransactionTypeList) TransactionTypeService {
	return &transactionService{
		transaction: transaction,
	}
}

func (t *transactionService) StoreTransactionTypeService(id int, s TransactionType) (*TransactionType, error) {
	err := t.transaction.StoreTransactionTypeList(id, s)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *transactionService) UpdateTransactionTypeService(id int, idTransactionType int, s TransactionType) (*IdTransactionType, error) {
	err := t.transaction.UpdateTransactionTypeList(id, idTransactionType, s)
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

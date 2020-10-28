package singleToken

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/address"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/config"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/logger"
	"io/ioutil"
	"net/http"
)

type SingleTokenReq struct {
	MerchantRefNum string   `json:"merchantRefNum"`
	PaymentTypes   []string `json:"paymentTypes"`
}

type SingleTokenRes struct {
	Id                     string            `json:"id"`
	MerchantRefNum         string            `json:"merchantRefNum"`
	CustomerId             string            `json:"customerId"`
	TimeToLiveSeconds      int               `json:"timeToLiveSeconds"`
	Status                 string            `json:"status"`
	SingleUseCustomerToken string            `json:"singleUseCustomerToken"`
	PaymentTypes           []string          `json:"paymentTypes"`
	Locale                 string            `json:"locale"`
	FirstName              string            `json:"firstName"`
	MiddleName             string            `json:"middleName"`
	LastName               string            `json:"lastName"`
	DateOfBirth            DateOfBirth       `json:"dateOfBirth"`
	Email                  string            `json:"email"`
	Phone                  string            `json:"phone"`
	Ip                     string            `json:"ip"`
	Nationality            string            `json:"nationality"`
	Addresses              []address.Address `json:"addresses"`
}

type DateOfBirth struct {
	Year  int  `json:"year"`
	Month int  `json:"month"`
	Day   int  `json:"day"`
}

func CreateSingleUseToken(s *SingleTokenReq, cid string) (*SingleTokenRes, error) {
	client := http.Client{}
	b, err := json.Marshal(s)
	if err != nil {
		logger.ErrorLogger.Printf("error in marshalling to create single use token, error:%v", err)
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.test.paysafe.com/paymenthub/v1/customers/%s/singleusecustomertokens", cid), bytes.NewBuffer(b))
	if err != nil {
		logger.ErrorLogger.Printf("error in creating a new request to create single use token, error:%v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s", config.C.Common.ApiSecret))
	req.Header.Set("Simulator", "EXTERNAL")

	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorLogger.Printf("error in calling api to create single use token, error:%v", err)
		return nil, err
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLogger.Printf("error in reading response to create single use token, resp:%v, error:%v", string(b), err)
		return nil, err
	}
	st := &SingleTokenRes{}
	err = json.Unmarshal(b, st)
	if err != nil {
		logger.ErrorLogger.Printf("error in UnMarshalling response to create single use token, resp:%v error:%v", resp, err)
		return nil, err
	}
	logger.InfoLogger.Printf("CreateSingleUseToken, req:%v, resp:%v\n", s, st)
	logger.InfoLogger.Printf("CreateSingleUseToken byteResp:%v, rawResp:%v\n", string(b), resp)
	return st, nil
}
package pay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/config"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/logger"
	"io/ioutil"
	"net/http"
)

type PayReq struct {
	MerchantRefNum     string `json:"merchantRefNum"`
	Amount             int    `json:"amount"`
	CurrencyCode       string `json:"currencyCode"`
	DupCheck           bool   `json:"dupCheck"`
	SettleWithAuth     bool   `json:"settleWithAuth"`
	PaymentHandleToken string `json:"paymentHandleToken"`
	CustomerIp         string `json:"customerIp"`
	Description        string `json:"description"`
}

type PayResp struct {
	Id                 string          `json:"id"`
	Amount             int             `json:"amount"`
	MerchantRefNum     string          `json:"merchantRefNum"`
	SettleWithAuth     bool            `json:"settleWithAuth"`
	PaymentHandleToken string          `json:"paymentHandleToken"`
	TxnTime            string          `json:"txnTime"`
	DupCheck           bool            `json:"dupCheck"`
	Description        string          `json:"description"`
	CurrencyCode       string          `json:"currencyCode"`
	PaymentType        string          `json:"paymentType"`
	Status             string          `json:"status"`
	AvailableToSettle  int             `json:"availableToSettle"`
	GatewayResponse    GatewayResponse `json:"gatewayResponse"`
}

type GatewayResponse struct {
	AuthCode        string `json:"authCode"`
	AvsResponse     string `json:"avsResponse"`
	CvvVerification string `json:"cvvVerification"`
}

func ProcessPay(p *PayReq) (*PayResp, error) {
	client := http.Client{}
	b, err := json.Marshal(p)
	if err != nil {
		logger.ErrorLogger.Printf("error in marshalling for processing payment, error:%v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.test.paysafe.com/paymenthub/v1/payments"), bytes.NewBuffer(b))
	if err != nil {
		logger.ErrorLogger.Printf("error in creating new request to processing payment, error:%v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.C.Common.ApiSecret)
	req.Header.Set("Simulator", "EXTERNAL")
	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorLogger.Printf("error in api call to processing payment, error:%v", err)
		return nil, err
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLogger.Printf("error in reading the response to processing payment resp:%v error:%v", string(b), err)
		return nil, err
	}
	createProfileResp := &PayResp{}
	err = json.Unmarshal(b, createProfileResp)
	if err != nil {
		logger.ErrorLogger.Printf("error in UnMarshalling response to processing payment, error:%v", err)
		return nil, err
	}

	logger.InfoLogger.Printf("Process Pay Service, req:%+v, resp:%+v\n", p, createProfileResp)
	logger.InfoLogger.Printf("Process Pay Service byteResp:%v, rawResp:%v\n", string(b), resp)
	return createProfileResp, err
}

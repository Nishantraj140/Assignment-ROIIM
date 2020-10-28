package address

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Nishantraj140/Assignment-ROIIM/internal/config"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/logger"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/sql"
)

type Address struct {
	Id                              string `json:"id"`
	Status                          string `json:"status"`
	NickName                        string `json:"nickName"`
	Street                          string `json:"street"`
	Street2                         string `json:"street2"`
	City                            string `json:"city"`
	Zip                             string `json:"zip"`
	Country                         string `json:"country"`
	State                           string `json:"state"`
	Phone                           string `json:"phone"`
	DefaultShippingAddressIndicator bool `json:"defaultShippingAddressIndicator"`
}

func (a *Address) Get() (err error) {
	return sql.DB.Model(a).First(a).Error
}

func (a *Address) Create() (err error) {
	return sql.DB.Model(a).Create(a).Error
}

func CreateAddressService(a Address, pid string) (*Address, error) {
	client := http.Client{}
	b, err := json.Marshal(a)
	if err != nil {
		logger.ErrorLogger.Printf("error in marshalling for create address request, error:%v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.test.paysafe.com/paymenthub/v1/customers/%v/addresses", pid), bytes.NewBuffer(b))
	if err != nil {
		logger.ErrorLogger.Printf("error in creating new request for create address request, error:%v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.C.Common.ApiSecret)
	req.Header.Set("Simulator", "EXTERNAL")
	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorLogger.Printf("error in api call for create address request, error:%v", err)
		return nil, err
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLogger.Printf("error in reading response for create address request, resp:%v, error:%v", string(b), err)
		return nil, err
	}
	addressResp := &Address{}
	err = json.Unmarshal(b, addressResp)
	if err != nil {
		logger.ErrorLogger.Printf("error in UnMarshalling response for create address request, resp%v, error:%v", string(b), err)
		return nil, err
	}
	logger.InfoLogger.Printf("Create Address Service, req:%+v, resp:%+v\n", a, addressResp)
	logger.InfoLogger.Printf("Create Profile Service byteResp:%v, rawResp:%v\n", string(b), resp)
	return addressResp, err
}

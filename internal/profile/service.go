package profile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Nishantraj140/Assignment-ROIIM/internal/config"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/logger"
)

func CreateProfileService(p CreateProfile) (*CreateProfileRes, error) {
	client := http.Client{}
	b, err := json.Marshal(p)
	if err != nil {
		logger.ErrorLogger.Printf("error in marshalling create profile request, error:%v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.test.paysafe.com/paymenthub/v1/customers"), bytes.NewBuffer(b))
	if err != nil {
		logger.ErrorLogger.Printf("error in creating new request to create profile, error:%v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.C.Common.ApiSecret)
	req.Header.Set("Simulator", "EXTERNAL")
	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorLogger.Printf("error in api call to create profile, error:%v", err)
		return nil, err
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLogger.Printf("error in reading the response to create profile, resp:%v error:%v", string(b), err)
		return nil, err
	}
	createProfileResp := &CreateProfileRes{}
	err = json.Unmarshal(b, createProfileResp)
	if err != nil {
		logger.ErrorLogger.Printf("error in UnMarshalling response to create profile, error:%v", err)
		return nil, err
	}

	logger.InfoLogger.Printf("Create Profile Service, req:%+v, resp:%+v\n", p, createProfileResp)
	logger.InfoLogger.Printf("Create Profile Service byteResp:%v, rawResp:%v\n", string(b), resp)
	return createProfileResp, err
}

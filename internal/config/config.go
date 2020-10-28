package config

import (
	"encoding/json"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/sql"
	"io/ioutil"
	"log"
	"os"
)

var C *Config

type Config struct {
	Common CommonConfig `json:"common"`
	DBConfig sql.DBConfig `json:"db"`
}

type CommonConfig struct {
	ApiSecret string `json:"api_secret"`
	AccountId string `json:"account_id"`
	ServerPort string `json:"server_port"`
}

func ReadConfig(configFile string) {
	C = &Config{}
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &C); err != nil {
		log.Fatalf("Unable to marshal config data")
		return
	}
	log.Println("config loaded", C)
}


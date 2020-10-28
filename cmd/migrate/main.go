package main

import (
	"flag"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/address"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/config"
	"github.com/Nishantraj140/Assignment-ROIIM/internal/user"
	"github.com/Nishantraj140/Assignment-ROIIM/pkg/sql"
	"log"
)

var configFile = flag.String("config", "conf/config.json","config file")

func main() {
	flag.Parse()
	config.ReadConfig(*configFile)
	err := sql.DBConn(&config.C.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer sql.DB.Close()
	sql.DB.AutoMigrate(&user.User{},&user.UserCard{},&address.Address{})
}

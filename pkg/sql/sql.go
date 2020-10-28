package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type DBConfig struct {
	Url string `json:"url"`
	UserName string `json:"username"`
	Password string `json:"password"`
	DataBase string `json:"database"`
}

func DBConn(c *DBConfig) (err error) {
	url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.UserName, c.Password, c.Url, c.DataBase)
	DB, err = gorm.Open("mysql", url)
	DB = DB.LogMode(true)
	return
}



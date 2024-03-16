package database

import (
	"chatbot/internal/config"
	"chatbot/internal/config/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbConfigs map[string]dbConfig

type dbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
}

var DefaultPageSize = 20

var dbs map[string]*gorm.DB = make(map[string]*gorm.DB, 0)

func init() {
	config.LoadConfig()
	var dbConfigs dbConfigs
	cfg := global.Cfg
	cfg.UnmarshalKey("mysql", &dbConfigs)

	for k, v := range dbConfigs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", v.User, v.Password, v.Addr, v.Database, v.Charset)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed to initialize database: %v", err)
		}
		dbs[k] = db
	}

}

func GetDB(biz string) *gorm.DB {
	return dbs[biz]
}

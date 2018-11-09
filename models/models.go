package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gavinzhou/hello-gin/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                     error
		dbType, dbName, user, host, tablePrefix string
	)

	dbType = setting.Config.DBType
	dbName = setting.Config.DBName
	user = setting.Config.DBUser
	host = setting.Config.DBHost
	tablePrefix = setting.Config.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable",
		host,
		user,
		dbName,
	))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

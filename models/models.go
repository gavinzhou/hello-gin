package models

import (
	"log"

	"github.com/go-pg/pg"

	"github.com/gavinzhou/hello-gin/pkg/setting"
)

var db *pg.DB

type Model struct {
	ID         int64
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func ModelDB() *pg.DB {
	cfg, err := setting.LoadConfig()

	if err != nil {
		log.Fatal(2, "Fail to load config with database: %v", err)
	}

	db := pg.Connect(&pg.Options{
		User:     cfg.Username,
		Database: cfg.Database,
	})
	return db

}

func CloseDB() {
	defer db.Close()
}

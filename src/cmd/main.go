package main

import (
	"github.com/miladshalikar/golang-clean-web-api/src/api"
	"github.com/miladshalikar/golang-clean-web-api/src/config"
	"github.com/miladshalikar/golang-clean-web-api/src/data/cache"
	"github.com/miladshalikar/golang-clean-web-api/src/data/db"
	"log"
)

func main() {

	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}

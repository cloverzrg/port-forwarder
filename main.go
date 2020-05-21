package main

import (
	"fmt"
	"github.com/cloverzrg/go-portforward/config"
	"github.com/cloverzrg/go-portforward/db"
	"github.com/cloverzrg/go-portforward/logger"
	"github.com/cloverzrg/go-portforward/model"
	"github.com/cloverzrg/go-portforward/web"
)

// @title go-portforward
// @version 1.0

// @contact.name API Support
// @contact.url https://github.com/cloverzrg/go-portforward
// @contact.email cloverzrg@gmail.com

// @license.name go-portforward

var (
	BuildTime string
	GoVersion string
	GitHead   string
)

func main() {
	var err error
	err = web.Start()
	if err != nil {
		logger.Panic(err)
	}
}

func init() {
	var err error
	fmt.Printf("BuildTime: %s\nGoVersion: %s\nGitHead: %s\n", BuildTime, GoVersion, GitHead)
	config.Parse("./data/config.json")
	err = db.Connect()
	if err != nil {
		logger.Panic(err)
	}
	err = model.CreateAllTable()
	if err != nil {
		logger.Panic(err)
	}
}

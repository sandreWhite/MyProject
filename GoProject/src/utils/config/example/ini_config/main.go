package main

import (
	"encoding/json"
	"utils/config"
	"utils/logger"
	_ "utils/config/ini"
)

type mongoDB struct {
	Host string `json:"host" ini:"host" cfg:"host" cfgDefault:"example.com"`
	Port int    `json:"port" ini:"port" cfg:"port" cfgDefault:"999"`
}

type systemUser struct {
	Name     string `json:"name" ini:"name" cfg:"name"`
	Password string `json:"passwd" ini:"passwd" cfg:"passwd"`
}

type configTest struct {
	DebugMode bool `json:"debug" ini:"debug" cfg:"debug" cfgDefault:"false"`
	Domain    string
	User      systemUser `json:"user" ini:"user" cfg:"user"`
	MongoDB   mongoDB    `json:"mongodb" ini:"mongodb" cfg:"mongodb"`
}

func main() {
	var log = logger.New()	
	log.Formatter = new(logger.JSONFormatter) 
	log.Formatter.(*logger.JSONFormatter).PrettyPrint = true
	log.Formatter.(*logger.JSONFormatter).DisableTimestamp = false
	log.Level = logger.TraceLevel	
	config := configTest{}

	goconfig.File = "config.ini"
	err := goconfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	// just print struct on screen
	j, _ := json.MarshalIndent(config, "", "\t")
	log.WithFields(logger.Fields{
			"debug":			config.DebugMode,
			"domain":			config.Domain,
			"username":			config.User.Name,
			"mongodb_host":		config.MongoDB.Host,
			"mongodb_porr":		config.MongoDB.Port,
	}).Info("config.ini")
	
	println(string(j))
}
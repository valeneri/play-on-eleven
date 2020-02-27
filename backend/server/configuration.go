package server

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	port         = ":8080"
	databaseType = "mongodb"
	databaseUrl  = "mongodb://play-on-eleven-mongo"
	databaseName = "play-dev"
	username     = "valou5940"
	password     = "5ekr3t@"
	poolSize     = 35
	timeout      = 10
)

type Db struct {
	DatabaseType string `json:"database_type"`
	DatabaseUrl  string `json:"database_url"`
	DatabaseName string `json:"database_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PoolSize     uint64 `json:"pool_size"`
	Timeout      int    `json:"timeout"`
}

type Server struct {
	Port string `json:"port"`
}

type Configuration struct {
	Db     Db     `json:"Db`
	Server Server `json:"Server`
}

func SetConfig() Configuration {
	var configuration Configuration
	cfg, err := readConfigFile()
	if err != nil {
		configuration = setDefaultConfig()
	} else {
		configuration = cfg
	}
	return configuration
}

func readConfigFile() (Configuration, error) {
	configuration := Configuration{}
	file, err := os.Open("../config/config-dev.json")
	if err != nil {
		fmt.Println("error opening config file")
		return configuration, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return configuration, err
	}
	fmt.Println("Setting configuration from file")
	fmt.Println(configuration)

	return configuration, nil
}

func setDefaultConfig() Configuration {
	configuration := Configuration{}
	db := Db{
		DatabaseType: databaseType,
		DatabaseUrl:  databaseUrl,
		DatabaseName: databaseName,
		Username:     username,
		Password:     password,
		PoolSize:     poolSize,
		Timeout:      timeout,
	}
	server := Server{
		Port: port,
	}
	configuration.Db = db
	configuration.Server = server
	fmt.Println("Setting default configuration")
	return configuration
}

// func setConfigFromFile(file *os.File) Configuration {
// 	configuration := Configuration{}
// 	decoder := json.NewDecoder(file)
// 	err := decoder.Decode(&configuration)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 		setDefaultConfig()
// 	}
// 	fmt.Println("Setting configuration from file")
// 	return configuration
// }

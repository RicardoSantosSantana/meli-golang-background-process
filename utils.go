package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/RicardoSantosSantana/apiTokenControl"
	"github.com/RicardoSantosSantana/dbTokenControl"
)

func checkErr(err error) {

	if err != nil {
		fmt.Println(err.Error())
		dbTokenControl.LogMessage(err.Error())
		panic(err)
	}
}

// get variable content if exist
func environment(value_default string, env string) string {

	enviromentVariable := strings.Trim(os.Getenv(strings.ToUpper(env)), " ")

	if enviromentVariable != "" {
		return os.Getenv(env)
	}
	return value_default
}

func StringConnection() dbTokenControl.STRConn {

	sconn := dbTokenControl.STRConn{
		DbName:     environment("dbmeli", "MYSQL_DATABASE"),
		DbHost:     environment("127.0.0.1", "MYSQL_SERVER"),
		DbUser:     environment("meli", "MYSQL_USER"),
		DbPassword: environment("meli123", "MYSQL_PASSWORD"),
		DbPort:     environment("3306", "MYSQL_PORT"),
	}
	return sconn
}

func InitialParams() apiTokenControl.SInitialParams {

	conf := apiTokenControl.SInitialParams{
		Client_id:     environment("2200107891711741", "CLIENT_ID"),
		Client_secret: environment("FJATnQFyHWIjcwaJg0e2eM9Y9z9ByXEH", "CLIENT_SECRET"),
		Code:          environment("TG-62fd2ad21eca6e0001222ee1-1176272578", "CODE"),
		Redirect_uri:  environment("https://reciboonline.com", "REDIRECT_URL"),
	}
	return conf
}

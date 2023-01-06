package main

import (
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
)

func init(){
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main(){
	config.DB.AutoMigrate(
		&models.User{},
		&models.Memo{},
	)
}
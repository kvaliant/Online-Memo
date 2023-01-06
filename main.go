package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/middlewares"
	"github.com/kvaliant/online-memo/routes"
)

func init(){
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main(){
	r := gin.Default()
	r.GET("/ping", routes.Ping)
	
	r.POST("/login", routes.Login)
	r.POST("/register", routes.Register)
	r.POST("/logout", routes.Logout)

	authorized := r.Group("/")
	authorized.Use(middlewares.JWTMiddleware())
	{
		authorized.GET("/auth_ping", routes.AuthPing)
		authorized.GET("/memos", routes.GetMemos)
		authorized.POST("/memos", routes.PostMemo)
		authorized.GET("/memo/:memoid", routes.GetMemo)
		authorized.PATCH("/memo/:memoid", routes.PatchMemo)
		authorized.DELETE("/memo/:memoid", routes.DeleteMemo)
	}
	

	r.Run() // 0.0.0.0:8080
}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/helloteemo/wine"
	"log"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"username"`
}

func main() {
	g := gin.Default()
	g.GET("/error", wine.Wine(Error))
	g.POST("/print", wine.Wine(Print))
	g.Run(":9898")
}

func Error(c *gin.Context, req *User, req2 User) interface{} {
	log.Printf("req:%+v", req)
	log.Printf("req2:%+v", req2)
	return wine.SystemError
}

func Print(c *gin.Context, req User, req2 *User) interface{} {
	log.Printf("req:%+v", req)
	log.Printf("req2:%+v", req2)
	return User{"alice", "123456"}
}

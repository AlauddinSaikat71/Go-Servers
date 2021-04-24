package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rahmanfadhil/gin-bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.Run()
}

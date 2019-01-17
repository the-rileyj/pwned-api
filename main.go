package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-rileyj/pwned-api/functionality"
)

func main() {
	router := gin.Default()

	apiGroup := router.Group("/api")

	apiGroup.POST("/notify-pwnage", functionality.NotifyOfPwnage)

	apiGroup.POST("/notify-pwnage-without-cache", functionality.NotifyOfPwnage)

	apiGroup.POST("/notify-pwnage-with-cache", functionality.NotifyOfPwnage)

	apiGroup.POST("/add-to-pwnage-check", functionality.AddToPwnageCheck)

	apiGroup.POST("/delete-from-pwnage-check", functionality.DeleteFromPwnageCheck)

	router.Run(":80")
}

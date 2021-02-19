package main

import "github.com/gin-gonic/gin"
func main() {
	startServer()
}

func startServer() {
	router := gin.Default()
	router.GET("/maximizeFont", maximizeFont)
	router.Run(":9091")
}


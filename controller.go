package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func maximizeFont(c *gin.Context) {
	//reqObject := MaximizeFontRequest{}
	//if err := c.BindJSON(&reqObject); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//resp,err := MaximizeFont(&reqObject)
	fontFamily := c.Query("fontFamily")
	text := c.Query("text")
	boxWidth,_ := strconv.Atoi(c.Query("boxWidth"))
	boxHeight,_ := strconv.Atoi(c.Query("boxHeight"))
	resp,err := MaximizeFont(&MaximizeFontRequest{
		fontFamily: fontFamily,
		text:       text,
		boxWidth:   boxWidth,
		boxHeight:  boxHeight,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"status": "ERROR",
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H {
		"status": "SUCCESS",
		"text": resp.text,
		"fontSize": resp.fontSize,
	})
}
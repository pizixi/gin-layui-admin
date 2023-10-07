package router

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getClientIP(c *gin.Context) string {
	s := strings.Split(c.Request.RemoteAddr, ":")
	return s[0]
}

func getUserID(c *gin.Context) int {
	if v, ok := c.Get("userId"); ok {
		userID, _ := strconv.Atoi(v.(string))
		return userID
	}
	return 0
}

func getUserName(c *gin.Context) string {
	if v, ok := c.Get("userName"); ok {
		name := v.(string)
		return name
	}
	return ""
}

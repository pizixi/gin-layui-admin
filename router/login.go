package router

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gin-layui-admin/libs"
	"gin-layui-admin/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SiteName SiteName
const SiteName = "gin-layui后台 v1.0"

// LoginIndex LoginIndex
func LoginIndex(c *gin.Context) {
	data := gin.H{"siteName": SiteName}
	c.HTML(200, "login", data)
}

// Login Login
func Login(c *gin.Context) {
	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))
	if len(username) == 0 || len(password) == 0 {
		c.AbortWithStatusJSON(200, gin.H{"status": -1, "message": "请输入账号密码"})
		return
	} //else {
	user := models.UserGetByName(username)
	if user == nil {
		c.AbortWithStatusJSON(200, gin.H{"status": -1, "message": "用户不存在"})
		return
	} else if user.Password != libs.Md5([]byte(password+user.Salt)) {
		c.AbortWithStatusJSON(200, gin.H{"status": -1, "message": "帐号或密码错误"})
		return
	} else if user.Status == 0 {
		c.AbortWithStatusJSON(200, gin.H{"status": -1, "message": "该帐号已禁用"})
		return
	} else {
		ip := getClientIP(c)
		user.LastIP = ip
		user.LastLogin = time.Now().Unix()

		fields := map[string]interface{}{
			"last_ip":    user.LastIP,
			"last_login": user.LastLogin,
		}
		user.Update(fields)

		// session
		session := sessions.Default(c)
		authkey := libs.Md5([]byte(ip + "|" + user.Password + user.Salt))
		cookie := strconv.Itoa(user.ID) + "|" + user.LoginName + "|" + authkey
		session.Set("auth", cookie)
		session.Save()

		c.JSON(200, gin.H{"status": 0})
	}
	fmt.Println(user)
	// }
}

// Logout Logout
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("auth")

	c.Redirect(301, "/")
}

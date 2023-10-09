package controllers

import (
	"log"
	"strconv"
	"strings"
	"time"

	"gin-layui-admin/libs"
	"gin-layui-admin/models"

	"github.com/gin-gonic/gin"
)

// UserStatusText UserStatusText
var UserStatusText = map[int]string{
	0: "<font color='red'>禁用</font>",
	1: "正常",
}

// UserIndex UserIndex
func UserIndex(c *gin.Context) {
	userName := getUserName(c)
	data := gin.H{
		"siteName":      SiteName,
		"loginUserName": userName,
		"pageTitle":     "用户管理",
	}
	c.HTML(200, "userIndex", data)
}

// UserAddIndex 角色列表
func UserAddIndex(c *gin.Context) {
	// 角色列表
	roles, _ := models.RoleGetList("status", 1)
	list := make([]gin.H, 0, len(roles))
	for _, v := range roles {
		row := make(gin.H)
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list = append(list, row)
	}

	data := gin.H{
		"role":      list,
		"pageTitle": "新增用户",
		"hideTop":   true, // 设置hideTop为true来隐藏公共头部
	}
	c.HTML(200, "userAdd", data)
}

// UserEditIndex UserEditIndex
func UserEditIndex(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("id"))
	user := models.UserGetByID(userID)
	if user != nil {
		row := make(gin.H)
		row["id"] = user.ID
		row["login_name"] = user.LoginName
		row["real_name"] = user.RealName
		row["phone"] = user.Phone
		row["email"] = user.Email
		row["role_id"] = user.RoleID

		roles, _ := models.RoleGetList("status", 1)
		list := make([]gin.H, 0, len(roles))
		for _, v := range roles {
			row := make(gin.H)
			row["id"] = v.Id
			row["role_name"] = v.RoleName
			row["checked"] = 0
			if v.Id == user.RoleID {
				row["checked"] = 1
			}

			list = append(list, row)
		}
		data := gin.H{
			"role":      list,
			"admin":     row,
			"pageTitle": "编辑用户",
			"hideTop":   true, // 设置hideTop为true来隐藏公共头部
		}
		c.HTML(200, "userEdit", data)
	}
}

// UserAjaxSave UserAjaxSave
func UserAjaxSave(c *gin.Context) {
	curUserID := getUserID(c)

	userID, _ := strconv.Atoi(c.PostForm("id"))
	if userID == 0 {
		user := new(models.User)
		c.Bind(user)

		user.UpdateTime = time.Now().Unix()
		user.UpdateID = curUserID
		user.Status = 1

		oldUser := models.UserGetByName(user.LoginName)
		if oldUser != nil {
			c.JSON(200, gin.H{"status": -1, "message": "登录名已经存在"})
			return
		}
		pwd, salt := libs.Password(4, "")
		user.Password = pwd
		user.Salt = salt
		user.CreateTime = time.Now().Unix()
		user.CreateID = curUserID

		if err := user.Insert(); err != nil {
			c.JSON(200, gin.H{"status": -1, "message": err.Error()})
		} else {
			c.JSON(200, gin.H{"status": 0})
		}
	} else {
		user := models.UserGetByID(userID)
		if user == nil {
			c.JSON(200, gin.H{"status": -1, "message": "用户不存在"})
			return
		} //else {
		c.Bind(user)
		log.Println(user)
		// }

		//普通管理员不可修改超级管理员资料
		if curUserID != 1 && user.ID == 1 {
			c.JSON(200, gin.H{"status": -1, "message": "不可修改超级管理员资料"})
			return
		}

		user.UpdateTime = time.Now().Unix()
		fields := map[string]interface{}{}
		if err := user.Update(fields); err != nil {
			c.JSON(200, gin.H{"status": -1, "message": err.Error()})
		} else {
			c.JSON(200, gin.H{"status": 0})
		}
	}
}

// UserAjaxDel UserAjaxDel
func UserAjaxDel(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("id"))
	status := c.PostForm("status")
	user := models.UserGetByID(userID)
	if user != nil {
		if user.ID == 1 {
			c.JSON(200, gin.H{"status": -1, "message": "超级管理员不允许操作"})
			return
		}

		user.UpdateTime = time.Now().Unix()
		user.Status = 0
		if status == "enable" {
			user.Status = 1
		}
		fields := map[string]interface{}{
			"update_time": user.UpdateTime,
			"status":      user.Status,
		}
		if err := user.Update(fields); err != nil {
			c.JSON(200, gin.H{"status": -1, "message": err.Error()})
		} else {
			c.JSON(200, gin.H{"status": 0})
		}
	}
}

// UserTable UserTable
func UserTable(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1000"))
	realName := strings.TrimSpace(c.DefaultQuery("realName", ""))

	var count int
	var users []models.User
	if len(realName) > 0 {
		users, count = models.UserGetPageList(page, limit, "real_name", realName)
	} else {
		users, count = models.UserGetPageList(page, limit)
	}

	list := make([]gin.H, 0, len(users))
	for _, v := range users {
		row := make(gin.H)
		row["id"] = v.ID
		row["login_name"] = v.LoginName
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_id"] = v.RoleID
		row["create_time"] = libs.FormatTime(v.CreateTime)
		row["update_time"] = libs.FormatTime(v.UpdateTime)
		row["status"] = v.Status
		row["status_text"] = UserStatusText[v.Status]
		list = append(list, row)
	}
	c.JSON(200, gin.H{"code": 0, "count": count, "data": list})
}

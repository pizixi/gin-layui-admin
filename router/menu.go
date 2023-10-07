package router

import (
	"fmt"
	"strconv"
	"time"

	"gin-layui-admin/models"

	"github.com/gin-gonic/gin"
)

// MenuIndex MenuIndex
func MenuIndex(c *gin.Context) {
	userName := getUserName(c)
	data := gin.H{
		"siteName":      SiteName,
		"loginUserName": userName,
		"zTree":         true,
		"pageTitle":     "权限因子",
	}
	c.HTML(200, "menuIndex", data)
}

// GetNodes 获取全部节点
func GetNodes(c *gin.Context) {
	menus, count := models.MenuGetList("status", 1)
	list := make([]gin.H, 0, len(menus))
	for _, v := range menus {
		row := make(gin.H)
		row["id"] = v.ID
		row["pId"] = v.Pid
		row["name"] = v.AuthName
		row["open"] = true
		list = append(list, row)
	}

	c.JSON(200, gin.H{"code": 0, "count": count, "data": list})
}

// GetNode 获取一个节点
func GetNode(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	menu := models.MenuGetByID(id)
	if menu != nil {
		row := make(gin.H)
		row["id"] = menu.ID
		row["pid"] = menu.Pid
		row["auth_name"] = menu.AuthName
		row["auth_url"] = menu.AuthURL
		row["sort"] = menu.Sort
		row["is_show"] = menu.IsShow
		row["icon"] = menu.Icon

		c.JSON(200, gin.H{"code": 0, "count": 0, "data": row})
	}
}

// MenuAjaxSave MenuAjaxSave
func MenuAjaxSave(c *gin.Context) {
	userID := getUserID(c)
	menuID, _ := strconv.Atoi(c.PostForm("id"))
	if menuID == 0 {
		menu := new(models.Menu)
		c.Bind(menu)

		menu.Status = 1
		menu.CreateTime = time.Now().Unix()
		menu.CreateID = userID
		menu.UpdateID = userID
		if err := menu.Insert(); err != nil {
			c.JSON(200, gin.H{"status": -1, "message": err.Error()})
		} else {
			c.JSON(200, gin.H{"status": 0})
		}
	} else {
		menu := models.MenuGetByID(menuID)
		if menu == nil {
			c.JSON(200, gin.H{"status": -1, "message": "菜单不存在"})
		} else {
			c.Bind(menu)
		}

		menu.UpdateID = userID
		menu.UpdateTime = time.Now().Unix()
		menu.Update()
		c.JSON(200, gin.H{"status": 0})
	}
}

// MenuAjaxDel MenuAjaxDel
func MenuAjaxDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	menu := models.MenuGetByID(id)
	fmt.Println(menu)
	if menu != nil {
		menu.Status = 0
		menu.UpdateID = getUserID(c)
		menu.UpdateTime = time.Now().Unix()
		menu.Update()
		c.JSON(200, gin.H{"status": 0})
	}
}

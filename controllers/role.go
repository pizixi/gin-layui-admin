package controllers

import (
	"strconv"
	"strings"
	"time"

	"gin-layui-admin/libs"
	"gin-layui-admin/models"

	"github.com/gin-gonic/gin"
)

// RoleStatusText RoleStatusText
var RoleStatusText = map[int]string{
	0: "<font color='red'>禁用</font>",
	1: "正常",
}

// RoleIndex RoleIndex
func RoleIndex(c *gin.Context) {
	userName := getUserName(c)
	data := gin.H{
		"siteName":      SiteName,
		"loginUserName": userName,
		"pageTitle":     "角色管理",
	}
	c.HTML(200, "roleIndex", data)
}

// RoleAddIndex RoleAddIndex
func RoleAddIndex(c *gin.Context) {
	userName := getUserName(c)
	data := gin.H{
		"siteName":      SiteName,
		"loginUserName": userName,
		"zTree":         true,
		"pageTitle":     "新增角色",
	}
	c.HTML(200, "roleAdd", data)
}

// RoleEditIndex RoleEditIndex
func RoleEditIndex(c *gin.Context) {
	userName := getUserName(c)
	data := gin.H{
		"siteName":      SiteName,
		"loginUserName": userName,
		"zTree":         true,
		"pageTitle":     "编辑角色",
	}

	// 角色信息
	id, _ := strconv.Atoi(c.Query("id"))
	role := models.RoleGetByID(id)
	if role != nil {
		row := make(map[string]interface{})
		row["id"] = role.Id
		row["role_name"] = role.RoleName
		row["detail"] = role.Detail
		data["role"] = row

		// 角色的权限列表
		auths := models.MenuGetAuth(role.Id)
		authID := make([]int, 0, len(auths))
		for _, v := range auths {
			authID = append(authID, v.ID)
		}
		data["auth"] = authID
	}

	c.HTML(200, "roleEdit", data)
}

// RoleAjaxSave RoleAjaxSave
func RoleAjaxSave(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.PostForm("id"))
	auths := strings.TrimSpace(c.DefaultPostForm("nodes_data", ""))

	userID := getUserID(c)
	if roleID == 0 {
		// 新增
		role := new(models.Role)
		c.Bind(role)

		role.CreateTime = time.Now().Unix()
		role.UpdateTime = time.Now().Unix()
		role.CreateId = userID
		role.UpdateId = userID
		role.Status = 1

		if err := role.Insert(); err != nil {
			c.JSON(200, gin.H{"status": -1, "message": "新增角色失败"})
		} else {
			roleID = role.Id

			// 更新菜单权限
			authsSlice := strings.Split(auths, ",")
			for _, v := range authsSlice {
				aid, _ := strconv.Atoi(v)
				menu := models.MenuGetByID(aid)
				if menu != nil {
					menu.AuthBit |= (1 << uint(roleID-1))
					menu.Update()
				}
			}
			c.JSON(200, gin.H{"status": 0})
		}
	} else {
		role := models.RoleGetByID(roleID)
		if role == nil {
			c.JSON(200, gin.H{"status": -1, "message": "角色不存在"})
		} else {
			c.Bind(role)
		}

		// 修改
		role.UpdateId = userID
		role.UpdateTime = time.Now().Unix()
		role.UpdateId = userID
		fields := map[string]interface{}{
			"role_name":   role.RoleName,
			"detail":      role.Detail,
			"update_id":   role.UpdateId,
			"update_time": role.UpdateTime,
		}
		if err := role.Update(fields); err != nil {
			c.JSON(200, gin.H{"status": -1, "message": "更新角色失败"})
		} else {
			authsSlice := strings.Split(auths, ",")
			menus := models.MenuGetAuth(roleID)
			for _, menu := range menus {
				hasDel := true
				for _, v := range authsSlice {
					aid, _ := strconv.Atoi(v)
					if aid == menu.ID {
						hasDel = false
						break
					}
				}

				if hasDel {
					menu.AuthBit &= ^(1 << uint(roleID-1))
					menu.Update()
				}
			}

			// 新增的权限
			for _, v := range authsSlice {
				aid, _ := strconv.Atoi(v)
				menu := models.MenuGetByID(aid)
				if menu != nil && menu.Status > 0 {
					menu.AuthBit |= (1 << uint(roleID-1))
					menu.Update()
				}
			}

			c.JSON(200, gin.H{"status": 0})
		}
	}
}

// RoleAjaxDel RoleAjaxDel
func RoleAjaxDel(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.PostForm("id"))
	role := models.RoleGetByID(roleID)
	if role != nil {
		role.Status = 0
		role.UpdateId = getUserID(c)
		role.UpdateTime = time.Now().Unix()
		fields := map[string]interface{}{
			"status":      role.Status,
			"update_id":   role.UpdateId,
			"update_time": role.UpdateTime,
		}
		role.Update(fields)

		// 角色对应的菜单权限要取消
		menus := models.MenuGetAuth(roleID)
		for _, menu := range menus {
			menu.AuthBit &= ^(1 << uint(roleID-1))
			menu.Update()
		}
	}
}

// RoleTable RoleTable
func RoleTable(c *gin.Context) {
	//列表
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1000"))
	roleName := strings.TrimSpace(c.DefaultQuery("roleName", ""))

	var count int
	var roles []models.Role
	if len(roleName) > 0 {
		roles, count = models.RoleGetPageList(page, limit, "role_name", roleName)
	} else {
		roles, count = models.RoleGetPageList(page, limit)
	}

	list := make([]gin.H, 0, len(roles))
	for _, v := range roles {
		row := make(gin.H)
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		row["detail"] = v.Detail
		row["create_time"] = libs.FormatTime(v.CreateTime)
		row["update_time"] = libs.FormatTime(v.UpdateTime)
		row["status_text"] = RoleStatusText[v.Status]
		list = append(list, row)
	}
	c.JSON(200, gin.H{"code": 0, "count": count, "data": list})
}

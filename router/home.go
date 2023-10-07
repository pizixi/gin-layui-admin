package router

import (
	//"log"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"

	"gin-layui-admin/libs"
	"gin-layui-admin/models"

	"github.com/gin-gonic/gin"
)

// HomeIndex HomeIndex
func HomeIndex(c *gin.Context) {
	userID := getUserID(c)
	userName := getUserName(c)

	user := models.UserGetByID(userID)
	menus := models.MenuGetAuth(user.RoleID)
	list := make([]gin.H, 0, len(menus))
	list2 := make([]gin.H, 0, len(menus))
	for _, v := range menus {
		row := make(gin.H)
		if v.Pid == 1 && v.IsShow == 1 {
			// 一级菜单
			row["Id"] = int(v.ID)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthURL
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list = append(list, row)
		}

		if v.Pid != 1 && v.IsShow == 1 {
			//二级菜单
			row["Id"] = int(v.ID)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthURL
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list2 = append(list2, row)
		}
	}

	data := gin.H{
		"title":         SiteName,
		"siteName":      SiteName,
		"loginUserName": userName,
		"SideMenu1":     list,
		"SideMenu2":     list2,
	}
	c.HTML(200, "home", data)
}

// HomeStartIndex HomeStartIndex
func HomeStartIndex(c *gin.Context) {
	userName := getUserName(c)

	// gsCount := models.GameGetCount()
	// agentCount := models.AgentGetCount()
	// cdnCount := models.CDNGetCount()

	// // 最近新开20个的服务器
	// lastSvrs := models.GameGetNearList()
	gsCount := 0
	agentCount := 0
	cdnCount := 0
	type Game struct {
		ID          int    `form:"id"`
		Aid         int    `form:"game_agent" binding:"required"`
		Sid         int    `form:"game_sid" binding:"required"`
		Serial      int    `form:"game_serial"`
		Gid         int    `db:"gid"`
		Mid         int    `db:"mid"`
		Version     int    `form:"game_version" binding:"required"`
		Name        string `db:"name" form:"game_name" binding:"required"`
		Vpsid       int    `form:"game_vps" binding:"required"`
		CreateTime  int64  `db:"create_time"`
		InstallTime int64  `db:"install_time"`
		UpdateTime  int64  `db:"update_time"`
		Port        int    `form:"game_port" binding:"required"`
		DbPort      int    `db:"db_port" form:"db_port"`
		DbShare     int    `db:"db_share"`
		OpenTime    int64  `db:"open_time"`
		MergeTime   int64  `db:"merge_time"`
		IsTLS       int    `db:"is_tls" form:"is_tls"`
		Domain      string
		Procs       string
		Status      int
		NginxID     int `db:"nginx_id" form:"game_nginx"`
		Ws          string
		Single      string
		Mode        int    `form:"game_mode" binding:"required"`
		InstallLog  string `db:"install_log"`
		HotLog      string `db:"hot_log"`
		Hoted       string `db:"hoted"`
		StartLog    string `db:"start_log"`
		StopLog     string `db:"stop_log"`
		Cid         int    `db:"cid"`
		JoinTime    int64  `db:"join_time"`
	}
	// 最近新开20个的服务器
	lastSvrs := []Game{}
	nearGss := make([]gin.H, 0, len(lastSvrs))
	for _, v := range lastSvrs {
		row := make(gin.H)
		row["id"] = v.ID
		row["name"] = v.Name
		row["flag"] = ""
		row["version"] = v.Version
		row["createTime"] = libs.FormatTime(v.CreateTime)
		nearGss = append(nearGss, row)
	}

	// 所有的版本
	vers := GetAllVersions()
	nearVers := make([]gin.H, 0, 20)
	for i, v := range vers {
		if i >= cap(nearVers) {
			break
		}
		nearVers = append(nearVers, gin.H{
			"name": v.Name(),
			"time": v.ModTime().Format("2006-01-02 15:04:05"),
		})
	}

	//系统运行信息
	info := libs.SystemInfo(models.StartTime)

	data := gin.H{
		"siteName":      SiteName,
		"loginUserName": userName,
		"pageTitle":     "系统概况",
		"gameAmount":    gsCount,
		"agentAmount":   agentCount,
		"cdnAmount":     cdnCount,
		"versionAmount": len(vers),
		"nearGsList":    nearGss,
		"nearVers":      nearVers,
		"sysInfo":       info,
	}

	// 系统概况
	c.HTML(http.StatusOK, "start", data)
}

// GetAllVersions 获取所有的版本
func GetAllVersions() []os.FileInfo {
	files, err := ioutil.ReadDir("/data/update_zip")
	if err == nil {
		vers := make([]os.FileInfo, 0, len(files))
		for _, f := range files {
			if f.IsDir() {
				s := regexp.MustCompile(`\w+?_(\w+?)_(\d+)$`).FindStringSubmatch(f.Name())
				if len(s) == 3 {
					vers = append(vers, f)
				}
			}
		}
		sort.Slice(vers, func(i, j int) bool {
			return vers[i].ModTime().Unix() > vers[j].ModTime().Unix()
		})
		return vers
	}
	return nil
}

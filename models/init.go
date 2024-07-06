package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gin-layui-admin/conf"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// StartTime StartTime
var StartTime int64

// Init Init
func Init() {
	StartTime = time.Now().Unix()
	dsn := conf.YamlConf.DBLink
	log.Println(dsn)

	var err error

	// 判断dsn前缀以确定数据库类型
	if strings.HasPrefix(dsn, "mysql:") {
		dsn = strings.TrimPrefix(dsn, "mysql:")
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// 全局禁止表名复数
				SingularTable: true,
			},
			// 日志等级
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else if strings.HasPrefix(dsn, "sqlite:") {
		dsn = strings.TrimPrefix(dsn, "sqlite:")
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// 全局禁止表名复数
				SingularTable: true,
			},
			// 日志等级
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		log.Fatal("不支持的数据库类型")
		return
	}

	if err != nil {
		log.Fatal("数据库连接失败", err.Error())
	}

	// 注册模型
	db.AutoMigrate(&Menu{}, &Role{}, &User{})
	// 插入基础数据
	insertBaseData()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库连接池失败", err.Error())
	}

	sqlDB.SetMaxIdleConns(10)  // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100) // 设置最大打开连接数
}

// initData 插入基础数据
func insertBaseData() {
	// 插入角色数据
	var count int64
	db.Model(&Role{}).Count(&count)
	if count == 0 {
		roles := []Role{
			{Id: 1, RoleName: "超级管理员", Detail: "超级管理员，具有所有权限", CreateId: 1, UpdateId: 1, Status: 1, CreateTime: 0, UpdateTime: 0},
			{Id: 2, RoleName: "普通管理员", Detail: "普通管理员，无菜单管理权限", CreateId: 1, UpdateId: 1, Status: 1, CreateTime: 0, UpdateTime: 0},
			{Id: 3, RoleName: "测试", Detail: "测试角色1", CreateId: 1, UpdateId: 1, Status: 1, CreateTime: 1696664214, UpdateTime: 1696675302},
		}
		db.Create(&roles)
	}

	// 插入用户数据
	db.Model(&User{}).Count(&count)
	if count == 0 {
		users := []User{
			{ID: 1, RoleID: 1, LoginName: "admin", RealName: "admin", Password: "97aae6915f853246abf22773112d2fb4", Phone: "13866668888", Email: "xxx@qq.com", Salt: "58QM", LastLogin: 1720174952, LastIP: "127.0.0.1", Status: 1, CreateID: 0, UpdateID: 0, CreateTime: 0, UpdateTime: 1548146771},
			{ID: 2, RoleID: 2, LoginName: "admin2", RealName: "admin2", Password: "97aae6915f853246abf22773112d2fb4", Phone: "13855556666", Email: "gjs3ob_t@streetsinus.com", Salt: "58QM", LastLogin: 1696599396, LastIP: "127.0.0.1", Status: 1, CreateID: 1, UpdateID: 1, CreateTime: 1696598348, UpdateTime: 1696598348},
			{ID: 4, RoleID: 3, LoginName: "admin4", RealName: "admin4", Password: "55ba69f874cf544b2fa46f4ac9412c54", Phone: "13178687958", Email: "cpktbf@chaco.net", Salt: "pHRV", LastLogin: 1696667404, LastIP: "127.0.0.1", Status: 0, CreateID: 1, UpdateID: 1, CreateTime: 1696664346, UpdateTime: 1696674031},
		}
		db.Create(&users)
	}

	// 插入菜单数据
	db.Model(&Menu{}).Count(&count)
	if count == 0 {
		menus := []Menu{
			{ID: 1, Pid: 0, AuthName: "所有权限", AuthURL: "", Sort: 1, Icon: "", IsShow: 1, AuthBit: 7, UserID: 1, CreateID: 1, UpdateID: 1, Status: 1, CreateTime: 1505620970, UpdateTime: 0},
			{ID: 2, Pid: 1, AuthName: "权限管理", AuthURL: "/", Sort: 999, Icon: "fa-id-card", IsShow: 1, AuthBit: 1, UserID: 1, CreateID: 1, UpdateID: 1, Status: 1, CreateTime: 1505622360, UpdateTime: 0},
			{ID: 3, Pid: 2, AuthName: "用户管理", AuthURL: "/home/user", Sort: 1, Icon: "fa-user-o", IsShow: 1, AuthBit: 1, UserID: 0, CreateID: 0, UpdateID: 1, Status: 1, CreateTime: 1528385411, UpdateTime: 0},
			{ID: 4, Pid: 2, AuthName: "角色管理", AuthURL: "/home/role", Sort: 2, Icon: "fa-user-circle-o", IsShow: 1, AuthBit: 1, UserID: 0, CreateID: 1, UpdateID: 1, Status: 1, CreateTime: 1505621852, UpdateTime: 0},
			{ID: 5, Pid: 2, AuthName: "菜单管理", AuthURL: "/home/menu", Sort: 3, Icon: "fa-list", IsShow: 1, AuthBit: 1, UserID: 1, CreateID: 1, UpdateID: 1, Status: 1, CreateTime: 1505621986, UpdateTime: 0},
			{ID: 6, Pid: 1, AuthName: "运维管理", AuthURL: "/oam", Sort: 1, Icon: "fa-tasks", IsShow: 1, AuthBit: 0, UserID: 1, CreateID: 1, UpdateID: 1, Status: 1, CreateTime: 0, UpdateTime: 0},
			{ID: 8, Pid: 1, AuthName: "个人中心", AuthURL: "/personal", Sort: 999, Icon: "fa-user-circle-o", IsShow: 1, AuthBit: 7, UserID: 1, CreateID: 1, UpdateID: 1, Status: 1, CreateTime: 1547000410, UpdateTime: 0},
			{ID: 12, Pid: 8, AuthName: "资料修改", AuthURL: "/home/personal", Sort: 1, Icon: "fa-edit", IsShow: 1, AuthBit: 7, UserID: 1, CreateID: 1, UpdateID: 1, Status: 1, CreateTime: 1547000565, UpdateTime: 1696675325},
		}
		db.Create(&menus)
	}
}

// MysqlGetList 分页查找范围条目总数，全部查询返回0
func MysqlGetList(table string, des interface{}, filters ...interface{}) int {
	page, pageSize := 0, 0

	tx := db.Table(table)

	for i := 0; i < len(filters); i += 2 {
		field := filters[i].(string)
		if field == "_page" {
			page = filters[i+1].(int)
		} else if field == "_limit" {
			pageSize = filters[i+1].(int)
		} else if field == "_order" {
			tx = tx.Order(filters[i+1].(string))
		} else if field == "_where" {
			tx = tx.Where(filters[i+1].(string))
		} else {
			// 精确匹配
			// tx = tx.Where(fmt.Sprintf("%s = ?", field), filters[i+1])

			// 使用LIKE操作符进行模糊查询
			tx = tx.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%%%v%%", filters[i+1]))
		}
	}

	var count int64
	if page > 0 && pageSize > 0 {
		// 分页查找
		offset := (page - 1) * pageSize

		err := tx.Count(&count).Error
		if err != nil {
			log.Printf("GormGetList [Table=%s] Count failed, err:%s\n", table, err.Error())
		}
		if count <= 0 {
			return 0
		}

		err = tx.Offset(offset).Limit(pageSize).Find(des).Error
		if err != nil {
			log.Printf("GormGetList [Table=%s] Page failed, err:%s\n", table, err.Error())
			return 0
		}
	} else {
		err := tx.Find(des).Error
		if err != nil {
			log.Printf("GormGetList [Table=%s] failed, err:%s\n", table, err.Error())
			return 0
		}
	}

	return int(count)
}

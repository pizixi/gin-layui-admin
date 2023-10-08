package models

import (
	"fmt"
	"log"
	"time"

	"gin-layui-admin/conf"

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
	cnf := conf.YamlConf.SqlCnf
	dbPort := cnf.Port
	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s", cnf.User, cnf.Password, cnf.Host, dbPort, cnf.Name)
	log.Println(dsn)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 全局禁止表名复数
			SingularTable: true,
		},
		// 日志等级
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("mysql 连接失败", err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库连接池失败", err.Error())
	}
	sqlDB.SetMaxIdleConns(10)  // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100) // 设置最大打开连接数
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

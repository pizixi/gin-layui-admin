package models

import (
	"log"
)

// Role Role
type Role struct {
	Id         int    `form:"id"`
	RoleName   string `db:"role_name" form:"role_name" binding:"required"`
	Detail     string `form:"detail" binding:"required"`
	CreateId   int    `db:"create_id"`
	UpdateId   int    `db:"update_id"`
	Status     int
	CreateTime int64 `db:"create_time"`
	UpdateTime int64 `db:"update_time"`
}

// RoleFields RoleFields
var RoleFields = []string{
	"role_name",
	"detail",
	"create_id",
	"update_id",
	"status",
	"create_time",
	"update_time",
}

// Insert Insert
func (r *Role) Insert() error {
	result := db.Create(r)
	if result.Error != nil {
		log.Println("Role Insert Failed, err:", result.Error)
		return result.Error
	}
	r.Id = int(result.RowsAffected)
	return nil
}

// Update Update
func (r *Role) Update(fields map[string]interface{}) error {
	if len(fields) == 0 {
		// 如果没有传递任何参数，则更新所有字段
		fields = map[string]interface{}{
			"role_name":   r.RoleName,
			"detail":      r.Detail,
			"create_id":   r.CreateId,
			"update_id":   r.UpdateId,
			"status":      r.Status,
			"create_time": r.CreateTime,
			"update_time": r.UpdateTime,
		}
	}
	return db.Model(r).Select("*").Updates(fields).Error
}

// RoleGetList RoleGetList
func RoleGetList(filters ...interface{}) ([]Role, int) {
	roles := []Role{}
	count := MysqlGetList("role", &roles, filters...)
	if count == 0 {
		count = len(roles)
	}
	return roles, count
}

// RoleGetPageList 分页查找
func RoleGetPageList(page, limit int, filters ...interface{}) ([]Role, int) {
	fs := []interface{}{"_page", page, "_limit", limit}
	fs = append(fs, filters...)
	return RoleGetList(fs...)
}

// RoleGetByName RoleGetByName
func RoleGetByName(name string) *Role {
	role := &Role{}
	err := db.Where("role_name = ?", name).First(role).Error
	if err != nil {
		log.Println("RoleGetByName failed, err:", err, name)
		return nil
	}
	return role
}

// RoleGetByID RoleGetByID
func RoleGetByID(id int) *Role {
	role := &Role{}
	result := db.First(role, id)
	if result.Error != nil {
		log.Println("RoleGetById failed, err:", result.Error, id)
		return nil
	}
	return role
}

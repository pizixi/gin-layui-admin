package models

import (
	"fmt"
	"log"
)

// Menu Menu
type Menu struct {
	ID         int    `form:"id"`
	Pid        int    `form:"pid"`
	AuthName   string `db:"auth_name" form:"auth_name" binding:"required"`
	AuthURL    string `db:"auth_url" form:"auth_url"`
	AuthBit    uint   `db:"auth_bit"`
	Sort       int    `form:"sort" binding:"required"`
	Icon       string `form:"icon" binding:"required"`
	IsShow     int    `db:"is_show" form:"is_show"`
	UserID     int    `db:"user_id"`
	CreateID   int    `db:"create_id"`
	UpdateID   int    `db:"update_id"`
	Status     int    `db:"status"`
	CreateTime int64  `db:"create_time"`
	UpdateTime int64  `db:"update_time"`
}

// MenuFields MenuFields
var MenuFields = []string{
	"pid",
	"auth_name",
	"auth_url",
	"auth_bit",
	"sort",
	"icon",
	"is_show",
	"user_id",
	"create_id",
	"update_id",
	"status",
	"create_time",
	"update_time",
}

// Insert Insert
func (m *Menu) Insert() error {
	result := db.Create(m)
	if result.Error != nil {
		log.Println("Menu Insert Failed, err:", result.Error)
		return result.Error
	}

	return nil
}

// Update Update
func (m *Menu) Update() error {
	if len(MenuFields) > 0 {
		// 使用Select方法指定需要更新的字段
		result := db.Model(m).Where("id = ?", m.ID).Select("*").Updates(m)
		if result.Error != nil {
			log.Println("Menu Update Failed, err:", result.Error)
			return result.Error
		}
		return nil
	}
	return fmt.Errorf("Menu Update Need fields")
}

// MenuGetList MenuGetList
func MenuGetList(filters ...interface{}) ([]Menu, int) {
	menus := []Menu{}
	count := MysqlGetList("menu", &menus, filters...)
	if count == 0 {
		count = len(menus)
	}
	return menus, count
}

// MenuGetByID MenuGetByID
func MenuGetByID(id int) *Menu {
	menu := &Menu{}
	err := db.First(menu, id).Error
	if err != nil {
		log.Println("MenuGetById failed, err:", err, id)
		return nil
	}
	return menu
}

// MenuGetAuth MenuGetAuth
func MenuGetAuth(roleID int) []Menu {
	if roleID <= 30 && roleID > 0 {
		menus := []Menu{}
		authBit := uint(1) << uint(roleID-1)
		db.Where("status = ? AND auth_bit & ? = ?", 1, authBit, authBit).Order("pid, sort").Find(&menus)
		return menus
	} else {
		return nil
	}
}

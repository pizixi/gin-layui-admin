package models

import (
	"log"
)

// User User
type User struct {
	ID         int    `form:"id"`
	RoleID     int    `db:"role_id" form:"role_id"`
	LoginName  string `db:"login_name" form:"login_name" binding:"required"`
	RealName   string `db:"real_name" form:"real_name" binding:"required"`
	Password   string
	Salt       string
	Phone      string `form:"phone" binding:"required"`
	Email      string `form:"email" binding:"required"`
	LastLogin  int64  `db:"last_login"`
	LastIP     string `db:"last_ip"`
	Status     int    `form:"status"`
	CreateID   int    `db:"create_id"`
	UpdateID   int    `db:"update_id"`
	CreateTime int64  `db:"create_time"`
	UpdateTime int64  `db:"update_time"`
}

// UserFields UserFields
// var UserFields = []string{
// 	"role_id",
// 	"login_name",
// 	"real_name",
// 	"password",
// 	"salt",
// 	"phone",
// 	"email",
// 	"last_login",
// 	"last_ip",
// 	"status",
// 	"create_id",
// 	"update_id",
// 	"create_time",
// 	"update_time",
// }

// Insert Insert
func (u *User) Insert() error {
	result := db.Create(u)
	if result.Error != nil {
		log.Println("User Insert Failed, err:", result.Error)
		return result.Error
	}
	u.ID = int(result.RowsAffected)
	return nil
}

// Update Update
func (u *User) Update(fields map[string]interface{}) error {
	if len(fields) == 0 {
		// 如果没有传递任何参数，则更新所有字段
		fields = map[string]interface{}{
			"id":          u.ID,
			"role_id":     u.RoleID,
			"login_name":  u.LoginName,
			"real_name":   u.RealName,
			"phone":       u.Phone,
			"email":       u.Email,
			"last_login":  u.LastLogin,
			"last_ip":     u.LastIP,
			"status":      u.Status,
			"create_id":   u.CreateID,
			"update_id":   u.UpdateID,
			"create_time": u.CreateTime,
			"update_time": u.UpdateTime,
		}
	}
	return db.Model(u).Select("*").Updates(fields).Error
}

// UserGetList UserGetList
func UserGetList(filters ...interface{}) ([]User, int) {
	users := []User{}
	count := MysqlGetList("user", &users, filters...)
	if count == 0 {
		count = len(users)
	}
	return users, count
}

// UserGetPageList 分页查找
func UserGetPageList(page, limit int, filters ...interface{}) ([]User, int) {
	fs := []interface{}{"_page", page, "_limit", limit}
	fs = append(fs, filters...)
	return UserGetList(fs...)
}

// UserGetByName UserGetByName
func UserGetByName(name string) *User {
	user := &User{}
	err := db.Where("login_name = ?", name).First(user).Error
	if err != nil {
		log.Println("UserGetByName failed, err:", err, name)
		return nil
	}
	return user
}

// UserGetByID UserGetByID
func UserGetByID(id int) *User {
	user := &User{}
	err := db.First(user, id).Error
	if err != nil {
		log.Println("UserGetById failed, err:", err, id)
		return nil
	}
	return user
}

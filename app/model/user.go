package model

import (
	"github.com/jinzhu/gorm"
)

// 定义 User 模型字段
type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int `json:"age"`
}

func init() {
	// 自动迁移
	// 只会 创建表、缺失的列、缺失的索引， 不会 更改现有列的类型或删除未使用的列
	DB.AutoMigrate(&User{})
}

// 创建
func (user *User) Create() bool {
	create := DB.Create(&user)
	if create.Error != nil {
		return false
	}

	return true
}

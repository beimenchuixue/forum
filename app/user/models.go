package user

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	UserId   int64  `gorm:"type:bigint(20);not null;uniqueIndex;comment:用户编号" json:"user_id"`
	Username string `gorm:"type:varchar(64);not null;uniqueIndex;comment:用户名" valid:"Required" json:"username"`
	Password string `gorm:"type:varchar(64);not null;comment:密码" valid:"Required" json:"password"`
	Email    string `gorm:"type:varchar(64);comment:邮箱" json:"email"`
	Gender   int    `gorm:"type:tinyint(4);comment:性别" json:"gender"`
}

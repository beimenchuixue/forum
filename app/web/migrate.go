package web

import (
	"forum/app/user"
	"forum/conn"
)

// auto 自动做数据模型迁移，将数据模型映射到表
func autoMigrate() error {
	return conn.SQL.AutoMigrate(user.User{})
}

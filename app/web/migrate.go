package web

import (
	"bbs/app/user"
	"bbs/conn"
)

// auto 自动做数据模型迁移，将数据模型映射到表
func autoMigrate() error {
	return conn.SQL.AutoMigrate(user.User{})
}

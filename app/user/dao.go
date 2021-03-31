package user

import (
	"bbs/conn"
	"bbs/settings"
	"bbs/utils/sf"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// SignupUser 注册用户
//查询该用户是否存在
//生成uuid
//密码加盐加密
//存入数据库
func SignupUser(u *User) error {
	_, ok := QueryUser(u.Username)
	if !ok {
		u.UserId = sf.GetInt64Id()
		u.Password = GeneratePwd(u.Password)
		conn.SQL.Create(u)
		return nil
	}
	return errors.New("用户已经存在")
}

// 通过用户名查询该用户是否存在
func QueryUser(username string) (*User, bool) {
	var users []User
	conn.SQL.Where("username = ?", username).Find(&users)
	if len(users) > 0 {
		return &users[0], true
	}
	return nil, false
}

// GeneratePwd 生成加密过后的密码
func GeneratePwd(pwd string) string {
	m := sha256.New()
	m.Write([]byte(pwd + settings.Conf.SecretKey))
	return hex.EncodeToString(m.Sum(nil))
}

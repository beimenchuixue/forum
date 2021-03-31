package conn

import (
	"bbs/settings"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SQL *gorm.DB

// init 初始化mysql连接对象
func init() {
	// 连接参数 用户名 密码 主机 端口
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.Conf.MySQLConfig.User, settings.Conf.MySQLConfig.Password, settings.Conf.MySQLConfig.Host,
		settings.Conf.MySQLConfig.Port, settings.Conf.MySQLConfig.DBName,
	)
	fmt.Println(dns)
	// 建立连接
	var err error
	SQL, err = gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}
}

package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// Conf 作为包全局变量，单例模式
var Conf = new(AppConfig)

//AppConfig webApp的总配置
type AppConfig struct {
	Name      string `mapstructure:"name"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
	Bind      string `mapstructure:"bind"`
	Mode      string `mapstructure:"mode"`
	SecretKey string `mapstructure:"secret_key"`
	Author    string `mapstructure:"author"`
	TokenExp  int    `mapstructure:"token_exp"`

	*MySQLConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*SnowflakeConfig `mapstructure:"snowflake"`
	*LogConfig       `mapstructure:"log"`
}

//MySQLConfig mysql连接参数
type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"dbname"`
	MaxConn     string `mapstructure:"max_conn"`
	MaxIdleConn string `mapstructure:"max_idle_conn"`
	Author      string `mapstructure:"author"`
}

//RedisConfig redis连接参数
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
	PoolSize string `mapstructure:"pool_size"`
}

// SnowflakeConfig 雪花算法参数
type SnowflakeConfig struct {
	Time      string `mapstructure:"time"`
	MachineId int64  `mapstructure:"machine_id"`
}

// LogConfig 日志参数
type LogConfig struct {
	Dir       string `mapstructure:"dir"`
	Filename  string `mapstructure:"filename"`
	Level     string `mapstructure:"level"`
	MaxSize   int    `mapstructure:"max_size"`
	MaxAge    int    `mapstructure:"max_age"`
	MaxBackup int    `mapstructure:"max_backup"`
}

func init() {
	//// 1.指定文件名
	//viper.SetConfigFile("config.yaml")
	//// 2. 指定文件查找路径
	//viper.AddConfigPath("./")

	// 配置文件绝对路径
	viper.SetConfigFile("D:\\work\\src\\forum\\config.yaml")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("settings:配置文件未找到")
		} else {
			log.Fatalln("settings:", err)
		}
	}
	// 将配置文件数据映射填充到结构中
	err = viper.Unmarshal(Conf)
	if err != nil {
		log.Fatal(err)
	}

	// 配置文件修改自动重新读取
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(Conf)
		err = viper.Unmarshal(Conf)
		if err != nil {
			log.Fatalln("settings:", err)
		}
	})
}

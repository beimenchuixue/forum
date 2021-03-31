package sf

import (
	"bbs/settings"
	"github.com/bwmarrin/snowflake"
	"time"
)

// 雪花算法 生成分布式id，具有全局唯一性和随时间增长两大特性
// 传入参数: 起始时间和节点id，获取一个id号

var node *snowflake.Node

func init() {
	// 1. 初始化开始时间
	setStartTime()
	var err error
	// 2. 初始化节点
	node, err = snowflake.NewNode(settings.Conf.SnowflakeConfig.MachineId)
	if err != nil {
		panic(err)
	}
}

// setStartTime 设置参考时间起始时间
func setStartTime() {
	startTime, err := time.Parse("2006-01-02", settings.Conf.SnowflakeConfig.Time)
	if err != nil {
		return
	}
	snowflake.Epoch = startTime.Unix()
}

// GetInt64Id 获取int64唯一标识
func GetInt64Id() int64 {
	return node.Generate().Int64()
}

// GetInt64Id 获取string唯一标识
func GetStringId() string {
	return node.Generate().String()
}

// GetBase64Id 获取string唯一标识
func GetBase64Id() string {
	return node.Generate().Base64()
}

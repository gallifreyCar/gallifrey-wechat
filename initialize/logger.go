package initialize

import (
	"github.com/gallifreyCar/gallifrey-wechat/config"
	"github.com/gallifreyCar/gallifrey-wechat/pkg/logger"
)

func InitLogger() {
	// 初始化日志
	logger.InitLogger(
		config.V.GetInt("log.max_size"),
		config.V.GetInt("log.max_backup"),
		config.V.GetInt("log.max_age"),
		config.V.GetBool("log.compress"),
		config.V.GetString("log.level"),
	)
}

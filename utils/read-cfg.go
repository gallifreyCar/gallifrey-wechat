package utils

import (
	"fmt"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"time"
)
import "github.com/spf13/viper"

// ReadMySQLConfig 根据配置文件读取mysql配置
func ReadMySQLConfig(filePath string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("read mysql config failed: %v\n", err)
	}

	config := mysqlCfg.Config{
		User:      viper.GetString("mysql.user"),
		Passwd:    viper.GetString("mysql.password"),
		Net:       "tcp",
		Addr:      viper.GetString("mysql.addr"),
		DBName:    viper.GetString("mysql.dbname"),
		Loc:       time.Local,
		ParseTime: true,
	}

	return config.FormatDSN()
}

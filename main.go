package main

import (
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

func main() {
	// 配置 logrus
	//	1.日志格式改成json
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	//  2.文件名和行号加进去
	log.SetReportCaller(true)
	//  3.设置日志级别
	log.SetLevel(log.DebugLevel)
	logMsg := make(map[string]interface{})
	logMsg["name"] = "wjc"
	logMsg["age"] = 18
	logMsg["sex"] = "男"
	log.WithFields(logMsg).Info("A info msg")

	logMsg2 := make(map[string]interface{})
	logMsg2["name"] = "wmy"
	logMsg2["age"] = 20

	//errorMsg := fmt.Sprintf("B %s msg", "debug")
	log.WithFields(logMsg2).Debug("123412")

	// viper: 配置文件解析
	//环境变量加载我们的程序配置
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("DB_USERNAME", "wjc")
	viper.SetDefault("DB_PASSWORD", "1234")
	viper.SetDefault("DB_PORT", 3306)
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	// 自动加载环境变量
	viper.AutomaticEnv()
	viper.GetString("DB_USERNAME")
	viper.GetInt("DB_PORT")

	logLevel := viper.GetString("LOG_LEVEL")
	username := viper.GetString("DB_USERNAME")
	log.WithFields(log.Fields{
		"日志级别": logLevel,
		"用户名":  username,
	}).Debug("程序信息")

}

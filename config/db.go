package config

import "time"

// 数据库配置
func GetDbConfig() map[string]interface{} {
	// 初始化数据库配置map
	dbConfig := make(map[string]interface{})

	dbConfig["DB_HOST"] = "127.0.0.1"   //数据库地址
	dbConfig["DB_PORT"] = "3306"	//端口
	dbConfig["DB_NAME"] = "jrtt"	//数据库名称
	dbConfig["DB_USER"] = "root"	//账号
	dbConfig["DB_PWD"] = "123456"	//密码
	dbConfig["DB_CHARSET"] = "utf8mb4"
	dbConfig["DB_PREFIX"] = "jr_"
	dbConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	dbConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	dbConfig["DB_MAX_LIFETIME_CONNS"] = time.Hour // 连接池链接最长生命周期

	//redis配置，无视就好
	dbConfig["REDIS_HOST"] = "127.0.0.1"
	dbConfig["REDIS_PORT"] = 6379
	dbConfig["REDIS_PWD"] = ""
	dbConfig["REDIS_SELECT"] = 4
	dbConfig["MAX_IDLE"] = 512
	dbConfig["MAX_ACTIVE"] = 10
	dbConfig["MAX_IDLE_TIMEOUT"] = 200
	dbConfig["TIMEOUT"] = 200


	return dbConfig
}

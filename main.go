package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"gomenter/models"
	"gomenter/routers"
)

func main() {

	// 获取配置
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Println("Failed to parse INI file:", err)
		return
	}
	databaseCfg, err := cfg.GetSection("Database")
	if err != nil {
		fmt.Println("Failed to get Database:", err)
		return
	}
	serverCfg, err := cfg.GetSection("Server")
	if err != nil {
		fmt.Println("Failed to get Serve:", err)
		return
	}

	// 初始化数据库
	models.InitDB(
		databaseCfg.Key("host").String(),
		databaseCfg.Key("name").String(),
	)

	// 启动服务
	r := routers.InitRouter()
	r.Run(fmt.Sprintf(":%s", serverCfg.Key("port").String()))

}

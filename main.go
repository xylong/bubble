package main

import (
	config2 "bubble/config"
	"bubble/pkg/config"
	"bubble/routes"
	"flag"
	"github.com/xylong/bingo"
)

func init() {
	config2.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	var env string

	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	bingo.Init().
		Mount("v1", routes.Controllers...)().
		Lunch(config.GetInt("app.port"))
}

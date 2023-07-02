package config

import (
	viperlib "github.com/spf13/viper"
	"os"
)

var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 1. 初始化 Viper 库
	viper = viperlib.New()
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")
	// 3. 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 5. 读取环境变量（支持 flags）
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量及config信息的加载
func InitConfig(env string) {
	// 1.加载环境变量
	loadEnv(env)
	// 2.注册配置信息
	loadConfig()
}

func loadEnv(envSuffix string) {
	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"

	if envSuffix != "" {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err != nil {
			envPath = filepath
		}
	}

	// 加载env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
}

func loadConfig() {
	for name, f := range ConfigFuncs {
		viper.Set(name, f())
	}
}

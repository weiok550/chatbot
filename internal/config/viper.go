package config

import (
	"chatbot/internal/config/global"
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg *viper.Viper

func init() {
	var configFile string
	var env string
	envPtr := flag.String("env", "prod", "program run env mode")
	flag.Parse()
	fmt.Printf("当前env:%s", *envPtr)
	switch *envPtr {
	case "debug":
	case "test":
	case "release":
		env = *envPtr
	default:
		env = "debug"
	}
	configFile = fmt.Sprintf("./config/config-%s.yaml", env)

	Cfg = viper.New()

	Cfg.SetConfigFile(configFile)
	Cfg.SetConfigType("yaml")

	// 读取配置文件
	err := Cfg.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	}
	// 监听配置文件变化
	Cfg.WatchConfig()
	Cfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Warn Config file changed:", e.Name)

		// 重新读取配置文件
		err := Cfg.ReadInConfig()
		global.Cfg = Cfg
		if err != nil {
			fmt.Println("Error reading config file:", err)
		}
	})
}

func LoadConfig() {
	global.Cfg = Cfg
}

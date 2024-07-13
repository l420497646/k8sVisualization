package initialize

import (
	"fmt"
	"k8sVisualization/global"
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitViper() {
	//获取环境变量env
	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}

	//设置config路径和文件名称
	configFile := fmt.Sprintf("%s-config.yaml", env)

	//viper初始化
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败, %s", err.Error())
	}

	v.Unmarshal(&global.CONF)

}

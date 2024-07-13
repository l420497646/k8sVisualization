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
	configPath := "conf/"
	configName := fmt.Sprintf("%s-config.yaml", env)

	//viper初始化
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败, %s", err.Error())
	}

	if err := v.Unmarshal(&global.CONF); err != nil {
		log.Fatalf("配置文件有误失败, %s", err.Error())
	}

}

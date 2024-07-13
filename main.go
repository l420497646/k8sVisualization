package main

import (
	"k8sVisualization/initialize"
)

// 项目启动入口
func main() {
	//初始化gin路由
	initialize.InitRouters()
}

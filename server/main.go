// main
package main

import (
	"flag"
	"fmt"

	"github.com/coffeehc/logger"
	"github.com/coffeehc/microserviceboot/consultool"
	"github.com/coffeehc/microserviceboot/server"
)

var (
	nodeId = flag.Int("nodeid", 0, "节点编号,最大255")
)

func main() {
	logger.InitLogger()
	serviceRegister, err := consultool.NewConsulServiceRegister(nil)
	if err != nil {
		fmt.Printf("创建服务注册器失败:%s", err)
	}
	if err != nil {
		fmt.Printf("配置文件加载失败:%s", err)
	}
	server.ServiceLauncher(newSequenceService(*nodeId), serviceRegister)
}

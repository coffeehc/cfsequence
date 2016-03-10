// main
package main

import (
	"flag"
	"fmt"

	"github.com/coffeehc/microserviceboot/consultool"
	"github.com/coffeehc/microserviceboot/server"
	"github.com/coffeehc/web"
)

var (
	nodeId    = flag.Int("nodeid", 0, "节点编号,最大255")
	http_Addr = flag.String("http_ip", "", "服务器地址")
)

func main() {
	config := new(server.MicorServiceCofig)
	webConfig := new(web.ServerConfig)
	webConfig.ServerAddr = *http_Addr
	webConfig.DefaultTransport = web.Transport_Json
	config.Service = newSequenceService(0)
	config.DevModule = true
	serviceRegister, err := consultool.NewConsulServiceRegister(nil)
	if err != nil {
		fmt.Printf("创建服务注册器失败:%s", err)
	}
	server.ServiceLauncher(config, serviceRegister)
}

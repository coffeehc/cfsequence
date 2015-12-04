// main
package main

import (
	"flag"
	"fmt"
	"github.com/coffeehc/logger"
	"github.com/coffeehc/web"
	"github.com/coffeehc/utils"
)


var (
	nodeId = flag.Int("nodeid", 0, "节点编号,最大255")
	http_ip = flag.String("http_ip", "0.0.0.0", "服务器地址")
	http_port = flag.Int("http_port", 8888, "服务提供地址")
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	logger.InitLogger()
	if *nodeId < 0 || *nodeId > 255 {
		fmt.Errorf("节点为0-255之间的值,请重新设置")
		return
	}
	var webServer *web.Server
	var sequenceService *SequenceService
	service := utils.NewService(func() error{
		webServer = web.NewServer(&web.ServerConfig{Addr: *http_ip, Port: *http_port})
		sequenceService = newSequenceService(*nodeId)
		sequenceService.regeditRestService(webServer)
		return webServer.Start()
	}  , func()error{
		logger.Debug("webServer is %#v",webServer)
		webServer.Stop()
		return nil
	})
	utils.StartService(service)
}



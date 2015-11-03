// main
package main

import (
	"flag"
	"fmt"
	"github.com/coffeehc/logger"
	"github.com/coffeehc/utils"
	"github.com/coffeehc/web"
)

type service struct {
	webServer       *web.Server
	sequenceService *SequenceService
}

var (
	nodeId    = flag.Int("nodeid", 0, "节点编号,最大255")
	http_ip   = flag.String("http_ip", "0.0.0.0", "服务器地址")
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
	_service := newService()
	args := make([]string, 3)
	args[0] = fmt.Sprintf("-nodeid=%d", *nodeId)
	args[1] = fmt.Sprintf("-http_ip=%s", *http_ip)
	args[2] = fmt.Sprintf("-")
	utils.ProcessServiceWithFlag(utils.CreatService("cfsequence", "cfsequence.service", "coffee's a sequence service", args, _service.start, _service.stop))
}

func newService() *service {
	_service := new(service)
	_service.webServer = web.NewServer(&web.ServerConfig{Addr: *http_ip, Port: *http_port})
	_service.sequenceService = newSequenceService(*nodeId)
	_service.sequenceService.regeditRestService(_service.webServer)
	return _service
}

func (this *service) start() {
	this.webServer.Start()
}

func (this *service) stop() {
	this.webServer.Stop()
}

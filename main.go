// main
package main

import (
	"github.com/coffeehc/logger"
	"github.com/coffeehc/utils"
	"github.com/coffeehc/web"
)

type service struct {
	webServer       *web.Server
	sequenceService *SequenceService
}

func main() {
	logger.SetDefaultLevel("/", logger.LOGGER_LEVEL_INFO)
	_service := newService()
	utils.ProcessServiceWithFlag(utils.CreatService("cfsequence", "cfsequence.service", "coffee's a sequence service", []string{}, _service.start, _service.stop))

}

func newService() *service {
	_service := new(service)
	_service.webServer = web.NewServer(nil)
	_service.sequenceService = newSequenceService()
	_service.sequenceService.regeditRestService(_service.webServer)
	return _service
}

func (this *service) start() {
	this.webServer.Start()
}

func (this *service) stop() {
	this.webServer.Stop()
}

package common

import (
	"io/ioutil"

	"github.com/coffeehc/logger"
	"github.com/coffeehc/microserviceboot/common"
)

type SequenceServiceInfo struct {
	apiDefine string
}

func (this *SequenceServiceInfo) GetApiDefine() string {
	if this.apiDefine == "" {
		data, err := ioutil.ReadFile("apis.raml")
		if err == nil {
			this.apiDefine = string(data)
		} else {
			logger.Error("read file error :%s", err)
			this.apiDefine = "no define"
		}
	}
	return this.apiDefine
}

func (this SequenceServiceInfo) GetServiceName() string {
	return "sequences"
}
func (this SequenceServiceInfo) GetVersion() string {
	return "v1"
}
func (this SequenceServiceInfo) GetDescriptor() string {
	return "a sequence service"
}

func (this SequenceServiceInfo) GetServiceTags() []string {
	return []string{"dev"}
}

func (this SequenceServiceInfo) GetServerPort() int {
	return 8080
}

func (this SequenceServiceInfo) GetScheme() common.RpcScheme {
	return common.RpcScheme_Http
}

func (this SequenceServiceInfo) GetTLSCert() (cartFile, keyFiler string) {
	return "server.crt", "server.key"
}

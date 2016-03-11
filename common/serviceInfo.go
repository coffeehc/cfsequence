package common

import (
	"io/ioutil"

	"github.com/coffeehc/logger"
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
	return "sequence"
}
func (this SequenceServiceInfo) GetVersion() string {
	return "v1"
}
func (this SequenceServiceInfo) GetDescriptor() string {
	return "a sequence service"
}

func (this SequenceServiceInfo) GetServiceTags() []string {
	return nil
}

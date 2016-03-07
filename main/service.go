// web
package main

import (
	"net/http"
	"strconv"

	"io/ioutil"

	"github.com/coffeehc/cfsequence"
	"github.com/coffeehc/logger"
	"github.com/coffeehc/microserviceboot/common"
	"github.com/coffeehc/web"
)

type SequenceService struct {
	_snowflake cfsequence.SequenceService
	apiDefine  string
}

func newSequenceService(nodeId int) *SequenceService {
	_snowflake := cfsequence.NewSequenceService(int64(nodeId))
	return &SequenceService{_snowflake: _snowflake}
}

func (this *SequenceService) Run() error {
	return nil
}
func (this *SequenceService) Stop() error {
	return nil
}

func (this *SequenceService) GetEndPoints() []common.EndPoint {
	return []common.EndPoint{
		common.EndPoint{
			Path:        "/v1/sequences",
			Method:      web.POST,
			HandlerFunc: this.GetNextId,
		},
		common.EndPoint{
			Path:        "/v1/sequences/{id}",
			Method:      web.GET,
			HandlerFunc: this.ParseId,
		},
	}
}

func (this *SequenceService) GetServiceInfo() common.ServiceInfo {
	return this
}

func (this *SequenceService) GetApiDefine() string {
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
func (this *SequenceService) GetServiceName() string {
	return "sequences"
}
func (this *SequenceService) GetVersion() string {
	return "v1"
}
func (this *SequenceService) GetDescriptor() string {
	return "a sequence service"
}

func (this *SequenceService) GetServiceTags() []string {
	return []string{"dev"}
}

type Sequence_Response struct {
	Sequence int64 `json:"sequence"`
}

func (this *SequenceService) GetNextId(request *http.Request, pathFragments map[string]string, reply web.Reply) {
	reply.With(Sequence_Response{this._snowflake.NextId()})
}

func (this *SequenceService) ParseId(request *http.Request, pathFragments map[string]string, reply web.Reply) {
	var response interface{}
	if id, ok := pathFragments["id"]; ok {
		sequence, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			reply.SetStatusCode(422)
			response = common.NewErrorResponse(common.Error{Code: 422, Message: "id is not Number"})
		} else {
			response = this._snowflake.ParseSequence(sequence)
		}
	} else {
		reply.SetStatusCode(500)
		response = common.NewErrorResponse(common.Error{Code: 500, Message: "not parse Path"})
	}
	reply.With(response)
}

// web
package main

import (
	"net/http"
	"strconv"

	"github.com/coffeehc/cfsequence"
	internal "github.com/coffeehc/cfsequence/common"
	"github.com/coffeehc/microserviceboot/common"
	"github.com/coffeehc/web"
)

type SequenceService struct {
	_snowflake cfsequence.SequenceService
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
	return &internal.SequenceServiceInfo{}
}

func (this *SequenceService) GetNextId(request *http.Request, pathFragments map[string]string, reply web.Reply) {
	reply.With(internal.Sequence_Response{this._snowflake.NextId()})
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

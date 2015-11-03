// web
package main

import (
	"net/http"
	"strconv"

	"github.com/coffeehc/cfsequence"
	"github.com/coffeehc/web"
	"github.com/coffeehc/web/rest"
)

type SequenceService struct {
	_snowflake cfsequence.Snowflake
}

func newSequenceService(nodeId int) *SequenceService {
	_snowflake := cfsequence.NewSnowflake(int64(nodeId))
	return &SequenceService{_snowflake: _snowflake}
}

func (this *SequenceService) regeditRestService(server *web.Server) {
	server.Regedit("/v1.0/sequences", web.POST, this.GetNextId)
	server.Regedit("/v1.0/sequences/{id}", web.GET, this.ParseId)
}

type Sequence_Response struct {
	Sequence int64 `json:"sequence"`
}

func (this *SequenceService) GetNextId(request *http.Request, pathFragments map[string]string, reply *web.Reply) {
	reply.With(Sequence_Response{this._snowflake.NextId()}).As(web.Default_JsonTransport)
}

func (this *SequenceService) ParseId(request *http.Request, pathFragments map[string]string, reply *web.Reply) {
	var response interface{}
	if id, ok := pathFragments["id"]; ok {
		sequence, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			reply.SetCode(422)
			response = rest.NewErrorResponse(rest.Error{Code: 422, Message: "id is not Number"})
		} else {
			response = this._snowflake.ParseSequence(sequence)
		}
	} else {
		reply.SetCode(500)
		response = rest.NewErrorResponse(rest.Error{Code: 500, Message: "not parse Path"})
	}
	reply.With(response).As(web.Default_JsonTransport)
}

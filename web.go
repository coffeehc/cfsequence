// web
package main

import (
	"net/http"
	"strconv"

	"github.com/coffeehc/cfsequence/snowflake"
	"github.com/coffeehc/web"
	"github.com/coffeehc/web/rest"
)

type SequenceService struct {
	_snowflake snowflake.Snowflake
}

func newSequenceService() *SequenceService {
	_snowflake := snowflake.NewSnowflake(0)
	return &SequenceService{_snowflake: _snowflake}
}

func (this *SequenceService) regeditRestService(server *web.Server) {
	server.Regedit("/service/sequences", web.POST, this.GetNextId)
	server.Regedit("/service/sequences/{id}", web.GET, this.ParseId)
}

func (this *SequenceService) GetNextId(request *http.Request, pathFragments map[string]string, reply *web.Reply) {
	sequence := this._snowflake.NextId()
	reply.With(rest.RestResponse{Code: 200, Msg: sequence}).As(web.Default_JsonTransport)
}

func (this *SequenceService) ParseId(request *http.Request, pathFragments map[string]string, reply *web.Reply) {
	response := new(rest.RestResponse)
	if id, ok := pathFragments["id"]; ok {
		sequence, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.Code = 422
			response.Msg = &rest.Error{Name: "422", Message: "id is not Number"}
		} else {
			meta := this._snowflake.ParseSequence(sequence)
			response.Code = 200
			response.Msg = meta
		}
	} else {
		response.Code = 500
		response.Msg = &rest.Error{Name: "500", Message: "not parse Path"}
	}
	reply.With(response).As(web.Default_JsonTransport)
}

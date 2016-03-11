package client

import (
	"fmt"

	"flag"

	"github.com/coffeehc/cfsequence"
	"github.com/coffeehc/cfsequence/common"
	"github.com/coffeehc/logger"
	"github.com/coffeehc/microserviceboot/client"
	"github.com/coffeehc/resty"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type SequenceApi struct {
	client *client.ServiceClient
	config *client.ServiceClientConfig
}

var (
	domain     = flag.String("domain", "", "根域")
	datacenter = flag.String("datacenter", "dc", "数据中心")
	dns        = flag.String("dns", "127.0.0.1:8600", "dns 地址")
)

func NewSequenceApi() (*SequenceApi, error) {
	if *domain == "" {
		return nil, errors.New("没有定义根域")
	}
	config := &client.ServiceClientConfig{
		Domain:     *domain,
		DataCenter: *datacenter,
		Info:       &common.SequenceServiceInfo{},
		DNSAddress: *dns,
	}
	serviceClient, err := client.NewServiceClient(config, nil)
	if err != nil {
		return nil, err
	}
	api := &SequenceApi{
		client: serviceClient,
		config: config,
	}
	serviceClient.ApiRegiter(command_newSequemce, api_NewSequemce)
	serviceClient.ApiRegiter(command_SequenceInfo, api_SequenceInfo)
	return api, nil
}

func api_NewSequemce(request *resty.Request, query map[string]string, body interface{}) (*resty.Response, error) {
	return request.Post("/v1/sequences")
}

func api_SequenceInfo(request *resty.Request, query map[string]string, body interface{}) (*resty.Response, error) {
	return request.Get(fmt.Sprintf("/v1/sequences/%s", query["id"]))
}

func (this *SequenceApi) NewSequence() int64 {
	sequenceResp := new(common.Sequence_Response)
	err := this.client.SyncCallApiExt(command_newSequemce, nil, nil, sequenceResp)
	if err != nil {
		fmt.Println(logger.Error("获取新的 Sequence出错:%s", err))

		return 0
	}
	return sequenceResp.Sequence

}

func (this *SequenceApi) ParseSequnece() *cfsequence.Sequence {
	sequence := new(cfsequence.Sequence)
	err := this.client.SyncCallApiExt(command_SequenceInfo, nil, nil, sequence)
	if err != nil {
		return nil
	}
	return sequence
}

package client_test

import (
	"testing"

	"flag"

	"github.com/coffeehc/cfsequence/client"
)

func BenchmarkClientAPI(b *testing.B) {
	flag.Set("domain", "test")
	sequenceApi, err := client.NewSequenceApi()
	if err != nil {
		b.Logf("创建 Api失败:%s", err)
		b.FailNow()
	}
	for i := 0; i < b.N; i++ {
		id := sequenceApi.NewSequence()
		if id == 0 {
			b.FailNow()
		}
	}
}

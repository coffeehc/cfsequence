// snowflake_test
package cfsequence_test

import (
	"sync"
	"testing"
	"time"

	"github.com/coffeehc/cfsequence"
	"gopkg.in/check.v1"
)

type SequenceSuite struct {
	sequenceService cfsequence.SequenceService
}

func Test(t *testing.T) {
	check.TestingT(t)
}

func (this *SequenceSuite) SetUpSuite(c *check.C) {
	this.sequenceService = cfsequence.NewSequenceService(0)
}

var _ = check.Suite(&SequenceSuite{})

func (this *SequenceSuite) TestNewSequenceService(c *check.C) {
	c.Assert(cfsequence.NewSequenceService(10000000000), check.Equals, nil)
}
func (this *SequenceSuite) TestEpoch(c *check.C) {
	c.Assert(this.sequenceService.GetEpoch(), check.Equals, int64(1444281954363))
}

func (this *SequenceSuite) TestGetNodeId(c *check.C) {
	c.Assert(this.sequenceService.GetNodeId(), check.Equals, int64(0))
}

func (this *SequenceSuite) TestNextId(c *check.C) {
	waitGroup := new(sync.WaitGroup)
	for i := 0; i < 100000; i++ {
		waitGroup.Add(1)
		go func(group *sync.WaitGroup) {
			defer group.Done()
			c.Assert(this.sequenceService.NextId(), check.Not(check.Equals), this.sequenceService.NextId())
		}(waitGroup)
	}
	waitGroup.Wait()
}

func (this *SequenceSuite) TestParseSequence(c *check.C) {
	sequence := this.sequenceService.ParseSequence(0)
	c.Assert(sequence.NodeId, check.Equals, int64(0))
	c.Assert(sequence.Index, check.Equals, int64(0))
	c.Assert(sequence.CreateTime, check.Equals, time.Unix(0, int64(1444281954363*time.Millisecond)))
	const (
		epoch = int64(1444281954363)
		//节点Id位数
		nodeIdBits = 8
		//最大节点Id
		//	maxNodeId = -1 ^ (-1 << nodeIdBits)
		//	nodeMask  = -1 ^ (-1 << nodeIdBits)
		//序列号位数
		indexBits = 10
		//序列号偏移量
		indexShift = nodeIdBits
		indexMask  = -1 ^ (-1 << indexBits)
		//时间戳偏移量
		timestampLeftShift = indexBits + nodeIdBits

		millisecond = int64(time.Millisecond)
	)
	timestamp := time.Now().UnixNano() / millisecond
	index := int64(indexMask)
	nodeId := int64(0)
	seq := (timestamp-epoch)<<timestampLeftShift | index<<indexShift | nodeId
	sequence = this.sequenceService.ParseSequence(seq)
	c.Assert(sequence.NodeId, check.Equals, nodeId)
	c.Assert(sequence.Index, check.Equals, index)
	c.Assert(sequence.CreateTime, check.Equals, time.Unix(0, timestamp*millisecond))
}

func (this *SequenceSuite) BenchmarkSequence(c *check.C) {
	waitGroup := new(sync.WaitGroup)
	n := c.N
	result := make([]int64, 10000*n)
	for i := 0; i < n; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer func() {
				waitGroup.Done()
			}()
			index := i * 10000
			for j := index; j < index+10000; j++ {
				result[j] = this.sequenceService.NextId()
			}
		}(i)
	}
	waitGroup.Wait()
	cache := make(map[int64]int, 10000*n)
	for i, data := range result {
		if d, ok := cache[data]; ok {
			c.Errorf("重复的Id:%d,old index is %d,new index is %d", data, d, i)
		} else {
			cache[data] = i
		}
	}
}

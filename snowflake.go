// snowflake
package cfsequence

import (
	"time"
	"sync"
)


type _snowflake struct {
	//节点ID
	nodeId int64
	mutex *sync.Mutex
	index int64
	lastTimestamp int64
}

func NewSequenceService(nodeId int64) SequenceService {
	if nodeId > maxNodeId {
		return nil
	}
	s := &_snowflake{nodeId, new(sync.Mutex),0,0}
	return s
}

func (this *_snowflake) GetEpoch() int64 {
	return epoch
}

func (this *_snowflake) GetNodeId() int64 {
	return this.nodeId
}


func (this *_snowflake)getTimeStampAndIndex()(int64,int64){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	index := this.index
	timestamp := getMillisecond()
	if this.lastTimestamp == timestamp/millisecond{
		index = (index + 1) & indexMask
		if index == 0 {
			//当前毫秒内计数满了，则等待下一秒
			timestamp = tilNextMillis(this.lastTimestamp*millisecond)
		}
	}else{
		index = 0
	}
	this.index = index
	this.lastTimestamp = timestamp/millisecond
	return this.lastTimestamp,index
}

func (this *_snowflake) NextId() int64 {
	t,i:=this.getTimeStampAndIndex()
	return (t-epoch)<<timestampLeftShift | i<<indexShift | this.nodeId
}

func (this *_snowflake) ParseSequence(sequence int64) *Sequence {
	index := sequence >> indexShift & indexMask
	unixTime := (epoch + (sequence >> timestampLeftShift)) * millisecond
	createTime := time.Unix(0, unixTime)
	nodeId := sequence & nodeMask
	return &Sequence{sequence, createTime, nodeId, index}
}

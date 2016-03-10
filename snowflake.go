// snowflake
package cfsequence

import "time"

type _snowflake struct {
	//节点ID
	nodeId int64
	//
	idChannle chan chan int64
}

func NewSequenceService(nodeId int64) SequenceService {
	if nodeId > maxNodeId {
		return nil
	}
	s := &_snowflake{nodeId, make(chan chan int64, 1024)}
	go s.idsbuilder()
	return s
}

func (this *_snowflake) GetEpoch() int64 {
	return epoch
}

func (this *_snowflake) GetNodeId() int64 {
	return this.nodeId
}

func (this *_snowflake) idsbuilder() {
	var index, lastTimestamp int64 = 0, 0
	for pipe := range this.idChannle {
		timestamp := getMillisecond()
		if lastTimestamp == timestamp {
			index = (index + 1) & indexMask
			if index == 0 {
				//当前毫秒内计数满了，则等待下一秒
				timestamp = tilNextMillis(lastTimestamp)
			}
		} else {
			index = 0
		}
		lastTimestamp = timestamp
		pipe <- (timestamp-epoch)<<timestampLeftShift | index<<indexShift | this.nodeId

	}
}

func (this *_snowflake) NextId() int64 {
	//t1 := time.Now()
	pipe := make(chan int64)
	this.idChannle <- (pipe)
	id := <-pipe
	//fmt.Printf("比较 is %s\n", time.Since(t1))
	return id
}

func (this *_snowflake) ParseSequence(sequence int64) Sequence {
	index := sequence >> indexShift & indexMask
	unixTime := (epoch + (sequence >> timestampLeftShift)) * millisecond
	createTime := time.Unix(0, unixTime)
	nodeId := sequence & nodeMask
	return Sequence{sequence, createTime, nodeId, index}
}

// snowflake_test
package snowflake

import (
	"sync"
	"testing"
)

func BenchmarkSnowflake(b *testing.B) {
	snowflake := NewSnowflake(0)
	if snowflake != nil {
		waitGroup := new(sync.WaitGroup)
		n := b.N
		result := make([]int64, 10000*n)
		for i := 0; i < n; i++ {
			waitGroup.Add(1)
			go func(i int) {
				defer func() {
					waitGroup.Done()
				}()
				index := i * 10000
				for j := index; j < index+10000; j++ {
					result[j] = snowflake.NextId()
				}
			}(i)
		}
		waitGroup.Wait()
		cache := make(map[int64]int, 10000*n)
		for i, data := range result {
			if d, ok := cache[data]; ok {
				b.Errorf("重复的Id:%d,old index is %d,new index is %d", data, d, i)
			} else {
				cache[data] = i
			}
		}
	} else {
		b.Errorf("不能获取snowflake实现")
	}

}

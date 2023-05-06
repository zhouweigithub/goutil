package threadutil

import (
	"sync"
	"sync/atomic"
)

// 启用多线程完成任务
func Threading[T any](sources []T, threadCount int, action func(item *T)) {
	var sourceCount = len(sources)
	if len(sources) == 0 || threadCount == 0 || action == nil {
		return
	}
	if threadCount > sourceCount {
		threadCount = sourceCount
	}
	var group = &sync.WaitGroup{}
	var runningIndex atomic.Int32

	for i := 0; i < threadCount; i++ {
		group.Add(1)
		go func() {
			for {
				var index = int(runningIndex.Load())
				if index >= sourceCount {
					break
				}
				runningIndex.Add(1)
				action(&sources[index])
			}
			group.Done()
		}()
	}
	group.Wait()
}

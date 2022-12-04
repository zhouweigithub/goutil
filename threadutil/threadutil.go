package threadutil

import (
	"sync"
)

// 启用多线程完成任务
func Threading[T any](sources []T, threadCount int, action func(item *T)) {
	if len(sources) == 0 || threadCount == 0 || action == nil {
		return
	}
	if threadCount > len(sources) {
		threadCount = len(sources)
	}
	var mutex = &sync.Mutex{}
	var group = &sync.WaitGroup{}
	var runningIndex = 0

	for i := 0; i < threadCount; i++ {
		group.Add(1)
		go func() {
			for {
				mutex.Lock()
				if runningIndex >= len(sources) {
					mutex.Unlock()
					break
				}
				//fmt.Println(runningIndex)
				var curIndex = runningIndex
				runningIndex++
				mutex.Unlock()

				action(&sources[curIndex])
			}
			group.Done()
		}()
	}

	group.Wait()
}

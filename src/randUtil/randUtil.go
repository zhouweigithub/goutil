package randutil

import (
	"math/rand"
	"time"
)

// GetRandInts 生成随机数,区间为[min,max]
func GetRandInts(min, max, count int) []int {
	var result []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		result = append(result, min+rand.Intn(max-min+1))
	}
	return result
}

// GetRandInt 生成随机数,区间为[min,max]
func GetRandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

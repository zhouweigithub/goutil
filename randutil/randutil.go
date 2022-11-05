package randutil

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/zhouweigithub/goutil/sliceutil"
)

var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var upChars = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var lowChars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var upAndLowChars []string
var allLetters []string

func init() {
	upAndLowChars = append(sliceutil.CopySlice(upChars), lowChars...)
	allLetters = append(sliceutil.CopySlice(upAndLowChars), numbers...)
}

// GetRandInts 生成随机数,区间为[min,max]
func GetRandInts(min, max, count int) []int {
	var result []int
	time.Sleep(time.Millisecond * 1)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		result = append(result, min+rand.Intn(max-min+1))
	}
	return result
}

// GetRandInt 生成随机数,区间为[min,max]
func GetRandInt(min, max int) int {
	time.Sleep(time.Millisecond * 1)
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

// 获取随机数字
func GetRandNumbers(count int) string {
	var sb strings.Builder
	var ints = GetRandInts(0, 9, count)
	for i := 0; i < len(ints); i++ {
		sb.WriteString(strconv.Itoa(ints[i]))
	}
	return sb.String()
}

// 获取随机大写字母
func GetRandUpperChars(count int) string {
	var sb strings.Builder
	var ints = GetRandInts(0, len(upChars)-1, count)
	for i := 0; i < len(ints); i++ {
		sb.WriteString(upChars[ints[i]])
	}
	return sb.String()
}

// 获取随机小写字母
func GetRandLowerChars(count int) string {
	var sb strings.Builder
	var ints = GetRandInts(0, len(lowChars)-1, count)
	for i := 0; i < len(ints); i++ {
		sb.WriteString(lowChars[ints[i]])
	}
	return sb.String()
}

// 获取随机大小写字母
func GetRandUpperAndLowerChars(count int) string {
	var sb strings.Builder
	var ints = GetRandInts(0, len(upAndLowChars)-1, count)
	for i := 0; i < len(ints); i++ {
		sb.WriteString(upAndLowChars[ints[i]])
	}
	return sb.String()
}

// 获取随机大小写字母和数字
func GetRandChars(count int) string {
	var sb strings.Builder
	var ints = GetRandInts(0, len(allLetters)-1, count)
	for i := 0; i < len(ints); i++ {
		sb.WriteString(allLetters[ints[i]])
	}
	return sb.String()
}

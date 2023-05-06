package sliceutil

import "sort"

// 查找第一个满足条件的元素，无匹配项则返回nil
func First[T any](datas []T, filter func(item *T) bool) *T {
	if len(datas) == 0 {
		return nil
	}
	for i := range datas {
		if filter(&datas[i]) {
			return &datas[i]
		}
	}
	return nil
}

// 查找最后一个满足条件的元素，无匹配项则返回nil
func Last[T any](datas []T, filter func(item *T) bool) *T {
	if len(datas) == 0 {
		return nil
	}
	for i := len(datas) - 1; i >= 0; i-- {
		if filter(&datas[i]) {
			return &datas[i]
		}
	}
	return nil
}

// 查找所有满足条件的元素副本，无匹配项则返回空切片
func Where[T any](datas []T, filter func(item *T) bool) []T {
	var result = make([]T, 0, len(datas))
	if len(datas) == 0 {
		return nil
	}
	for i := range datas {
		if filter(&datas[i]) {
			result = append(result, datas[i])
		}
	}
	return result
}

// 查找所有满足条件的元素的引用，无匹配项则返回空切片
func WhereReference[T any](datas []T, filter func(item *T) bool) []*T {
	var result = make([]*T, 0, len(datas))
	if len(datas) == 0 {
		return nil
	}
	for i := range datas {
		if filter(&datas[i]) {
			result = append(result, &datas[i])
		}
	}
	return result
}

// 检查元素是否存在
func Contains[T comparable](datas []T, item T) bool {
	if len(datas) == 0 {
		return false
	}
	for i := range datas {
		if datas[i] == item {
			return true
		}
	}
	return false
}

// 检查是否存在满足条件的元素
func ContainsFunc[T any](datas []T, filter func(item *T) bool) bool {
	if len(datas) == 0 {
		return false
	}
	for i := range datas {
		if filter(&datas[i]) {
			return true
		}
	}
	return false
}

// 过滤部分字段
func Select[T any, K any](datas []T, filter func(item *T) K) []K {
	var result = []K{}
	if len(datas) == 0 {
		return nil
	}
	for i := range datas {
		result = append(result, filter(&datas[i]))
	}
	return result
}

// 查找第一个元素的索引
func IndexOf[T comparable](datas []T, item T) int {
	if len(datas) == 0 {
		return -1
	}
	for i := range datas {
		if datas[i] == item {
			return i
		}
	}
	return -1
}

// 查找第一个满足条件元素的索引
func IndexOfFunc[T any](datas []T, filter func(item *T) bool) int {
	if len(datas) == 0 {
		return -1
	}
	for i := range datas {
		if filter(&datas[i]) {
			return i
		}
	}
	return -1
}

// 查找最后一个元素的索引
func LastIndexOf[T comparable](datas []T, item T) int {
	if len(datas) == 0 {
		return -1
	}
	for i := len(datas) - 1; i >= 0; i-- {
		if datas[i] == item {
			return i
		}
	}
	return -1
}

// 查找最后一个满足条件元素的索引
func LastIndexOfFunc[T any](datas []T, filter func(item *T) bool) int {
	if len(datas) == 0 {
		return -1
	}
	for i := len(datas) - 1; i >= 0; i-- {
		if filter(&datas[i]) {
			return i
		}
	}
	return -1
}

// 遍历元素，并执行指定操作
func ForEach[T any](datas []T, action func(item *T)) {
	if len(datas) == 0 {
		return
	}
	for i := range datas {
		action(&datas[i])
	}
}

// 查找满足条件元素的数量
func Count[T any](datas []T, filter func(item *T) bool) int {
	if len(datas) == 0 {
		return 0
	}
	var count = 0
	for i := range datas {
		if filter(&datas[i]) {
			count++
		}
	}
	return count
}

// 删除元素
func Remove[T comparable](datas []T, item T) []T {
	if len(datas) == 0 {
		return datas
	}
	var result = []T{}
	for i := range datas {
		if datas[i] != item {
			result = append(result, datas[i])
		}
	}
	return result
}

// 删除满足条件的所有元素
func RemoveFunc[T any](datas []T, filter func(item *T) bool) []T {
	if len(datas) == 0 {
		return datas
	}
	var result = []T{}
	for i := range datas {
		if !filter(&datas[i]) {
			result = append(result, datas[i])
		}
	}
	return result
}

// 获取去重后的元素集
func Distinct[T comparable](datas []T) []T {
	var result = []T{}
	if len(datas) == 0 {
		return result
	}
	var distinctFieldValues []T
	for i := range datas {
		var item = datas[i]
		if !Contains(distinctFieldValues, item) {
			distinctFieldValues = append(distinctFieldValues, item)
			result = append(result, datas[i])
		}
	}
	return result
}

// 自定义排序
func OrderBy[T any](datas []T, filter func(i, j int) bool) {
	if len(datas) == 0 {
		return
	}
	sort.Slice(datas, filter)
}

// 复制切片
func Copy[T any](datas []T) []T {
	var result = make([]T, len(datas))
	copy(result, datas)
	return result
}

// 合并切片
func Union[T any](datas1 []T, datas2 []T) []T {
	return append(datas1, datas2...)
}

// 排除切片
func Exclude[T comparable](datas []T, excludeDatas []T) []T {
	if len(datas) == 0 {
		return []T{}
	}
	if len(excludeDatas) == 0 {
		return datas
	}
	var result = []T{}
	for i := range datas {
		if !Contains(excludeDatas, datas[i]) {
			result = append(result, datas[i])
		}
	}
	return result
}

// 对切片倒序
func Reverse[T any](datas []T) {
	for i, j := 0, len(datas)-1; i < j; i, j = i+1, j-1 {
		datas[i], datas[j] = datas[j], datas[i]
	}
}

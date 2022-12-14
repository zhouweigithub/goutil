package queryutil

import (
	"sort"
)

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

// 查找所有满足条件的元素，无匹配项则返回空切片
func Where[T any](datas []T, filter func(item *T) bool) []*T {
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

// 检查是否存在满足条件的元素
func Contains[T any](datas []T, filter func(item *T) bool) bool {
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

// 查找第一个满足条件元素的索引
func IndexOf[T any](datas []T, filter func(item *T) bool) int {
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

// 查找最后一个满足条件元素的索引
func LastIndexOf[T any](datas []T, filter func(item *T) bool) int {
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

// 删除满足条件的所有元素
func Remove[T any](datas []T, filter func(item *T) bool) []T {
	if len(datas) == 0 {
		return datas
	}
	var validIndex = 0
	for i := range datas {
		if !filter(&datas[i]) {
			datas[validIndex] = datas[i]
			validIndex++
		}
	}
	return datas[:validIndex]
}

// 获取去重后的元素集
func Distinct[T comparable](datas []T) []T {
	var result = []T{}
	if len(datas) == 0 {
		return result
	}
	var distinctFieldValues []T
	for i := range datas {
		var v = datas[i]
		if !Contains(distinctFieldValues, func(subItem *T) bool { return *subItem == v }) {
			distinctFieldValues = append(distinctFieldValues, v)
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

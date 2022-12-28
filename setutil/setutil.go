package setutil

import "sort"

// 不包含重复元素的集合（非线程安全）
type myset[T string | int | int8 | int16 | int32 | int64 | float32 | float64] map[T]struct{}

// 创建不包含重复元素的集合新实例（非线程安全）
func NewSet[T string | int | int8 | int16 | int32 | int64 | float32 | float64]() myset[T] {
	return make(myset[T])
}

// 添加元素
func (s myset[T]) Add(keys ...T) {
	for _, v := range keys {
		s[v] = struct{}{}
	}
}

// 删除元素
func (s myset[T]) Delete(keys ...T) {
	for _, v := range keys {
		delete(s, v)
	}
}

// 是否存在元素
func (s myset[T]) Exists(key T) bool {
	if _, isOk := s[key]; isOk {
		return true
	}
	return false
}

// 获取所有元素
func (s myset[T]) GetAll() []T {
	var datas = make([]T, 0, len(s))
	for key := range s {
		datas = append(datas, key)
	}
	return datas
}

// 获取所有元素并按照从小到大排序
func (s myset[T]) GetAllSorted() []T {
	var datas = s.GetAll()
	sort.Slice(datas, func(i, j int) bool { return datas[i] < datas[j] })
	return datas
}

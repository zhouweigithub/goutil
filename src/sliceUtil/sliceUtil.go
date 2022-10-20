package sliceutil

// FindSlice 查找切片内容
func FindSlice(slice []string, val string) (int, bool) {
	for i := range slice {
		if slice[i] == val {
			return i, true
		}
	}
	return -1, false
}

// ExcludeSlice 排除部分元素
func ExcludeSlice(slice []string, excludes []string) []string {
	var newSlice = make([]string, 0)
	for i := range slice {
		_, isFind := FindSlice(excludes, slice[i])
		if !isFind {
			newSlice = append(newSlice, slice[i])
		}
	}

	return newSlice
}

// CopySlice 复制切片内容
func CopySlice(slice []string) []string {
	var newSlice = make([]string, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// DeleteSlice 删除切片指定元素，会修改原切片a
func DeleteSlice(source []string, elem string) []string {
	j := 0
	for i := range source {
		if source[i] != elem {
			source[j] = source[i]
			j++
		}
	}
	return source[:j]
}

// DeleteSlice2 删除切片指定元素，不会修改原切片a
func DeleteSlice2(source []string, elem string) []string {
	tmp := make([]string, 0, len(source))
	for i := range source {
		if source[i] != elem {
			tmp = append(tmp, source[i])
		}
	}
	return tmp
}

// GetDistinct 获取去重后的新切片
func GetDistinct(slice []string) []string {
	var result []string
	for i := range slice {
		if _, isExist := FindSlice(result, slice[i]); !isExist {
			result = append(result, slice[i])
		}
	}
	return result
}

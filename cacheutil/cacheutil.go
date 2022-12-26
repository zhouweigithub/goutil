package cacheutil

import (
	"errors"
	"time"
)

// 缓存管理中心
type cacheManager struct {
	datas map[string]cacheItem
}

// 缓存的具体内容
type cacheItem struct {
	Value      interface{} //缓存的值
	ExpireTime time.Time   //缓存的过期时间
}

// 是否过期
func (c cacheItem) IsExpired() bool {
	return c.ExpireTime.Before(time.Now())
}

// 创建新实例
func New() *cacheManager {
	var c cacheManager
	c.datas = make(map[string]cacheItem)
	return &c
}

// 获取缓存内容
func (c *cacheManager) Get(key string) interface{} {
	if key == "" {
		return nil
	}
	if value, isExists := c.datas[key]; !isExists {
		return nil
	} else if value.IsExpired() {
		delete(c.datas, key)
		return nil
	} else {
		return value.Value
	}
}

// 设置缓存内容
func (c *cacheManager) Set(key string, value interface{}, timeDuration time.Duration) error {
	if key == "" {
		return errors.New("key不能为空")
	}
	c.datas[key] = cacheItem{
		Value:      value,
		ExpireTime: time.Now().Add(timeDuration),
	}
	return nil
}

// 删除缓存内容
func (c *cacheManager) Delete(key string) bool {
	if key == "" {
		return false
	}

	delete(c.datas, key)
	return true
}

// 检测缓存是否存在
func (c *cacheManager) Exists(key string) bool {
	if key == "" {
		return false
	}
	if value, isExists := c.datas[key]; !isExists {
		return false
	} else {
		if !value.IsExpired() {
			return true
		} else {
			// 删除已过期的项
			delete(c.datas, key)
			return false
		}
	}
}

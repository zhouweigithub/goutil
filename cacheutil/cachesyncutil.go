package cacheutil

import (
	"errors"
	"sync"
	"time"
)

// 缓存管理中心
type cacheSyncManager struct {
	// 缓存数据集
	datas sync.Map
}

// 创建新实例（线程安全）
func NewSync() cacheBase {
	var c cacheSyncManager
	return &c
}

// 获取缓存内容
func (c *cacheSyncManager) Get(key string) interface{} {
	if key == "" {
		return nil
	}

	if value, isExists := c.datas.Load(key); !isExists {
		return nil
	} else if value.(cacheItem).IsExpired() {
		c.datas.Delete(key)
		return nil
	} else {
		return value.(cacheItem).Value
	}
}

// 设置缓存内容
func (c *cacheSyncManager) Set(key string, value interface{}, timeDuration time.Duration) error {
	if key == "" {
		return errors.New("key不能为空")
	}
	c.datas.Store(key, cacheItem{
		Value:      value,
		ExpireTime: time.Now().Add(timeDuration),
	})
	return nil
}

// 删除缓存内容
func (c *cacheSyncManager) Delete(key string) bool {
	if key == "" {
		return false
	}

	c.datas.Delete(key)
	return true
}

// 检测缓存是否存在
func (c *cacheSyncManager) Exists(key string) bool {
	if key == "" {
		return false
	}
	if value, isExists := c.datas.Load(key); !isExists {
		return false
	} else {
		if !value.(cacheItem).IsExpired() {
			return true
		} else {
			// 删除已过期的项
			c.datas.Delete(key)
			return false
		}
	}
}

// 获取缓存数量
func (c *cacheSyncManager) Len() int {
	var count = 0
	c.datas.Range(func(key, value any) bool {
		count++
		return true //需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。
	})
	return count
}

package cacheutil

import "time"

// 缓存接口
type cacheBase interface {
	Get(key string) interface{}
	Set(key string, value interface{}, timeDuration time.Duration) error
	Delete(key string) bool
	Exists(key string) bool
	Len() int
}

// 缓存的具体内容
type cacheItem struct {
	//缓存的值
	Value interface{}
	//缓存的过期时间
	ExpireTime time.Time
}

// 是否过期
func (c cacheItem) IsExpired() bool {
	return c.ExpireTime.Before(time.Now())
}

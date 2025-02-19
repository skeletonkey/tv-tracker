package tvdb

import (
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

var globalCache *cache.Cache

const (
	cacheDefaultExp = 30 // minutes
	cacheCleanupInt = 50 // minutes
)

func getCache() *cache.Cache {
	if globalCache == nil {
		globalCache = cache.New(cacheDefaultExp*time.Minute, cacheCleanupInt*time.Minute)
	}
	return globalCache
}

func GenCacheKey(parts ...string) string {
	return strings.Join(parts, "_")
}

func SetCacheItem(name string, value interface{}, expiration time.Duration) {
	getCache().Set(name, value, expiration)
}

func GetCacheItem(name string) (interface{}, bool) {
	return getCache().Get(name)
}

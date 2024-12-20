package tvdb

import "github.com/patrickmn/go-cache"

var globalCache *cache.Cache

const (
	cacheDefaultExp = 30 // minutes
	cacheCleanupInt = 50 // minutes
)

func genCacheKey(name string) string {
	return name
}

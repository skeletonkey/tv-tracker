package tvdb

import (
	"strings"

	"github.com/patrickmn/go-cache"
)

var globalCache *cache.Cache

const (
	cacheDefaultExp = 30 // minutes
	cacheCleanupInt = 50 // minutes
)

func genCacheKey(parts ...string) string {
	return strings.Join(parts, "_")
}

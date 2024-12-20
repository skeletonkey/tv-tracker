package tvdb

import (
	"fmt"
	"net/url"
	"time"

	cache "github.com/patrickmn/go-cache"
)

const (
	maxEntriesReturned = 100
	searchCacheExp     = 60 * 24 // minutes
)

var c client

func getClient() (client, error) {
	if c.token == "" || c.refreshToken {
		if globalCache == nil {
			globalCache = cache.New(cacheDefaultExp*time.Minute, cacheCleanupInt*time.Minute)
		}
		cfg := getConfig()
		c = client{
			login: tvdbLogin{
				ApiKey: cfg.ApiKey,
				Pin:    cfg.Pin,
			},
			baseUrl:      cfg.BaseUrl,
			refreshToken: c.refreshToken,
		}

		var login loginPage
		c.makeCall("POST", c.baseUrl+"/login", c.login, &login)
		c.token = fmt.Sprintf("Bearer %s", login.Data.Token)
	}

	return c, nil
}

func Search(search string) ([]SearchResult, error) {
	cacheKey := genCacheKey("search_" + search)
	c, err := getClient()
	if err != nil {
		return nil, err
	}

	if item, found := globalCache.Get(cacheKey); found {
		fmt.Printf("Cache hit for %s\n", cacheKey)
		return item.([]SearchResult), nil
	}

	var sp SearchPage
	result := make([]SearchResult, 0)

	c.makeCall("GET", c.baseUrl+"/search?query="+url.QueryEscape(search), nil, &sp)
	result = append(result, sp.Data...)
	for sp.Links.Next != "" && sp.Links.Next != sp.Links.Self {
		c.makeCall("GET", sp.Links.Next, nil, &sp)
		result = append(result, sp.Data...)
		if len(result) >= maxEntriesReturned {
			break
		}
	}

	globalCache.Set(cacheKey, result, searchCacheExp*time.Minute)

	return result, nil
}

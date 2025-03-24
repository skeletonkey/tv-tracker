package tvdb

import (
	"fmt"
	"net/url"
	"time"

	"github.com/skeletonkey/lib-core-go/logger"
)

const (
	maxEntriesReturned = 100
	searchCacheExp     = 60 * 24 // minutes
)

var c client

func getClient() (client, error) {
	if c.token == "" || c.refreshToken {
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
		err := c.makeCall("POST", c.baseUrl+"/login", c.login, &login)
		if err != nil {
			return c, fmt.Errorf("error logging in: %s", err)
		}
		c.token = "Bearer " + login.Data.Token
	}

	return c, nil
}

func Search(search string, skip_cache bool) ([]SearchResult, error) {
	log := logger.Get()
	log.Trace().Msg("Search")
	log.Info().Str("search", search).Bool("skip_cache", skip_cache).Msg("Searching")

	cacheKey := GenCacheKey("search_" + search)
	c, err := getClient()
	if err != nil {
		return nil, err
	}

	if !skip_cache {
		if item, found := GetCacheItem(cacheKey); found {
			log.Info().Str("Cache Key", cacheKey).Msg("Cache hit")
			return item.([]SearchResult), nil
		}
	}

	var sp SearchPage
	result := make([]SearchResult, 0)

	err = c.makeCall("GET", c.baseUrl+"/search?query="+url.QueryEscape(search), nil, &sp)
	if err != nil {
		return nil, fmt.Errorf("error searching for %s: %s", search, err)
	}
	result = append(result, sp.Data...)
	for sp.Links.Next != "" && sp.Links.Next != sp.Links.Self {
		err = c.makeCall("GET", sp.Links.Next, nil, &sp)
		if err != nil {
			log.Error().Err(err).Str("Next", sp.Links.Next).Msg("Error searching for next page results")
			break
		}
		result = append(result, sp.Data...)
		if len(result) >= maxEntriesReturned {
			break
		}
	}

	SetCacheItem(cacheKey, result, searchCacheExp*time.Minute)

	return result, nil
}

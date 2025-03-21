package tvdb

import (
	"fmt"
	"net/url"
	"time"
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

func Search(search string) ([]SearchResult, error) {
	cacheKey := GenCacheKey("search_" + search)
	c, err := getClient()
	if err != nil {
		return nil, err
	}

	if item, found := GetCacheItem(cacheKey); found {
		fmt.Printf("Cache hit for %s\n", search)
		return item.([]SearchResult), nil
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
			return nil, fmt.Errorf("error searching for next page results(%s): %s", sp.Links.Next, err)
		}
		result = append(result, sp.Data...)
		if len(result) >= maxEntriesReturned {
			break
		}
	}

	SetCacheItem(cacheKey, result, searchCacheExp*time.Minute)

	return result, nil
}

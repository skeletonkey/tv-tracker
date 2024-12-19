package tvdb

import (
	"fmt"
	"net/url"
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
			baseUrl: cfg.BaseUrl,
			refreshToken: c.refreshToken,
		}

		var login loginPage
		c.makeCall("POST", c.baseUrl+"/login", c.login, &login)
		c.token = fmt.Sprintf("Bearer %s", login.Data.Token)
	}

	return c, nil
}

func Search(search string) ([]SearchResult, error) {
	c, err := getClient()
	if err != nil {
		return nil, err
	}

	var sp SearchPage
	result := make([]SearchResult, 0)

	c.makeCall("GET", c.baseUrl+"/search?query="+url.QueryEscape(search), nil, &sp)
	result = append(result, sp.Data...)
	for sp.Links.Next != "" && sp.Links.Next != sp.Links.Self {
		c.makeCall("GET", sp.Links.Next, nil, &sp)
		result = append(result, sp.Data...)
	}

	return result, nil
}

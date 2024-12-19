package tvdb

type client struct {
	login          tvdbLogin
	baseUrl, token string
	refreshToken bool
}

type tvdbLogin struct {
	ApiKey string `json:"apikey"`
	Pin    string `json:"pin"`
}

type tvdb struct {
	ApiKey  string `json:"api_key"`
	BaseUrl string `json:"base_url"`
	Pin     string `json:"pin"`
}

type loginPage struct {
	Status string `json:"status"`
	Data   struct {
		Token string `json:"token"`
	} `json:"data"`
}

type SearchPage struct {
	Status string         `json:"status"`
	Data   []SearchResult `json:"data"`
	Links  PageLinks      `json:"links"`
}

type PageLinks struct {
	Prev       string `json:"prev"`
	Self       string `json:"self"`
	Next       string `json:"next"`
	TotalItems int    `json:"total_items"`
	PageSize   int    `json:"page_size"`
}

type SearchResult struct {
	ObjectID        string     `json:"objectID"`
	Aliases         []string   `json:"aliases"`
	Country         string     `json:"country"`
	ID              string     `json:"id"`
	ImageURL        string     `json:"image_url"`
	Name            string     `json:"name"`
	FirstAirTime    string     `json:"first_air_time"`
	Overview        string     `json:"overview"`
	PrimaryLanguage string     `json:"primary_language"`
	PrimaryType     string     `json:"primary_type"`
	Status          string     `json:"status"`
	Type            string     `json:"type"`
	TvdbID          string     `json:"tvdb_id"`
	Year            string     `json:"year"`
	Slug            string     `json:"slug"`
	Overviews       Languages  `json:"overviews"`
	Translations    Languages  `json:"translations"`
	Network         string     `json:"network"`
	RemoteIDs       []RemoteID `json:"remote_ids"`
	Thumbnail       string     `json:"thumbnail"`
}

type RemoteID struct {
	ID         string `json:"id"`
	Type       int    `json:"type"`
	SourceName string `json:"sourceName"`
}

type Languages struct {
	Deu string `json:"deu"`
	Ell string `json:"ell"`
	Eng string `json:"eng"`
	Fin string `json:"fin"`
	Fra string `json:"fra"`
	Heb string `json:"heb"`
	Ita string `json:"ita"`
	Nld string `json:"nld"`
	Por string `json:"por"`
	Pt  string `json:"pt"`
	Rus string `json:"rus"`
	Spa string `json:"spa"`
	Srp string `json:"srp"`
	Swe string `json:"swe"`
	Zho string `json:"zho"`
}

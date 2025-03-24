package server

type server struct {
	CorsAllow       []string `json:"cors_allow"`               // list of URLs that are allowed to access the server
	Port            string   `json:"port"`                     // port service runs on
	ShutdownTimeout int      `json:"shutdown_timeout_seconds"` // max time (seconds) allowed for HTTP service to shutdown
}

type User struct {
	ID           int    `json:"id" validate:"omitempty"`
	UUID         string `json:"uuid"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"omitempty,email"`
	Active       bool   `json:"active"`
	Password     string `json:"password" validate:"required,min=8,max=20"`
	PasswordConf string `json:"password_confirm" validate:"omitempty,eqfield=Password"`
}

type Show struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Episode struct {
	ID            int    `json:"id"`
	UUID          string `json:"uuid"`
	ShowID        int    `json:"show_id"`
	Title         string `json:"title"`
	Season        int    `json:"season"`
	EpisodeNumber int    `json:"episode_number"`
	AirDate       string `json:"air_date"`
}

type SourceXInternal struct {
	ID         int    `json:"id"`
	ShowID     int    `json:"show_id"`
	EpisodeID  int    `json:"episode_id"`
	SourceID   int    `json:"source_id"`
	ExternalID string `json:"external_id"`
}

type Art struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	ArtTypeID int    `json:"art_type_id"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
}

type ArtType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EntityArt struct {
	ID         int  `json:"id"`
	ShowID     int  `json:"show_id"`
	EpisodeID  int  `json:"episode_id"`
	ArtID      int  `json:"art_id"`
	DefaultArt bool `json:"default_art"`
}

type Source struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	APIURL string `json:"api_url"`
	Site   string `json:"site"`
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type UserXEntityArt struct {
	ID          int `json:"id"`
	EntityArtID int `json:"entity_art_id"`
	UserID      int `json:"user_id"`
}

type UserXShow struct {
	ID      int  `json:"id"`
	UserID  int  `json:"user_id"`
	ShowID  int  `json:"show_id"`
	Rank    int  `json:"rank"`
	Visible bool `json:"visible"`
}

type UserXSource struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	SourceID int    `json:"source_id"`
	URL      string `json:"url"`
	APIToken string `json:"api_token"`
}

type UserHistory struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ShowID    int    `json:"show_id"`
	EpisodeID int    `json:"episode_id"`
	StatusID  int    `json:"status_id"`
	Timestamp string `json:"timestamp"` // or time.Time if you want to use Go's time package
}

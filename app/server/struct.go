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
	Password     string `json:"password" validate:"required,min=8,max=20"`
	PasswordConf string `json:"password_confirm" validate:"omitempty,eqfield=Password"`
}

type Show struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	ExternalID  int    `json:"external_id"`
	Thumbnail   string `json:"thumbnail"`
}

type Episode struct {
	ID            int    `json:"id"`
	UUID          string `json:"uuid"`
	ShowID        int    `json:"show_id"`
	Title         string `json:"title"`
	Season        int    `json:"season"`
	EpisodeNumber int    `json:"episode_number"`
	AirDate       string `json:"air_date"`
	ExternalID    int    `json:"external_id"`
}

type ExternalSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ExternalXInternal struct {
	ID               int    `json:"id"`
	ShowID           int    `json:"show_id"`
	EpisodeID        int    `json:"episode_id"`
	ExternalSourceID int    `json:"external_source_id"`
	ExternalID       string `json:"external_id"`
}

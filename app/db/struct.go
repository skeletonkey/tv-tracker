package db

type db struct {
	File string
}

type User struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

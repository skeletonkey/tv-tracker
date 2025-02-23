package server

type server struct {
	CorsAllow       []string `json:"cors_allow"`               // list of URLs that are allowed to access the server
	Port            string   `json:"port"`                     // port service runs on
	ShutdownTimeout int      `json:"shutdown_timeout_seconds"` // max time (seconds) allowed for HTTP service to shutdown
}

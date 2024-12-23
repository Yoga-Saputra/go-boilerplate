package config

// Redis configuration key value
// Same like redis adapter config from pkg/redis
type redis struct {
	// Host name where the DB is hosted
	Host string `json:"host"`

	// Port where the DB is listening on
	Port int `json:"port"`

	// Server username
	Username string `json:"username"`

	// Server password
	Password string `json:"password"`

	// Database to be selected after connecting to the server.
	Database int `json:"database"`

	// URL the standard format redis url to parse all other options. If this is set all other config options, Host, Port, Username, Password, Database have no effect.
	URL string `json:"url"`

	// Maximum number of retries before giving up.
	MaxRetries int `json:"maxRetries"`

	// Maximum number of socket connections.
	PoolSize int `json:"poolSize"`
}
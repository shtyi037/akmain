package config

type Config struct {
	RunMode string
	HTTP    HTTP
}

type HTTP struct {
	Port            int
	ShutdownTimeout int
}

package config

type Config struct {
	DB       DB       `json:"db"`
	Services Services `json:"services"`
}

type DB struct {
	PG_HOST   string `json:"host"`
	PG_PORT   string `json:"port"`
	PG_USER   string `json:"user"`
	PG_PASS   string `json:"pass"`
	PG_DBNAME string `json:"dbname"`
}

type Services struct {
	Auth   string `json:"auth"`
	User   string `json:"user"`
	Dialog string `json:"dialog"`
}

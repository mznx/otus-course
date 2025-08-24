package config

type Config struct {
	DB DB `json:"db"`
}

type DB struct {
	PG_HOST   string `json:"host"`
	PG_PORT   string `json:"port"`
	PG_USER   string `json:"user"`
	PG_PASS   string `json:"pass"`
	PG_DBNAME string `json:"dbname"`
}

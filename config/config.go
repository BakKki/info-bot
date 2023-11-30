package config

import "fmt"

type Config struct {
	Telegram `json:"telegram"`
	Database `json:"database"`
}

type Telegram struct {
	Token string `json:"token"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	DBName   string `json:"db_name"`
	Password string `json:"password"`
	SSLMode  string `json:"ssl_mode"`
}

func New() (*Config, error) {
	//file, err := os.ReadFile("../config/config.json")
	//if err != nil {
	//    return nil, fmt.Errorf("error while reading config file: %v", err)
	//}

	var config Config

	//config.Token = "6565885905:AAGpeOphD-WRQbHFQCVQ6A1ogdB4WS-UGv0"
	config.Token = "6891868565:AAH8EjTv9tSj_E6iDZvoXwSiMMDOTf3p1iA"
	config.Host = "127.0.0.1"
	config.Port = 5432
	config.User = "postgres"
	config.DBName = "postgres"
	config.Password = "root"
	config.SSLMode = "disable"

	//
	//if err = json.Unmarshal(file, &config); err != nil {
	//    return nil, fmt.Errorf("error while unmarshalling config file data: %v", err)
	//}

	return &config, nil
}

func (db *Database) ToDataString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", db.Host, db.Port, db.User, db.DBName, db.Password, db.SSLMode)
}

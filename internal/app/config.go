package app

type Config struct {
	Host string
	Port uint16

	DBHost     string
	DBPort     uint16
	DBUser     string
	DBPassword string
	DBName     string
	DBSslMode  string
}

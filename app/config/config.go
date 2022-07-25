package config

type Config struct {
	Database DatabaseConfigurations
}

type AccessConfigurations struct {
	ExpirationTime string
	SigningKey     string
}

type DatabaseConfigurations struct {
	User           string
	Password       string
	Host           string
	Name           string
	SSLMode        string
	MigrationsPath string
}

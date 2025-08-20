package config

import (
	"log"

	"github.com/spf13/viper"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	Server struct {
		PortalAPIPort   string `mapstructure:"portal_api_port"`
		AdminAPIPort    string `mapstructure:"admin_api_port"`
		NodeServicePort string `mapstructure:"node_service_port"`
	} `mapstructure:"server"`
	Database struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"database"`
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
	JWT struct {
		Secret      string `mapstructure:"secret"`
		ExpireHours int    `mapstructure:"expire_hours"`
	} `mapstructure:"jwt"`
}

// Cfg is the global configuration object.
var Cfg AppConfig

// Init loads the configuration from the config file.
func Init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs") // path to look for the config file in
	viper.AddConfigPath(".")         // optionally look for config in the working directory

	// Read the configuration file.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Unmarshal the configuration into the Cfg struct.
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Println("Configuration loaded successfully.")
}

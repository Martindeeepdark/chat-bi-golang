// config.config
package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	MySQL MySQLConfig `mapstructure:"mysql"`
	XfApi XfApiConfig `mapstructure:"xfapi"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type XfApiConfig struct {
	AppId     string `mapstructure:"appid"`
	ApiKey    string `mapstructure:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret"`
	HostUrl   string `mapstructure:"hostUrl"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	fmt.Printf(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func SetupDatabase() (*gorm.DB, error) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	dsn := formatDSN(config.MySQL)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func formatDSN(cfg MySQLConfig) string {
	return cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + formatPort(cfg.Port) + ")/" + cfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func formatPort(port int) string {
	return fmt.Sprint(port)
}

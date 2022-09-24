package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	RedisUrl string `mapstructure:"redis_url" yaml:"redis_url" json:"redis_url,omitempty"`
	MySqlUrl string `mapstructure:"my_sql_url" yaml:"my_sql_url" json:"my_sql_url,omitempty"`
	Env      string `mapstructure:"env" yaml:"env" json:"env,omitempty"`
}

func Init() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("__", "."))
	viper.AutomaticEnv()

	cfg := &Config{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

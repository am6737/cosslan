package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server      Server      `mapstructure:"server"`
	Persistence Persistence `mapstructure:"persistence"`
}

type Server struct {
	Port     string `mapstructure:"port"`
	Timeout  int    `mapstructure:"timeout"`
	LogLevel string `mapstructure:"log_level"`
}

type Persistence struct {
	Enabled bool   `yaml:"enabled"`
	Url     string `yaml:"url"`
	Type    string `yaml:"type"`
	DB      string `yaml:"db"`
}

// LoadConfig 加载配置文件并解析成 Config 结构体
func LoadConfig(path string) (Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("无法读取配置文件: %w", err)
	}

	// 解析配置文件到结构体
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("无法解析配置文件: %w", err)
	}

	return config, nil
}

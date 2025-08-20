package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var globalConfig *Config = nil

type Config struct {
	Database Database
	Server   Server
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Option   string
}

type Server struct {
	Port int
}

func LoadConfig() (*Config, error) {
	if globalConfig != nil {
		return globalConfig, nil
	}

	// Tìm đường dẫn gốc của dự án
	pwd, rootPathErr := findProjectRoot()
	if rootPathErr != nil {
		return nil, rootPathErr
	}

	suffix := "-local"
	viper.SetConfigName("config" + suffix) // Tên file config (không bao gồm phần mở rộng)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(pwd + "/config") // Tìm file config-'env'.yaml trong thư mục config

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("không thể đọc file config: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("không thể parse config: %w", err)
	}

	globalConfig = &config

	return &config, nil
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// Kiểm tra xem có file go.mod hay không (dùng để xác định thư mục gốc của dự án)
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// Di chuyển lên thư mục cha
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("can't file")
		}

		dir = parent
	}
}

func (cfg *Config) GetDatabaseUri() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.Option)
}

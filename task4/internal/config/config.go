package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/goccy/go-yaml"
)

// Configuration 结构体对应配置文件内容
type Configuration struct {
	Service struct {
		Port string `yaml:"port"`
	} `yaml:"service"`

	MySQL struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`

	Log struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"log"`
}

// 注意 * 和 不带*的区别
var (
	config *Configuration
	once   sync.Once
)

// InitConfig 初始化配置（单例模式）
func InitConfig(configPath string) {
	once.Do(func() {
		config = &Configuration{}
		err := loadConfig(configPath)
		if err != nil {
			log.Fatalf("yaml配置加载失败: %v", err)
		}
	})
}

// GetConfig 获取全局配置实例
func GetConfig() *Configuration {
	return config
}

// 加载并解析配置文件
func loadConfig(configPath string) error {
	// 读取文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败：%w", err)
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("解析配置文件失败：%w", err)
	}

	return nil
}

package configs

import (
	viper "github.com/spf13/viper"
	"hero/pkg/logger"
	"os"
	"path/filepath"
	"strings"
)

var (
	cacheMap map[string]string
	EnvPath  string
)

func init() {
	viper.SetConfigType("yaml")
	cacheMap = make(map[string]string)
	if EnvPath = os.Getenv("env"); EnvPath == "" {
		EnvPath = "local"
	}
	logger.Print("env", EnvPath)
}

func Get(target string) string {
	return getCache(target)
}

func setPath(path string) {
	if EnvPath != "" {
		path = EnvPath + "/" + path
	}

	dir, err := filepath.Abs("./configs/" + path)
	if err != nil {
		logger.Print("config init", err.Error())
	}
	viper.Reset()
	viper.AddConfigPath(dir)
	logger.Print("set config path", dir)
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		logger.Print("Fatal error config file", err.Error())
	}
}

func getCache(target string) string {
	if val, ok := cacheMap[target]; ok {
		return val
	}
	targets := strings.Split(target, ".")
	if len(targets) != 2 {
		return ""
	}
	path := targets[0]
	key := targets[1]
	setPath(path)
	value := viper.GetString(key)
	cacheMap[target] = value
	return value
}

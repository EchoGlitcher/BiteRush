package cmd

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
)

type Config struct {
	Server   Server                `yaml:"server"`
	Database DatabaseConfiguration `yaml:"db"`
}

type Server struct {
	Port int `yaml:"port"`
}

type DatabaseConfiguration struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

func (cfg DatabaseConfiguration) CreateDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}

func ResolveConfig() Config {

	klog.InitFlags(nil)
	klog.EnableContextualLogging(true)

	var (
		configFile = flag.String("cfg", "config.yaml", "Configuration file")
	)
	flag.Parse()

	file, err := os.Open(*configFile)
	if err != nil {
		klog.Fatalf("cannot read config file: %v", err)
	}
	defer file.Close()

	var config Config
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		klog.Fatalf("cannot unmarshal config yaml: %v", err)
	}

	return config
}

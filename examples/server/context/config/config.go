package config

import (
	"encoding/json"
	"strings"
	"time"
)

type serverConfig struct {
	HTTP struct {
		Listen       string        `default:":8080" field:"listen" json:"listen" yaml:"listen" cli:"http-listen" env:"SERVER_HTTP_LISTEN"`
		ReadTimeout  time.Duration `default:"120s" field:"read_timeout" json:"read_timeout" yaml:"read_timeout" env:"SERVER_HTTP_READ_TIMEOUT"`
		WriteTimeout time.Duration `default:"120s" field:"write_timeout" json:"write_timeout" yaml:"write_timeout" env:"SERVER_HTTP_WRITE_TIMEOUT"`
	} `json:"HTTP" yaml:"HTTP"`
	GRPC struct {
		Listen   string        `default:"tcp://:8081" field:"listen" json:"listen" yaml:"listen" cli:"grpc-listen" env:"SERVER_GRPC_LISTEN"`
		Timeout  time.Duration `default:"120s" field:"timeout" json:"timeout" yaml:"timeout" env:"SERVER_GRPC_TIMEOUT"`
		CertFile string        `field:"cert_file" json:"cert_file" yaml:"cert_file" env:"SERVER_GRPC_CERTFILE"`
		KeyFile  string        `field:"key_file" json:"key_file" yaml:"key_file" env:"SERVER_GRPC_KEYFILE"`
	} `json:"GRPC" yaml:"GRPC"`
}

// ConfigType contains all application options
type ConfigType struct {
	ServiceName    string `json:"service_name" yaml:"service_name" env:"SERVICE_NAME" default:"api"`
	DatacenterName string `json:"datacenter_name" yaml:"datacenter_name" env:"DC_NAME" default:"??"`
	Hostname       string `json:"hostname" yaml:"hostname" env:"HOSTNAME" default:""`
	Hostcode       string `json:"hostcode" yaml:"hostcode" env:"HOSTCODE" default:""`

	LogAddr  string `json:"log_addr" default:"" env:"LOG_ADDR"`
	LogLevel string `json:"log_level" default:"debug" env:"LOG_LEVEL"`

	Server serverConfig `json:"server" yaml:"server"`
}

// String implementation of Stringer interface
func (cfg *ConfigType) String() (res string) {
	if data, err := json.MarshalIndent(cfg, "", "  "); err != nil {
		res = `{"error":"` + err.Error() + `"}`
	} else {
		res = string(data)
	}
	return res
}

// IsDebug mode
func (cfg *ConfigType) IsDebug() bool {
	return strings.EqualFold(cfg.LogLevel, "debug")
}

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type FilterConfig struct {
	Type   string
	Config map[string]interface{}
}

type ListenerConfig struct {
	Name           string
	Address        string
	BindToPort     bool           `json:"bind_port"`
	NetworkFilters []FilterConfig `json:"network_filters"`
	StreamFilters  []FilterConfig `json:"stream_filters"`
}

type ServerConfig struct {
	LoggerPath string `json:"logger_path"`
	LoggerLevel  string `json:"logger_level"`
	// only used in http2 case
	DisableConnIo bool `json:"disable_conn_io"`

	Listeners []ListenerConfig
}

type HostConfig struct {
	Address  string
	Hostname string
	Weight   uint32
}

type HealthCheckConfig struct {
	Timeout            time.Duration
	HealthyThreshold   uint32 `json:"healthy_threshold"`
	UnhealthyThreshold uint32 `json:"unhealthy_threshold"`
	Interval           time.Duration
	IntervalJitter     time.Duration `json:"interval_jitter"`
	CheckPath          string
	ServiceName        string
}

type ClusterConfig struct {
	Name                 string
	Type                 string
	LbType               string `json:"lb_type"`
	MaxRequestPerConn    uint64 `json:"max_request_per_conn"`
	ConnBufferLimitBytes uint32
	HealthCheck          HealthCheckConfig `json:"healthcheck"`
	Hosts                []HostConfig
}

type ClusterManagerConfig struct {
	Clusters []ClusterConfig `json:"clusters"`
}

type MOSNConfig struct {
	Servers        []ServerConfig       `json:"servers"`
	ClusterManager ClusterManagerConfig `json:"cluster_manager"`
	//tracing config
}

func Load(path string) *MOSNConfig {
	fmt.Println("load config from : " + path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var config MOSNConfig
	json.Unmarshal(content, &config)
	return &config
}

//func main() {
//	config := Load("D:/workspace/go_files/src/gitlab.alipay-inc.com/afe/mosn/pkg/mosn/mosn_config.json")
//	fmt.Printf("config : %+v", config)
//
//}
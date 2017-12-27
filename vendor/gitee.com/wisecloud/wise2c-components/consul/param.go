package consul

import (
	"gitee.com/wisecloud/wise2c-components/network"
	"gitee.com/wisecloud/wise2c-components/tool"
	"github.com/Sirupsen/logrus"
)

var (
	DeRegisterCriticalServiceAfter = "30m"
	Interval                       = "15s"
	TTL                            = "10s"
	DefaultSchema                  = "http://"
)

type ConsulParam struct {
	ServerURL                      string //eg: consul:8500
	ServiceName                    string
	RegistryIp                     string
	RegistryID                     string
	RegistryPort                   int
	HealthCheckURL                 string
	DeRegisterCriticalServiceAfter string
	Schema                         string
	Interval                       string
	TTL                            string
	Timeout                        string
	Tags                           []string
	EnableTagOverride              bool
}

func (param *ConsulParam) Default() {
	if param == nil {
		logrus.Printf("consul param is nil!")
		return
	}

	if len(tool.Trim(param.DeRegisterCriticalServiceAfter)) == 0 {
		param.DeRegisterCriticalServiceAfter = DeRegisterCriticalServiceAfter
	}

	if len(tool.Trim(param.Interval)) == 0 {
		param.Interval = Interval
	}

	if len(tool.Trim(param.TTL)) == 0 {
		param.TTL = TTL
	}

	if len(tool.Trim(param.Schema)) == 0 {
		param.Schema = DefaultSchema
	}

	if len(tool.Trim(param.RegistryIp)) == 0 {
		param.RegistryIp = network.GetRegistryIp()
		logrus.Printf("registry ip from consul components is %s", param.RegistryIp)
	}
	return
}

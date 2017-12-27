package helper

import (
	"fmt"
	"os"
	"strings"

	"github.com/toolkits/net"
	"gitee.com/wisecloud/wise2c-components/consul"
	"github.com/astaxie/beego/logs"
)

const (
	CONSUL_SERVICE_MODEL = "model"
)

var ConsulClient *consul.ConsulClient

func RegistryConsul(serviceName string, registryPort int, healthCheckURL string) error {
	var err error

	e := EnvVar{}.Get()
	if len(e.ConsulURL) == 0 || len(e.ConsulPort) == 0 {
		logs.Error("invalid CONSUL_PORT or CONSUL_URL")
		os.Exit(-1)
	}
	consulAddr := fmt.Sprintf("%s:%s", e.ConsulURL, e.ConsulPort)

	ConsulClient, err = consul.NewConsulClient(
		&consul.ConsulParam{
			ServerURL:      consulAddr,
			ServiceName:    serviceName,
			RegistryPort:   registryPort,
			HealthCheckURL: healthCheckURL,
		})
	if err != nil {
		logs.Error("new consul client failed, error info is ", err.Error())
		return err
	}
	ConsulClient.AutoRegistry = true
	ConsulClient.AutoAgentRegistry()

	return nil
}

func GetServiceAddress(serviceName string) (string, error) {
	if ConsulClient == nil {
		logs.Error("ConsulClient is NULL!")
		return "", fmt.Errorf("invalid ConsulClient (null)! ")
	}

	address := ConsulClient.GetServiceAddress(serviceName)
	return address, nil
}

func GetMyIPAddr() string {
	ips, err := net.IntranetIP()
	if err != nil {
		logs.Error("can't get ip list!", err.Error())
		return ""
	}
	return ips[0]
}

func GetMyIPIdentity() string {
	return strings.Replace(GetMyIPAddr(), ".", "_", -1)
}

type EnvVar struct {
	DbUrl      string
	DbPort     string
	DbUser     string
	DbPwd      string
	ConsulURL  string
	ConsulPort string
	HostName   string
}

func (e EnvVar) Get() EnvVar {
	e.DbUrl = os.Getenv("DB_URL")
	e.DbPort = os.Getenv("DB_PORT")
	e.DbUser = os.Getenv("DB_USER")
	e.DbPwd = os.Getenv("DB_PASSWORD")
	e.ConsulURL = os.Getenv("CONSUL_URL")
	e.ConsulPort = os.Getenv("CONSUL_PORT")
	e.HostName = os.Getenv("MY_NODE_NAME")

	return e
}

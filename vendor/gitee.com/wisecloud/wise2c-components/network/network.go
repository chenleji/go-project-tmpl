package network

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/mojo-zd/go-library/http"
	"github.com/toolkits/net"
)

var (
	RANCHER_VERSION   = "http://rancher-metadata/latest"
	RANCHER_MANAGE_IP = fmt.Sprintf("%s/%s", RANCHER_VERSION, "self/container/primary_ip")
	RANCHER_HOST_IP   = fmt.Sprintf("%s/%s", RANCHER_VERSION, "self/host/agent_ip")
)

var (
	DEBUG = false
)

func GetRegistryIp() (ip string) {
	if len(os.Getenv("KUBERNETES_PORT")) > 0 || DEBUG {
		ip = internalIp()
		return
	}
	ip = getRancherManageIP()
	return
}

func internalIp() string {
	ips, err := net.IntranetIP()
	if err != nil {
		logrus.Errorf("get inner ip failed, error info is %s\n", err.Error())
		return ""
	}
	return ips[0]
}

func getRancherManageIP() (ip string) {
	URL := RANCHER_MANAGE_IP
	if len(os.Getenv("PROFILE")) > 0 {
		URL = RANCHER_HOST_IP
	}

	response := http.NewHttpClient().BuildRequestInfo(&http.RequestInfo{URL: URL}).Get()
	if response.Error != nil {
		logrus.Printf("get rancher manager ip failed, error info is %s\n", response.Error.Error())
		return
	}

	ip = string(response.Result)
	logrus.Printf("rancher manager ip is %s ", ip)
	return
}

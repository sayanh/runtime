package utils

import (
	yaml "github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

type RuntimeInfo struct {
	runtimeConfig     *corev1.ConfigMap
	RegistryInfo      DockerInfo
	AvailableRuntimes []RuntimesSupported
	ServiceAccount    string
}

type DockerInfo struct {
	DockerRegistry string
	DockerAccount  string
}

type RuntimesSupported struct {
	ID             string `yaml:"ID"`
	DockerFileName string `yaml:"DockerFileName"`
}

func New(config *corev1.ConfigMap) *RuntimeInfo {
	var dockerInfo DockerInfo
	var ars []RuntimesSupported
	sa := ""
	return &RuntimeInfo{
		runtimeConfig:     config,
		RegistryInfo:      dockerInfo,
		AvailableRuntimes: ars,
		ServiceAccount:    sa,
	}
}

func (ri *RuntimeInfo) ReadConfigMap() {
	ri.RegistryInfo.DockerAccount = ri.runtimeConfig.Data["dockerHubAccount"]
	ri.RegistryInfo.DockerRegistry = ri.runtimeConfig.Data["dockerRegistry"]
	if runtimeImages, ok := ri.runtimeConfig.Data["runtimes"]; ok {
		err := yaml.Unmarshal([]byte(runtimeImages), &ri.AvailableRuntimes)
		if err != nil {
			logrus.Fatalf("Unable to get the supported runtimes: %v", err)
		}
	}
	ri.ServiceAccount = ri.runtimeConfig.Data["serviceAccountName"]
}

func (ri *RuntimeInfo) GetDockerFileName(runtime string) string {
	result := ""
	for _, runtimeInf := range ri.AvailableRuntimes {
		if runtimeInf.ID == runtime {
			result = runtimeInf.DockerFileName
		}
	}
	if runtime == "" {
		logrus.Infof("Unable to find the docker file for runtime: %v", runtime)
	}
	return result
}

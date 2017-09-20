package main

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
	"github.com/clickyab/services/assert"
	"k8s.io/client-go/kubernetes"
)

func main() {
	// get cluster config
	config, err := rest.InClusterConfig()
	assert.Nil(err)
	logrus.Warn(config)

	clientSet,err:=kubernetes.NewForConfig(config)
	assert.Nil(err)
	logrus.Warn(clientSet)
}

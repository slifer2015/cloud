package main

import (
	"fmt"

	"github.com/clickyab/services/assert"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// get cluster config
	config, err := rest.InClusterConfig()
	assert.Nil(err)
	clientSet, err := kubernetes.NewForConfig(config)
	assert.Nil(err)
	nodes, err := clientSet.CoreV1().Nodes().List(metav1.ListOptions{})
	assert.Nil(err)
	for _, n := range nodes.Items {
		for _, s := range n.Status.Addresses {
			logrus.Warn(fmt.Sprintf("Address: %s, Type: %s", s.Address, s.Type))
		}
		for _, s := range n.Status.Conditions {
			logrus.Warn(fmt.Sprintf("Type: %s, status: %s, Message: %s", s.Type, s.Message, s.Status))
		}
		logrus.Info("//////////////////////////")

	}
}

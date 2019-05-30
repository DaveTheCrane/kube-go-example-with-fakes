package main

import (
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	v12 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v1b1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	restclient "k8s.io/client-go/rest"
)

type k8sSystem struct {
	deployments v12.DeploymentInterface
	services v1.ServiceInterface
	ingress v1b1.IngressInterface
}

func productionK8sSystem(config *restclient.Config) *k8sSystem {
	return &k8sSystem{
		deployments: deploymentsClient(config),
		services: servicesClient(config),
		ingress: ingressClient(config),
	}
}

func deploymentsClient(config *restclient.Config) v12.DeploymentInterface {
	clientset := kubernetes.NewForConfigOrDie(config)
	return clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
}

func servicesClient(config *restclient.Config) v1.ServiceInterface {
	coreV1Client := v1.NewForConfigOrDie(config)
	return coreV1Client.Services(apiv1.NamespaceDefault)
}

func ingressClient(config *restclient.Config) v1b1.IngressInterface {
	extV1B1Client := v1b1.NewForConfigOrDie(config)
	return extV1B1Client.Ingresses("default")

}



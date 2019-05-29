package main

import (
	v1b1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	apiv1b1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
)

func ingressClient(config *restclient.Config) v1b1.IngressInterface {
	extV1B1Client := v1b1.NewForConfigOrDie(config)
	return extV1B1Client.Ingresses("default")

}

func createIngress(config *restclient.Config, ingress *apiv1b1.Ingress) *apiv1b1.Ingress {
	ingressInstance, err := ingressClient(config).Create(ingress)
	if err != nil {
		panic(err)
	}
	return ingressInstance
}

func listIngress(config *restclient.Config) *apiv1b1.IngressList {
	ingressList, err := ingressClient(config).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return ingressList
}

func deleteIngress(config *restclient.Config, name string) {
	deletePolicy := metav1.DeletePropagationForeground
	err := ingressClient(config).Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	});
	if err != nil {
		panic(err)
	}
}


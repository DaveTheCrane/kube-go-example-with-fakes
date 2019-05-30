package main

import (
	apiv1b1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createIngress(sys *k8sSystem, ingress *apiv1b1.Ingress) *apiv1b1.Ingress {
	ingressInstance, err := sys.ingress.Create(ingress)
	if err != nil {
		panic(err)
	}
	return ingressInstance
}

func listIngress(sys *k8sSystem) *apiv1b1.IngressList {
	ingressList, err := sys.ingress.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return ingressList
}

func deleteIngress(sys *k8sSystem, name string) {
	deletePolicy := metav1.DeletePropagationForeground
	err := sys.ingress.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	});
	if err != nil {
		panic(err)
	}
}


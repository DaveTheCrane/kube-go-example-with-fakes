package main

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createService(sys *k8sSystem, service *apiv1.Service) *apiv1.Service {
	serviceInstance, err := sys.services.Create(service)
	if err != nil {
		panic(err)
	}
	return serviceInstance
}

func listServices(sys *k8sSystem) *apiv1.ServiceList {
	serviceList, err := sys.services.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return serviceList
}

func deleteService(sys *k8sSystem, name string) {
	deletePolicy := metav1.DeletePropagationForeground
	err := sys.services.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	});
	if err != nil {
		panic(err)
	}
}


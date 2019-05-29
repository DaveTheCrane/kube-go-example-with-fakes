package main

import (
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
)

func servicesClient(config *restclient.Config) v1.ServiceInterface {
	coreV1Client := v1.NewForConfigOrDie(config)
	return coreV1Client.Services(apiv1.NamespaceDefault)
}

func createService(config *restclient.Config, service *apiv1.Service) *apiv1.Service {
	serviceInstance, err := servicesClient(config).Create(service)
	if err != nil {
		panic(err)
	}
	return serviceInstance
}

func listServices(config *restclient.Config) *apiv1.ServiceList {
	serviceList, err := servicesClient(config).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return serviceList
}

func deleteService(config *restclient.Config, name string) {
	deletePolicy := metav1.DeletePropagationForeground
	err := servicesClient(config).Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	});
	if err != nil {
		panic(err)
	}
}


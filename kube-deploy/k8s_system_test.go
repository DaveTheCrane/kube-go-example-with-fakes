package main

import (
	apiv1 "k8s.io/api/core/v1"
	kf "k8s.io/client-go/kubernetes/fake"
	v12 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	v1b1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

func testK8sSystem() *k8sSystem {
	sys := &k8sSystem{
		deployments: testDeploymentsClient(),
		services: testServicesClient(),
		ingress: testIngressClient(),
	}
	return sys
}

func testDeploymentsClient() v12.DeploymentInterface {
	clientset := kf.NewSimpleClientset()
	return clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
}

func testServicesClient() v1.ServiceInterface {
	clientset := kf.NewSimpleClientset()
	coreV1Client := clientset.CoreV1()
	return coreV1Client.Services(apiv1.NamespaceDefault)
}

func testIngressClient() v1b1.IngressInterface {
	clientset := kf.NewSimpleClientset()
	extV1B1Client := clientset.ExtensionsV1beta1()
	return extV1B1Client.Ingresses("default")
}


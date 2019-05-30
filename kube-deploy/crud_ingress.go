package main

import (
	"fmt"
	apiv1b1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
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

func updateIngressExternalPath(sys *k8sSystem, name string, newPath string) {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := sys.ingress.Get(name, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Ingress: %v", getErr))
		}

		onlyPath := result.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0]
		currentPath := onlyPath.Path
		if currentPath != newPath {
			fmt.Println("ext path for %s: change from %s to %s", name, currentPath, newPath)
			result.Spec.Rules = ingressLoadBalancerRules(result.Spec.Rules[0].Host, newPath, onlyPath.Backend.ServiceName, int(onlyPath.Backend.ServicePort.IntVal))
			_, updateErr := sys.ingress.Update(result)
			return updateErr
		} else {
			fmt.Println("ext path for %s: no change %s to %s", name, currentPath, newPath)
			return nil
		}
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
}


package main

import(
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func createDeployment(sys *k8sSystem, deployment *appsv1.Deployment) *appsv1.Deployment {
	deploymentInstance, err := sys.deployments.Create(deployment)
	if err != nil {
		panic(err)
	}
	return deploymentInstance
}

func listDeployments(sys *k8sSystem) *appsv1.DeploymentList {
	deploymentList, err := sys.deployments.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return deploymentList
}

func deleteDeployment(sys *k8sSystem, name string) {
	deletePolicy := metav1.DeletePropagationForeground
	err := sys.deployments.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	});
	if err != nil {
		panic(err)
	}
}

func updateDeploymentSize(sys *k8sSystem, name string, newReplicaCount int) {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := sys.deployments.Get(name, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		currentCount := result.Spec.Replicas
		newCount := int32(newReplicaCount)
		if *currentCount != newCount {
			result.Spec.Replicas = int32Ptr(newCount)
			_, updateErr := sys.deployments.Update(result)
			return updateErr
		} else {
			return nil
		}
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
}




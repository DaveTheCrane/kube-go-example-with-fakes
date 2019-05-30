package main

import(
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// 	// Update Deployment
// 	prompt()
// 	fmt.Println("Updating deployment...")
// 	//    You have two options to Update() this Deployment:
// 	//
// 	//    1. Modify the "deployment" variable and call: Update(deployment).
// 	//       This works like the "kubectl replace" command and it overwrites/loses changes
// 	//       made by other clients between you Create() and Update() the object.
// 	//    2. Modify the "result" returned by Get() and retry Update(result) until
// 	//       you no longer get a conflict error. This way, you can preserve changes made
// 	//       by other clients between Create() and Update(). This is implemented below
// 	//			 using the retry utility package included with client-go. (RECOMMENDED)
// 	//
// 	// More Info:
// 	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency

// 	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
// 		// Retrieve the latest version of Deployment before attempting update
// 		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
// 		result, getErr := deploymentsClient.Get("demo-deployment", metav1.GetOptions{})
// 		if getErr != nil {
// 			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
// 		}

// 		result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
// 		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
// 		_, updateErr := deploymentsClient.Update(result)
// 		return updateErr
// 	})
// 	if retryErr != nil {
// 		panic(fmt.Errorf("Update failed: %v", retryErr))
// 	}
// 	fmt.Println("Updated deployment...")





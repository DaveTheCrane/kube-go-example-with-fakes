package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	v1b1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1b1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	restclient "k8s.io/client-go/rest"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {

	//specifications
	labels := map[string]string{
		"app": "hello-world",
	}
	containerPort := 8090
	deployment := deploymentSpec("hello-deployment",
		"hello-go", "0.0.1",
		"hello-web-container", containerPort,
		5, labels)
	service := serviceSpec("hello-service", containerPort, labels)
	ingress := ingressLoadBalancerSpec("hello-lb", "mini.kube.io", "/hi/(.*)", "/hello/$1", "hello-service", containerPort)

	config := resolveConfig()

	// Create Deployment
	fmt.Println("Creating deployment...")
	deploymentInstance := createDeployment(config, deployment)
	fmt.Printf("Created deployment %q.\n", deploymentInstance.GetObjectMeta().GetName())

	// Create Service To Wrap Deployment
	fmt.Println("Wrapping in service...")
	serviceInstance := createService(config, service)
	fmt.Printf("Wrapped service %q. \n", serviceInstance.Name)

	// Create ingress service
	fmt.Println("Creating Load Balancer...")
	lbInstance := createLoadBalancer(config, ingress)
	fmt.Printf("Created Load Balancer %q. \n", lbInstance.Name)
}


func resolveConfig() *restclient.Config {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	return config
}

func createDeployment(config *restclient.Config, deployment *appsv1.Deployment) *appsv1.Deployment {
	clientset := kubernetes.NewForConfigOrDie(config)
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deploymentInstance, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}

	return deploymentInstance
}

func createService(config *restclient.Config, service *apiv1.Service) *apiv1.Service {
	coreV1Client := v1.NewForConfigOrDie(config)
	serviceInstance, err := coreV1Client.Services(apiv1.NamespaceDefault).Create(service)
	if err != nil {
		panic(err)
	}
	return serviceInstance
}

func createLoadBalancer(config *restclient.Config, ingress *apiv1b1.Ingress) *apiv1b1.Ingress {

	extV1B1Client := v1b1.NewForConfigOrDie(config)
	ingressInstance, err := extV1B1Client.Ingresses("default").Create(ingress)
	if err != nil {
		panic(err)
	}
	return ingressInstance

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

// 	// List Deployments
// 	prompt()
// 	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
// 	list, err := deploymentsClient.List(metav1.ListOptions{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, d := range list.Items {
// 		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
// 	}

// 	// Delete Deployment
// 	prompt()
// 	fmt.Println("Deleting deployment...")
// 	deletePolicy := metav1.DeletePropagationForeground
// 	if err := deploymentsClient.Delete("demo-deployment", &metav1.DeleteOptions{
// 		PropagationPolicy: &deletePolicy,
// 	}); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Deleted deployment.")
// }

// func prompt() {
// 	fmt.Printf("-> Press Return key to continue.")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		break
// 	}
// 	if err := scanner.Err(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println()
// }


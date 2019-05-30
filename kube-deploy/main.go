package main

import (
	"flag"
	"fmt"
	"path/filepath"

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

	createMode := flag.Bool("create", false, "Create deployment, service and ingress in kubernetes cluster")
	listMode := flag.Bool("list", false, "List current deployment, service and ingress in kubernetes cluster")
	deleteMode := flag.Bool("delete", false, "Delete current deployment, service and ingress in kubernetes cluster")

	flag.Parse()
	//flags to set...
	/*
	docker image name and tag
	root name for everything
	exposed path fragment
	 */
	if (*createMode){
		create_all()
	}else if (*listMode){
		list_all()
	}else if (*deleteMode){
		delete_all()
	}else{
		flag.Usage()
	}
}

func create_all() {
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

	sys := resolveSys()

	// Create Deployment
	fmt.Println("Creating deployment...")
	deploymentInstance := createDeployment(sys, deployment)
	fmt.Printf("Created deployment %q.\n", deploymentInstance.GetObjectMeta().GetName())

	// Create Service To Wrap Deployment
	fmt.Println("Wrapping in service...")
	serviceInstance := createService(sys, service)
	fmt.Printf("Wrapped service %q. \n", serviceInstance.Name)

	// Create ingress service
	fmt.Println("Creating Load Balancer...")
	lbInstance := createIngress(sys, ingress)
	fmt.Printf("Created Load Balancer %q. \n", lbInstance.Name)
}


func list_all(){
	sys := resolveSys()

	listD := listDeployments(sys)
	fmt.Println("===DEPLOYMENTS===")
	for i:=0; i < len(listD.Items); i++ {
		it := listD.Items[i]
		fmt.Printf("%s %s %s\n", it.Name, it.UID, it.CreationTimestamp)
	}

	listS := listServices(sys)
	fmt.Println("\n===SERVICES===")
	for i:=0; i < len(listS.Items); i++ {
		it := listS.Items[i]
		fmt.Printf("%s %s %s\n", it.Name, it.UID, it.CreationTimestamp)
	}

	listI := listIngress(sys)
	fmt.Println("\n===INGRESSES===")
	for i:=0; i < len(listI.Items); i++ {
		it := listI.Items[i]
		fmt.Printf("%s %s %s\n", it.Name, it.UID, it.CreationTimestamp)
	}
}

func delete_all(){
	sys := resolveSys()

	deleteDeployment(sys, "hello-deployment")
	fmt.Println("deleted deployment")
	deleteService(sys, "hello-service")
	fmt.Println("deleted service")
	deleteIngress(sys, "hello-lb")
	fmt.Println("deleted ingress")
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

func resolveSys() *k8sSystem {
	return productionK8sSystem(resolveConfig())
}




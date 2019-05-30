package main

import "fmt"

type operations struct {
	sys *k8sSystem
	createMode bool
	listMode bool
	deleteMode bool
}

func (ops *operations) execute() bool {
	if (ops.createMode){
		ops.create_all()
	}else if (ops.listMode){
		ops.list_all()
	}else if (ops.deleteMode){
		ops.delete_all()
	}else{
		return false
	}
	return true
}

func (ops *operations) create_all() {

	sys := ops.sys

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


func (ops *operations) list_all() {

	sys := ops.sys

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

func (ops *operations) delete_all() {

	sys := ops.sys

	deleteDeployment(sys, "hello-deployment")
	fmt.Println("deleted deployment")
	deleteService(sys, "hello-service")
	fmt.Println("deleted service")
	deleteIngress(sys, "hello-lb")
	fmt.Println("deleted ingress")
}

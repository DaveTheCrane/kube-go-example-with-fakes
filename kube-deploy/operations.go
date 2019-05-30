package main

import "fmt"

type operations struct {
	sys *k8sSystem

	createMode bool
	listMode bool
	deleteMode bool

	prefix string
	imageName string
	imageTag string

	hostname string
	externalPath string

	replicas int

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

func (ops *operations) label(s string) string {
	return fmt.Sprintf("%s-%s", ops.prefix, s)
}

func (ops *operations) create_all() {

	sys := ops.sys

	//specifications
	labels := map[string]string{
		"app": ops.label("world"),
	}
	containerPort := 8090
	deployment := deploymentSpec(ops.label("deployment"),
		ops.imageName, ops.imageTag,
		ops.label("web-container"), containerPort,
		int32(ops.replicas), labels)
	service := serviceSpec(ops.label("service"), containerPort, labels)
	ingress := ingressLoadBalancerSpec(ops.label("lb"), ops.hostname, fmt.Sprintf("/%s/(.*)", ops.externalPath), "/hello/$1", ops.label("service"), containerPort)

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

	deleteDeployment(sys, ops.label("deployment"))
	fmt.Println("deleted deployment")
	deleteService(sys, ops.label("service"))
	fmt.Println("deleted service")
	deleteIngress(sys, ops.label("lb"))
	fmt.Println("deleted ingress")
}

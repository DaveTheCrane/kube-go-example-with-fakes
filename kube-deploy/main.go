package main

import (
	"flag"
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
	updateMode := flag.Bool("update", false, "Update selected features of current deployment, currently only external path and replica count supported")

	prefix := flag.String("prefix", "hello", "prefix used to identify artefacts in the kubernetes system")

	imageName := flag.String("image-name", "hello-go", "docker image name for service to be deployed")
	imageTag := flag.String("image-tag", "0.0.1", "docker image tag for service to be deployed")

	hostname := flag.String("host-name", "mini.kube.io", "hostname via which service is exposed in the load balancer")
	externalPath := flag.String("ext-path", "hi", "path used by load balancer rewrite rule to access the service")

	replicas := flag.Int("replicas", 5, "number of load-balanced replicas to deploy")

	flag.Parse()
	//flags to set...
	/*
	docker image name and tag
	root name for everything
	exposed path fragment
	 */

	ops := &operations{
		sys: resolveSys(),
		createMode: *createMode,
		listMode: *listMode,
		deleteMode: *deleteMode,
		updateMode: *updateMode,

		prefix: *prefix,

		imageName: *imageName,
		imageTag: *imageTag,

		hostname: *hostname,
		externalPath: *externalPath,

		replicas: *replicas,
	}

	if (!ops.execute()) {
		flag.Usage()
	}
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




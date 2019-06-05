package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperationRequiresAnAction(t *testing.T) {
	ops := &operations{
		sys: testK8sSystem(),
		createMode: false,
		listMode: false,
		deleteMode: false,
		updateMode: false,
	}
	execResult := ops.execute()

	assert.False(t, execResult)
}

func TestOperationCanCreate(t *testing.T) {

	testK8s := testK8sSystem()
	ops := &operations{
		sys: testK8s,
		createMode: true,
		listMode: false,
		deleteMode: false,
		updateMode: false,

		prefix: "testme",
		imageName: "testImage",
		imageTag: "1.2.3",

		hostname: "kube.test",
		externalPath: "hello-world",

		replicas: 42,
	}
	execResult := ops.execute()

	assert.True(t, execResult)

	deployments := listDeployments(testK8s)
	assert.Len(t, deployments.Items, 1)
	deployment := deployments.Items[0]
	assert.EqualValues(t, "testme-deployment", deployment.Name)
	assert.EqualValues(t, int32(42), *deployment.Spec.Replicas)

	services := listServices(testK8s)
	assert.Len(t, services.Items, 1)
	service := services.Items[0]
	assert.EqualValues(t, "testme-service", service.Name)

	ingresses := listIngress(testK8s)
	assert.Len(t, ingresses.Items, 1)
	ingress := ingresses.Items[0]
	assert.EqualValues(t, "testme-lb", ingress.Name)
}

func TestOperationCanUpdate(t *testing.T) {

	testK8s := testK8sSystem()
	ops := &operations{
		sys: testK8s,
		createMode: true,
		listMode: false,
		deleteMode: false,
		updateMode: false,

		prefix: "testme",
		imageName: "testImage",
		imageTag: "1.2.3",

		hostname: "kube.test",
		externalPath: "hello-world",

		replicas: 42,
	}
	ops.execute()
	ops.replicas = 7
	ops.createMode = false
	ops.updateMode = true
	execResult := ops.execute()

	assert.True(t, execResult)

	deployments := listDeployments(testK8s)
	assert.Len(t, deployments.Items, 1)
	deployment := deployments.Items[0]
	assert.EqualValues(t, "testme-deployment", deployment.Name)
	assert.EqualValues(t, int32(7), *deployment.Spec.Replicas)
}

func TestOperationCanDelete(t *testing.T) {

	testK8s := testK8sSystem()
	ops := &operations{
		sys: testK8s,
		createMode: true,
		listMode: false,
		deleteMode: false,
		updateMode: false,

		prefix: "testme",
		imageName: "testImage",
		imageTag: "1.2.3",

		hostname: "kube.test",
		externalPath: "hello-world",

		replicas: 42,
	}
	ops.execute()
	ops.createMode = false
	ops.deleteMode = true
	execResult := ops.execute()

	assert.True(t, execResult)

	deployments := listDeployments(testK8s)
	assert.Len(t, deployments.Items, 0)
	services := listServices(testK8s)
	assert.Len(t, services.Items, 0)
	ingresses := listIngress(testK8s)
	assert.Len(t, ingresses.Items, 0)
}

func TestOperationCanOperateIndependentlyOnTwoDifferentPrefixes(t *testing.T) {

	testK8s := testK8sSystem()

	ops1 := &operations{
		sys: testK8s,
		createMode: true,
		listMode: false,
		deleteMode: false,
		updateMode: false,

		prefix: "testme1",
		imageName: "testImage",
		imageTag: "1.2.3",

		hostname: "kube.test1",
		externalPath: "hello-one",

		replicas: 42,
	}
	execResult1 := ops1.execute()
	assert.True(t, execResult1)

	ops2 := &operations{
		sys: testK8s,
		createMode: true,
		listMode: false,
		deleteMode: false,
		updateMode: false,

		prefix: "testme2",
		imageName: "testImage",
		imageTag: "1.2.3",

		hostname: "kube.test2",
		externalPath: "hello-two",

		replicas: 999,
	}
	execResult2 := ops2.execute()
	assert.True(t, execResult2)

	deployments := listDeployments(testK8s)
	assert.Len(t, deployments.Items, 2)
	assert.EqualValues(t, "testme1-deployment", deployments.Items[0].Name)
	assert.EqualValues(t, int32(42), *deployments.Items[0].Spec.Replicas)
	assert.EqualValues(t, "testme2-deployment", deployments.Items[1].Name)
	assert.EqualValues(t, int32(999), *deployments.Items[1].Spec.Replicas)

	ops2.replicas = 3
	ops2.createMode = false
	ops2.updateMode = true
	ops2.execute()

	deployments = listDeployments(testK8s)
	assert.Len(t, deployments.Items, 2)
	assert.EqualValues(t, "testme1-deployment", deployments.Items[0].Name)
	assert.EqualValues(t, int32(42), *deployments.Items[0].Spec.Replicas)
	assert.EqualValues(t, "testme2-deployment", deployments.Items[1].Name)
	assert.EqualValues(t, int32(3), *deployments.Items[1].Spec.Replicas)

	ops1.replicas = 3
	ops1.createMode = false
	ops1.deleteMode = true
	ops1.execute()

	deployments = listDeployments(testK8s)
	assert.Len(t, deployments.Items, 1)
	assert.EqualValues(t, "testme2-deployment", deployments.Items[0].Name)
	assert.EqualValues(t, int32(3), *deployments.Items[0].Spec.Replicas)

}




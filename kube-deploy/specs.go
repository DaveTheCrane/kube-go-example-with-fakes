package main

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/* build a deployment, i.e. a persistent set of pods */
func deploymentSpec(deploymentName string, imageName string, imageTag string, 
	containerName string, containerPort int32, 
	labels map[string]string) *appsv1.Deployment {
	
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  containerName,
							Image: fmt.Sprintf("%s:%s", imageName, imageTag),
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: containerPort,
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment
}

/* wrap a deployment in a service, exposing it to outside world */
func serviceSpec(serviceName string, servicePort int32,
	labels map[string]string) *apiv1.Service {

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: apiv1.ServiceSpec{
			Type:     apiv1.ServiceTypeNodePort,
			Selector: labels,
			Ports: []apiv1.ServicePort{
				{
					Port: servicePort,
				},
			},
		},
	}

	return service
}

func int32Ptr(i int32) *int32 { return &i }
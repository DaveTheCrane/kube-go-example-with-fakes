package main

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	apiv1b1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

/* build a deployment, i.e. a persistent set of pods */
func deploymentSpec(deploymentName string, imageName string, imageTag string, 
	containerName string, containerPort int,
	numReplicas int32,
	labels map[string]string) *appsv1.Deployment {
	
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(numReplicas),
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
									ContainerPort: int32(containerPort),
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

/* wrap a deployment in a service (of type ClusterIP, i.e. without exposing it to outside world) */
func serviceSpec(serviceName string, servicePort int,
	labels map[string]string) *apiv1.Service {

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: apiv1.ServiceSpec{
			Type:     apiv1.ServiceTypeClusterIP,
			Selector: labels,
			Ports: []apiv1.ServicePort{
				{
					Port: int32(servicePort),
				},
			},
		},
	}

	return service
}

/* spec for an Ingress service to load balance a service rather than the user access it directly */
func ingressLoadBalancerSpec(lbName string, host string, path string, rewrittenPath string, serviceName string, servicePort int) *apiv1b1.Ingress {
	ingress := &apiv1b1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: lbName,
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/use-regex": "true",
				"nginx.ingress.kubernetes.io/rewrite-target": rewrittenPath,
			},
		},
		Spec: apiv1b1.IngressSpec{
			Rules:[]apiv1b1.IngressRule{
				{
					Host: host,
					IngressRuleValue: apiv1b1.IngressRuleValue{
						HTTP: &apiv1b1.HTTPIngressRuleValue{
							Paths: []apiv1b1.HTTPIngressPath{
								{
									Path: path,
									Backend:apiv1b1.IngressBackend{
										ServiceName: serviceName,
										ServicePort: intOrStr(servicePort),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return ingress
}

func int32Ptr(i int32) *int32 { return &i }
func intOrStr(i int) intstr.IntOrString { return intstr.FromInt(i) }

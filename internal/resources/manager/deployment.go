package manager

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ZupIT/horusec-operator/api/v2alpha1"
)

//nolint:funlen // to improve in the future
func NewDeployment(resource *v2alpha1.HorusecPlatform) appsv1.Deployment {
	return appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resource.GetManagerName(),
			Namespace: resource.GetNamespace(),
			Labels:    resource.GetManagerLabels(),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: resource.GetManagerReplicaCount(),
			Selector: &metav1.LabelSelector{MatchLabels: resource.GetManagerLabels()},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: resource.GetManagerLabels()},
				Spec: corev1.PodSpec{Containers: []corev1.Container{{
					Name:  resource.GetManagerName(),
					Image: resource.GetManagerImage(),
					Env: []corev1.EnvVar{
						{Name: "REACT_APP_HORUSEC_ENDPOINT_API", Value: resource.GetAPIEndpoint()},
						{Name: "REACT_APP_HORUSEC_ENDPOINT_ANALYTIC", Value: resource.GetAnalyticEndpoint()},
						{Name: "REACT_APP_HORUSEC_ENDPOINT_CORE", Value: resource.GetCoreEndpoint()},
						{Name: "REACT_APP_HORUSEC_ENDPOINT_WEBHOOK", Value: resource.GetWebhookEndpoint()},
						{Name: "REACT_APP_HORUSEC_ENDPOINT_AUTH", Value: resource.GetAuthEndpoint()},
						{Name: "REACT_APP_HORUSEC_MANAGER_PATH", Value: "\\/manager"},
						{Name: "REACT_APP_KEYCLOAK_BASE_PATH", Value: resource.Spec.Global.Keycloak.PublicURL},
						{Name: "REACT_APP_KEYCLOAK_CLIENT_ID", Value: resource.Spec.Global.Keycloak.Clients.Public.ID},
						{Name: "REACT_APP_KEYCLOAK_REALM", Value: resource.Spec.Global.Keycloak.Realm},
					},
					Ports: []corev1.ContainerPort{
						{Name: "http", ContainerPort: int32(resource.GetManagerPortHTTP())},
					},
				}}},
			},
		},
	}
}

func NewEnvFromSecret(variableName, secretName, secretKey string) corev1.EnvVar {
	return corev1.EnvVar{
		Name: variableName,
		ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{
			LocalObjectReference: corev1.LocalObjectReference{Name: secretName},
			Key:                  secretKey,
		}},
	}
}
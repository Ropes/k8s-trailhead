package kubernetes

import (
	"fmt"

	"k8s.io/api/core/v1"
	//apiv1b2 "k8s.io/api/extensions/v1beta1"
	apiv1b2 "k8s.io/api/apps/v1beta2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var deploymentName = "simpledeploy"

// SimpleDeployment creates an effectively static Kubernetes Deployment
// object with only the name and image being configured. No helper
// functions are used, all data structures are defined inline.
// This example is equivalent to a simple yaml file, but with type
// enforcement!
func SimpleDeployment(image, tag string) *apiv1b2.Deployment {
	name := deploymentName
	imgtag := fmt.Sprintf("%s:%s", image, tag)

	return &apiv1b2.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: map[string]string{"app": name, "tier": "api"},
		},
		Spec: apiv1b2.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": name, "tier": "api"},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Name: name,
					Labels: map[string]string{"app": name, "tier": "api"},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						v1.Container{
							Name:            name,
							Image:           imgtag,
							ImagePullPolicy: "Always",
						},
					},
				},
			},
		},
	}
}

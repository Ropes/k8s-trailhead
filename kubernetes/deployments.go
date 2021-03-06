package kubernetes

import (
	"fmt"

	"k8s.io/api/core/v1"
	//apiv1b2 "k8s.io/api/extensions/v1beta1"
	apiv1b2 "k8s.io/api/apps/v1beta2"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const helloKubeconDeploymentName = "hellokubecon"

// webappMeta defines ObjectMeta(data) for webapp objects...
func webappMeta(name, app string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:   name,
		Labels: map[string]string{"app": app, "tier": "webapp"},
	}
}

// cpuQuantity accepts a float64 of full CPU cores and returns
// a millicore *resource.Quantity.(rounds up at 0.5)
func cpuQuantity(cpu float64) resource.Quantity {
	f := cpu * float64(1000.0)
	v := round(f)
	q := fmt.Sprintf("%dm", int(v))
	return resource.MustParse(q)
}

// kubeconContainer returns a named v1.Container object with resources
// limited to 0.5 CPU and scheduled with 0.2CPU.
func kubeconContainer(name, image, tag string) v1.Container {
	return v1.Container{
		Name:            name,
		Image:           fmt.Sprintf("%s:%s", image, tag),
		ImagePullPolicy: "Always",
		Resources: v1.ResourceRequirements{
			Limits:   v1.ResourceList{"cpu": cpuQuantity(0.5)},
			Requests: v1.ResourceList{"cpu": cpuQuantity(0.2)},
		},
	}
}

// kubeconDeployment creates the Deployment Object structure or returns
// error if parameters are unspecified.
func kubeconDeployment(image, tag string) (*apiv1b2.Deployment, error) {
	om := webappMeta(helloKubeconDeploymentName, "api")
	if tag == "" || image == "" {
		return nil, fmt.Errorf("error: image and tag must be defined")
	}
	pts := &v1.PodTemplateSpec{
		ObjectMeta: om,
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				kubeconContainer(helloKubeconDeploymentName, image, tag),
			},
		},
	}
	d := &apiv1b2.Deployment{
		ObjectMeta: om,
		Spec: apiv1b2.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: om.Labels,
			},
			Template: *pts,
		},
	}
	return d, nil
}

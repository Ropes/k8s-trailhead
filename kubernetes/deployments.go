package kubernetes

import (
	"fmt"

	apiextv2 "k8s.io/api/autoscaling/v2beta1"
	"k8s.io/api/core/v1"
	apiextv1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

func kubeconContainers(name, image string) v1.Container { return v1.Container{} }

func kubeconDeployment(image, tag string, replicas int) (*apiextv1.Deployment, error) {
	name := "hellokubecon"
	om := webappMeta(name, "api")

	_ = apiextv2.HorizontalPodAutoscaler{}

	if tag == "" || image == "" {
		return nil, fmt.Errorf("error: image and tag must be defined")
	}

	pts := &v1.PodTemplateSpec{
		ObjectMeta: om,
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				v1.Container{
					Name:            name,
					Image:           fmt.Sprintf("%s:%s", image, tag),
					ImagePullPolicy: "Always",
				},
			},
		},
	}

	r := int32(replicas)
	d := &apiextv1.Deployment{
		ObjectMeta: om,
		Spec: apiextv1.DeploymentSpec{
			Template: *pts,
			Replicas: &r,
		},
	}

	return d, nil
}

func kubeconAutoscaller(image, tag, cpuLimit int) (*apiextv2.HorizontalPodAutoscaler, error) {
	return nil, nil
}

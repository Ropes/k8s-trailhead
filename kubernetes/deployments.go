package kubernetes

import (
	"fmt"

	"k8s.io/api/core/v1"
	apiextv1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func kubeconMeta(name string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:   name,
		Labels: map[string]string{"app": name, "tier": "webapp"},
	}
}

func HelloKubecon(image, tag string, replicas int) (*apiextv1.Deployment, error) {
	name := "hellokubecon"
	om := kubeconMeta(name)

	if tag == "" || image == "" {
		return nil, fmt.Errorf("error: tag undefined")
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

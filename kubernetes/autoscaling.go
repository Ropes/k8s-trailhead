package kubernetes

import apiextv2 "k8s.io/api/autoscaling/v2beta1"

var avgCpuUtil = int32(50)

func kubeconAutoscaller(image, tag string, cpuLimit, cpuavg float64) (*apiextv2.HorizontalPodAutoscaler, error) {
	om := webappMeta(helloKubeconDeploymentName, "api")
	minr, maxr := int32(1), int32(10)

	hpa := &apiextv2.HorizontalPodAutoscaler{
		ObjectMeta: om,
		Spec: apiextv2.HorizontalPodAutoscalerSpec{
			MinReplicas: &minr,
			MaxReplicas: maxr,
			ScaleTargetRef: apiextv2.CrossVersionObjectReference{
				Kind:       "Deployment",
				APIVersion: "apps/v1beta1",
				Name:       helloKubeconDeploymentName,
			},
			Metrics: []apiextv2.MetricSpec{
				apiextv2.MetricSpec{
					Type: "Resource",
					Resource: &apiextv2.ResourceMetricSource{
						Name: "cpu",
						TargetAverageUtilization: &avgCpuUtil,
					},
				},
			},
		},
	}
	return hpa, nil
}

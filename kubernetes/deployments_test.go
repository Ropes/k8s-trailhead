package kubernetes

import (
	"fmt"
	"testing"
)

func TestHelloKubecon(t *testing.T) {
	tests := []struct {
		namespace string
		image     string
		tag       string
		replicas  int
		err       bool
	}{
		{
			namespace: "default",
			image:     "kubecon",
			tag:       "",
			replicas:  3,
			err:       true,
		},
		{
			namespace: "default",
			image:     "kubecon",
			tag:       "0.1",
			replicas:  3,
			err:       false,
		},
		{
			namespace: "kubecon",
			image:     "",
			tag:       "0.1",
			replicas:  3,
			err:       true,
		},
		{
			namespace: "kubecon",
			image:     "hihi",
			tag:       "0.1",
			replicas:  3,
			err:       false,
		},
	}

	for i, T := range tests {
		t.Run(fmt.Sprintf("hello-deployment-test-%d", i), func(t *testing.T) {

			d, err := HelloKubecon(T.image, T.tag, T.replicas)
			t.Logf("%#v", T)
			if T.err && err == nil {
				t.Errorf("an error was expected for: %#v", T)
				return
			}
			if T.err && d != nil {
				t.Errorf("Deployment expected nil")
			}
			if T.err == false {
				exp := int32(T.replicas)
				reps := d.Spec.Replicas
				if !T.err && err != nil && *reps != exp {
					t.Errorf("replica value unexpected: %#v", d.Spec)
				}
			}
		})
	}
}

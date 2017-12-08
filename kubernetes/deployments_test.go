package kubernetes

import (
	"fmt"
	"os"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
			err:       true,
		},
		{
			namespace: "default",
			image:     "kubecon",
			tag:       "0.1",
			err:       false,
		},
		{
			namespace: "kubecon",
			image:     "",
			tag:       "0.1",
			err:       true,
		},
		{
			namespace: "kubecon",
			image:     "hihi",
			tag:       "0.1",
			err:       false,
		},
	}

	for i, T := range tests {
		t.Run(fmt.Sprintf("hello-deployment-test-%d", i), func(t *testing.T) {

			d, err := kubeconDeployment(T.image, T.tag)
			t.Logf("%#v", T)
			if T.err && err == nil {
				t.Errorf("an error was expected for: %#v", T)
				return
			}
			if T.err && d != nil {
				t.Errorf("Deployment expected nil")
			}
			if T.err == false {
			}
		})
	}
}

func TestQuantity(t *testing.T) {
	tests := []struct {
		v float64
		s string
	}{
		{
			v: float64(2.5),
			s: "2500m",
		},
		{
			v: float64(2.0),
			s: "2",
		},
		{
			v: float64(0.5),
			s: "500m",
		},
	}

	for i, T := range tests {
		t.Run(fmt.Sprintf("test-quantity-%d", i), func(t *testing.T) {
			q := cpuQuantity(T.v)
			qs := q.String()
			if qs != T.s {
				t.Errorf("quantity expected: %q, got %q", T.s, qs)
			}

		})
	}
}

func TestDeploymentMinikubeE2E(t *testing.T) {
	// TestMini
	if _, ok := os.LookupEnv("TESTMINIKUBE"); !ok {
		t.Skip("TESTMINIKUBE unset")
	}
	// Create Deployment
	d, err := kubeconDeployment("localhost", "5000")
	if err != nil || d == nil {
		t.Error(err)
		t.Logf("%#v", d)
	}

	kcp := os.Getenv("HOME") + "/.kube/config"
	// TODO: Create client to to minikube
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kcp},
		&clientcmd.ConfigOverrides{CurrentContext: "minikube"},
	).ClientConfig()
	if err != nil {
		t.Fatal("error reading config %#v: %v", config, err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Fatal("error creating client: %v", err)
	}

	// Create Minikube Deployment
	d, err = clientset.Extensions().Deployments("default").Create(d)
	if err != nil {
		t.Error(err)
		t.Fatal("error creating deployment")
	}
	time.Sleep(1 * time.Second) // give some time for the API

	// Verify Deployment exists
	ds, err := clientset.Extensions().Deployments("default").List(metav1.ListOptions{LabelSelector: "app=api"})
	if err != nil {
		t.Error(err)
	}
	if len(ds.Items) < 1 {
		t.Errorf("error: deployments should not be 0")
	}
	t.Logf("%#v", ds.Items[0])
	time.Sleep(1 * time.Second) // give some time for the API

	/*
		// Assert services
		svc, err := clientset.Core().Services("default").List(metav1.ListOptions{})
		if err != nil {
			t.Errorf("svc get err: %v", err)
		}
		if len(svc.Items) < 1 {
			t.Errorf("at least one service expected")
		}
	*/

	// Delete Deployment
	err = clientset.Extensions().Deployments("default").Delete(helloKubeconDeploymentName, nil)
	if err != nil {
		t.Errorf("error deleting deployment: %v", err)
	}
}

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

var (
	kubeconImg = "localhost:5000/kubecon"
	kubeconTag = "1.0"
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

func TestDeploymentCreateUpdateNoCleanup(t *testing.T) {
	if _, ok := os.LookupEnv("TESTMINIKUBE"); !ok {
		t.Skip("TESTMINIKUBE unset")
	}
	// Create Deployment
	d, err := kubeconDeployment(kubeconImg, kubeconTag)
	if err != nil || d == nil {
		t.Error(err)
		t.Logf("%#v", d)
	}

	kcp := os.Getenv("HOME") + "/.kube/config"
	// Create client to to minikube
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

	// Check if deployment already exists
	exists := false
	ds, err := clientset.AppsV1beta2().Deployments("default").List(metav1.ListOptions{LabelSelector: "app=api"})
	if err != nil {
		t.Error(err)
	}
	if len(ds.Items) != 1 {
		t.Errorf("error: deployments should be 1")
	} else {
		chkd := ds.Items[0]
		if chkd.Name != helloKubeconDeploymentName {
			t.Errorf("unexpected Deployment: %v", chkd.Name)
		}
	}

	// Create/Update Minikube Deployment
	if exists {
		d, err = clientset.AppsV1beta2().Deployments("default").Update(d)
		if err != nil {
			t.Error(err)
			t.Fatal("error creating deployment")
		}
	} else {
		d, err = clientset.AppsV1beta2().Deployments("default").Create(d)
		if err != nil {
			t.Error(err)
			t.Fatal("error creating deployment")
		}
	}
	t.Logf("Deployment \n%#v", d)
}

func testClientset(t *testing.T) *kubernetes.Clientset {
	// Test Minikube  Deployment
	if _, ok := os.LookupEnv("TESTMINIKUBE"); !ok {
		t.Skip("TESTMINIKUBE unset")
	}
	kcp := os.Getenv("HOME") + "/.kube/config"
	// Create client to to minikube
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
	t.Logf("clientset created")
	return clientset
}

func TestDeploymentMinikubeE2E(t *testing.T) {
	// Test Minikube  Deployment
	if _, ok := os.LookupEnv("TESTMINIKUBE"); !ok {
		t.Skip("TESTMINIKUBE unset")
	}
	// Create Deployment
	d, err := kubeconDeployment(kubeconImg, kubeconTag)
	if err != nil || d == nil {
		t.Error(err)
		t.Logf("%#v", d)
	}
	t.Logf("deployment struct created")

	kcp := os.Getenv("HOME") + "/.kube/config"
	// Create client to to minikube
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
	t.Logf("clientset created")

	// Create Minikube Deployment
	d, err = clientset.AppsV1beta2().Deployments("default").Create(d)
	if err != nil {
		t.Error(err)
		t.Fatal("error creating deployment")
	}
	time.Sleep(1 * time.Second) // give some time for the API
	t.Logf("deployment created by client")

	// Verify Deployment exists
	ds, err := clientset.AppsV1beta2().Deployments("default").List(metav1.ListOptions{LabelSelector: "app=api"})
	if err != nil {
		t.Error(err)
	}
	if len(ds.Items) < 1 {
		t.Fatal("error: deployments should not be 0")
	}
	time.Sleep(1 * time.Second) // give some time for the API
	t.Logf("deployment exists!")

	// Delete Deployment
	err = clientset.AppsV1beta2().Deployments("default").Delete(helloKubeconDeploymentName, nil)
	if err != nil {
		t.Errorf("error deleting deployment: %v", err)
	}
	t.Logf("deployment deleted")
}

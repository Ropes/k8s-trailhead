package operators

import (
	"context"
	"os"
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func TestSimpleObserve(t *testing.T) {
	kc := testClientset(t)
	ctx, can := context.WithCancel(context.Background())
	defer can()

	op := NewSimpleOperator(ctx, kc)
	op.Observe()

	go func() {
		for {
			err := <-op.errors
			t.Error(err)
		}
	}()

	if op.observed == nil {
		t.Errorf("observed config is nil")
	} else {
		t.Logf("%#v", op.observed)
	}
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
		t.Fatalf("error reading config %#v: %v", config, err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	t.Logf("clientset created")
	return clientset
}

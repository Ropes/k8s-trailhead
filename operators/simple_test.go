package operators

import (
	"context"
	"os"
	"strconv"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	testNamespace     = "default"
	testConfigMapName = "test-configmap"
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

func TestConfigMapCreate(t *testing.T) {
	kc := testClientset(t)
	testConfigmap(t, kc)
}

// testconfigmap deletes the existing ConfigMap and creates a new one
// containing the SimpleSpec fields
func testConfigmap(t *testing.T, kc kubernetes.Interface) {
	sc := SimpleSpec{
		Replicas: 2,
		Image:    "nginx:1.6",
	}
	// iff already exists: delete
	gcm := configmapGet(t, kc)
	if gcm != nil {
		t.Logf("configmap already exists")
		configmapDelete(t, kc)
	}

	testData := map[string]string{
		"replicas": strconv.Itoa(sc.Replicas),
		"image":    sc.Image,
	}
	testCm := v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: testConfigMapName,
			Labels: map[string]string{
				"tier": "testing",
			},
		},
		Data: testData,
	}
	t.Logf("no existing configmap, creating data for test")
	_, err := kc.CoreV1().ConfigMaps(testNamespace).Create(&testCm)
	if err != nil {
		t.Fatal(err)
	}
}

func configmapGet(t *testing.T, kc kubernetes.Interface) *v1.ConfigMap {
	gcm, err := kc.CoreV1().ConfigMaps(testNamespace).Get(
		testConfigMapName,
		metav1.GetOptions{
			TypeMeta: metav1.TypeMeta{Kind: "ConfigMap"},
		},
	)
	if err != nil {
		t.Logf("getting configmap err: %v", err)
		return nil
	}
	return gcm
}

func configmapDelete(t *testing.T, kc kubernetes.Interface) {
	err := kc.Core().ConfigMaps(testNamespace).Delete(
		testConfigMapName,
		&metav1.DeleteOptions{
			TypeMeta: metav1.TypeMeta{Kind: "ConfigMap"},
		},
	)
	if err != nil {
		t.Errorf("error deleting configmap: %v", err)
	}
}

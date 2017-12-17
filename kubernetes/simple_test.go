package kubernetes

import (
	"fmt"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSimpleDeploy(t *testing.T) {
	img, tag := "nginx", "1.13.7"

	d := SimpleDeployment(img, tag)

	ei := fmt.Sprintf("%s:%s", img, tag)
	if d.Spec.Template.Spec.Containers[0].Image != ei {
		t.Errorf("unexpected Image field value: %q", d.Spec.Template.Spec.Containers[0].Image)
	}
}

func TestSimpleDeployE2E(t *testing.T) {
	clientset := testClientset(t)

	img, tag := "nginx", "1.13.7"
	d := SimpleDeployment(img, tag)

	// Check Deployment exists
	ds, err := clientset.AppsV1beta2().Deployments("default").List(metav1.ListOptions{LabelSelector: fmt.Sprintf("app=%s", deploymentName)})
	if err != nil {
		t.Error(err)
	}
	if len(ds.Items) > 0 {
		t.Logf("nominal: deployment already exists, deleting")
		// Delete Deployment
		err = clientset.AppsV1beta2().Deployments("default").Delete(deploymentName, nil)
		if err != nil {
			t.Errorf("error deleting deployment: %v", err)
		}
		t.Logf("deployment deleted")
		time.Sleep(3 * time.Second) // give some time for the API
	}

	// Create Minikube Deployment
	d, err = clientset.AppsV1beta2().Deployments("default").Create(d)
	if err != nil {
		t.Errorf("creating deployment error: %s", err)
		t.Fatal("error creating deployment")
	}
	time.Sleep(1 * time.Second) // give some time for the API
	t.Logf("deployment created by client")

	// Verify Deployment exists
	time.Sleep(3 * time.Second) // give some time for the API
	ds, err = clientset.AppsV1beta2().Deployments("default").List(metav1.ListOptions{LabelSelector: fmt.Sprintf("app=%s", deploymentName)})
	if err != nil {
		t.Error(err)
	}
	if len(ds.Items) < 1 {
		t.Fatal("error: deployments should not be 0")
	}
	t.Logf("deployment exists!")

	// Delete Deployment
	err = clientset.AppsV1beta2().Deployments("default").Delete(deploymentName, nil)
	if err != nil {
		t.Errorf("error deleting deployment: %v", err)
	}
	t.Logf("deployment deleted")
}

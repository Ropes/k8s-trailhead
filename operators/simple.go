package operators

import (
	"context"
	"strconv"
	"time"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SimpleCmpArtifact SimpleConfig

// SimpleOperator implements the Operator interface for a Kubernetes Namespace.
type SimpleOperator struct {
	ctx       context.Context
	kclient   kubernetes.Interface
	errors    chan error
	namespace string

	observed *SimpleCmpArtifact
	expected *SimpleCmpArtifact
}

// NewSimpleOperator creates an Operator with context controlled exit.
func NewSimpleOperator(ctx context.Context, kclient kubernetes.Interface) *SimpleOperator {
	return &SimpleOperator{
		ctx:       ctx,
		kclient:   kclient,
		namespace: "default",
	}
}

// Run executes all Operator steps and
func (o *SimpleOperator) Run() {
	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-o.ctx.Done():
			return
		case <-tick.C:
			o.Observe()
			o.Analyze()
			o.Act()
		}
	}

}

func (o *SimpleOperator) Observe() {
	select {
	case <-o.ctx.Done():
		return
	default:
		//TODO: Query ConfigMap
		cm, err := o.kclient.Core().ConfigMaps(o.namespace).Get("simpleoperator", meta_v1.GetOptions{})
		if err != nil {
			o.errors <- err
			return
		}
		d := cm.Data

		//TODO: Query k8s api

		//TODO: Set state
		sz, err := strconv.Atoi(d["size"])
		if err != nil {
			o.errors <- err
			return
		}
		s := &SimpleConfig{
			Size: sz,
		}
	}
}

func (o *SimpleOperator) Analyze() {
	select {
	case <-o.ctx.Done():
		return
	default:
	}
}

// Act executes changes.
func (o *SimpleOperator) Act() {
	select {
	case <-o.ctx.Done():
		return
	default:
	}
}

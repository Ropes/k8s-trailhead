package operators

// Operator defines the pattern for managing complex services in Kubernetes.
type Operator interface {
	//Run()
	Observe()
	Analyze()
	Act()
}

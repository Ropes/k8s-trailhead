package operators

// SimpleConfigMapName identifies k8s API ConfigMap  name for a Namespace.
const SimpleConfigMapName = "simple"

// SimpleConfig is the data written to the
type SimpleConfig struct {
	Replicas int    `json:"size"`
	Image    string `json:"image"`
}

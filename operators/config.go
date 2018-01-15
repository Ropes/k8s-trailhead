package operators

// SimpleConfigMapName identifies k8s API ConfigMap  name for a Namespace.
const SimpleConfigMapName = "simple"

// SimpleSpec is the data written to the
type SimpleSpec struct {
	Replicas int    `json:"size"`
	Image    string `json:"image"`
}

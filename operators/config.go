package operators

// SimpleConfigMapName identifies k8s API ConfigMap  name for a Namespace.
const SimpleConfigMapName = "simple"

// SimpleConfig is the data written to the
type SimpleConfig struct {
	Size int `json:"size"`
}

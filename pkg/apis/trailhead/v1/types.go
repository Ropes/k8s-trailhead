package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SimpleSpec defines a k8s API Custome Defined Resource for specifying
// image, tag, and replica count.
type SimpleSpec struct {
	metav1.TypeMeta   `json:"typemeta,inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Replicas int    `json:"size,omitempty"`
	Image    string `json:"image,omitempty"`
	Tag      string `json:"tag,omitempty"`
}

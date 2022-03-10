package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



type EtcdRestorPhase string
var (
	EetcdRestorRunning   EtcdRestorPhase = "Running"
	EtcdRestorComplated  EtcdRestorPhase = "Complated"
	EtcdRestorFailed  EtcdRestorPhase = "Failed"
)
type EtcdRestorCondition struct{
	Ready  bool   `json:"ready"`
	Reason  string  `json:"reason,omitempty"`
	Message  string  `json:"message,omitempty"`
	LastedTranslationTime  metav1.Time  `json:"lastedTranslationTime"`
	Location   string    `json:"location,omitempty"`
}
// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EtcdRestoreSpec defines the desired state of EtcdRestore
type EtcdRestoreSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	ClusterRefernece  string `json:"clusterReference"`
	DataURL   string  `json:"dataURL"`
}

// EtcdRestoreStatus defines the observed state of EtcdRestore
type EtcdRestoreStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Conditions []EtcdRestorCondition `json:"conditions"`
	Phase   EtcdRestorPhase  `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EtcdRestore is the Schema for the etcdrestores API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=etcdrestores,scope=Namespaced
type EtcdRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdRestoreSpec   `json:"spec,omitempty"`
	Status EtcdRestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EtcdRestoreList contains a list of EtcdRestore
type EtcdRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EtcdRestore{}, &EtcdRestoreList{})
}

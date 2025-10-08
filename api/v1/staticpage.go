package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
type StaticPage struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec StaticPageSpec `json:"spec,omitempty"`
}

type StaticPageSpec struct {
    Contents string `json:"contents,omitempty"`
    Image    string `json:"image,omitempty"`
    Replicas int32  `json:"replicas,omitempty"`
}

// +kubebuilder:object:root=true
type StaticPageList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`

    Items []StaticPage `json:"items"`
}
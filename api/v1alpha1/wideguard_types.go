/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WideguardSpec defines the desired state of Wideguard
type WideguardSpec struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:ExclusiveMaximum=false
	// +kubebuilder:default=1
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Replicas int32 `json:"replicas,omitempty"`

	// +kubebuilder:default=51820
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ContainerPort int32 `json:"containerPort,omitempty"`

	// +kubebuilder:default="192.168.1.1/24"
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Network string `json:"network,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ExternalDNS ExternalDNS `json:"externalDns,omitempty"`
}

type ExternalDNS struct {
	// +kubebuilder:default=true
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:default="docker.io/klutchell/unbound:v1.17.1"
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Image string `json:"image,omitempty"`
}

// WideguardStatus defines the observed state of Wideguard
type WideguardStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Wideguard is the Schema for the wideguards API
type Wideguard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WideguardSpec   `json:"spec,omitempty"`
	Status WideguardStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WideguardList contains a list of Wideguard
type WideguardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wideguard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Wideguard{}, &WideguardList{})
}

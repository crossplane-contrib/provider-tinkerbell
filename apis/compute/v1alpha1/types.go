/*
Copyright 2020 The Crossplane Authors.

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

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// HardwareParameters are the configurable fields of a Hardware.
type HardwareParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// HardwareObservation are the observable fields of a Hardware.
type HardwareObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A HardwareSpec defines the desired state of a Hardware.
type HardwareSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  HardwareParameters `json:"forProvider"`
}

// A HardwareStatus represents the observed state of a Hardware.
type HardwareStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     HardwareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Hardware is an example API type
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type Hardware struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HardwareSpec   `json:"spec"`
	Status HardwareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HardwareList contains a list of Hardware
type HardwareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hardware `json:"items"`
}

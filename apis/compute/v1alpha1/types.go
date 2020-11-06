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

// TemplateParameters are the configurable fields of a Template.
type TemplateParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// TemplateObservation are the observable fields of a Template.
type TemplateObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A TemplateSpec defines the desired state of a Template.
type TemplateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  TemplateParameters `json:"forProvider"`
}

// A TemplateStatus represents the observed state of a Template.
type TemplateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     TemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Template is an example API type
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TemplateSpec   `json:"spec"`
	Status TemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TemplateList contains a list of Template
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Template `json:"items"`
}

// WorkflowParameters are the configurable fields of a Workflow.
type WorkflowParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// WorkflowObservation are the observable fields of a Workflow.
type WorkflowObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A WorkflowSpec defines the desired state of a Workflow.
type WorkflowSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WorkflowParameters `json:"forProvider"`
}

// A WorkflowStatus represents the observed state of a Workflow.
type WorkflowStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WorkflowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Workflow is an example API type
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkflowSpec   `json:"spec"`
	Status WorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkflowList contains a list of Workflow
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items"`
}

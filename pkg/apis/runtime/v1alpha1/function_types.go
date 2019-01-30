/*
Copyright 2019 The Kyma Authors.

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
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	// function defines the content of a function
	Function string `json:"function"`

	// functionContentType defines file content type (plaintext or base64)
	FunctionContentType string `json:"functionContentType"`

	// size defines as the size of a function pertaining to memory and cpu only. Values can be any one of these S, M, L, XL
	Size string `json:"size"`

	// runtime is the programming language used for a function e.g. nodejs8
	Runtime string `json:"runtime"`

	// timeout defines maximum duration alloted to a function to complete its execution, defaults to 180s
	Timeout int32 `json:"timeout,omitempty"`

	// deps defines the dependencies for a function
	Deps string `json:"deps,omitempty"`

	// envs defines an array of key value pairs need to be used as env variable for a function
	Env []v1.EnvVar `json:"env,omitempty"`
}

// FunctionStatus defines the observed state of Function
type FunctionStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Function is the Schema for the functions API
// +k8s:openapi-gen=true
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FunctionSpec   `json:"spec,omitempty"`
	Status FunctionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FunctionList contains a list of Function
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}

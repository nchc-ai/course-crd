/*
Copyright 2017 The Kubernetes Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Course is a specification for a Course resource
type Course struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CourseSpec   `json:"spec"`
	Status CourseStatus `json:"status"`
}

// CourseSpec is the spec for a Course resource
type CourseSpec struct {
	Schedule       []string         `json:"schedule"`
	WritableVolume *WritableVolume  `json:"writableVolume,omitempty"`
	Port           map[string]int32 `json:"port"`
	Gpu            int32            `json:"gpu"`
	Image          string           `json:"image"`
	Dataset        []string         `json:"dataset,omitempty"`
	AccessType     AccessType       `json:"accessType"`
}

// CourseStatus is the status for a Course resource
type CourseStatus struct {
	Accessible  bool              `json:"accessible"`
	ServiceName string            `json:"service"`
	NodePort    map[string]int32  `json:"nodePort,omitempty"`
	SubPath     map[string]string `json:"subPath,omitempty"`
}

type WritableVolume struct {
	Owner        string `json:"owner"`
	StorageClass string `json:"storageclass"`
	Uid          int64  `json:"uid"`
	MountPoint   string `json:"mountPoint"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CourseList is a list of Course resources
type CourseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Course `json:"items"`
}

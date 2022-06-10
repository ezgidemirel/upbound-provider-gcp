/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha2 "github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1alpha2"
	v1alpha2compute "github.com/upbound/official-providers/provider-gcp/apis/compute/v1alpha2"
	v1alpha2container "github.com/upbound/official-providers/provider-gcp/apis/container/v1alpha2"
	v1alpha2monitoring "github.com/upbound/official-providers/provider-gcp/apis/monitoring/v1alpha2"
	v1alpha2redis "github.com/upbound/official-providers/provider-gcp/apis/redis/v1alpha2"
	v1alpha2sql "github.com/upbound/official-providers/provider-gcp/apis/sql/v1alpha2"
	v1alpha2storage "github.com/upbound/official-providers/provider-gcp/apis/storage/v1alpha2"
	v1alpha1 "github.com/upbound/official-providers/provider-gcp/apis/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha2.SchemeBuilder.AddToScheme,
		v1alpha2compute.SchemeBuilder.AddToScheme,
		v1alpha2container.SchemeBuilder.AddToScheme,
		v1alpha2monitoring.SchemeBuilder.AddToScheme,
		v1alpha2redis.SchemeBuilder.AddToScheme,
		v1alpha2sql.SchemeBuilder.AddToScheme,
		v1alpha2storage.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
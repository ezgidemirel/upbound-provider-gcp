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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TenantDefaultSupportedIdPConfigObservation struct {

	// an identifier for the resource with format projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the default supported IDP config resource
	// The name of the default supported IDP config resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TenantDefaultSupportedIdPConfigParameters struct {

	// OAuth client ID
	// OAuth client ID
	// +kubebuilder:validation:Required
	ClientIDSecretRef v1.SecretKeySelector `json:"clientIdSecretRef" tf:"-"`

	// OAuth client secret
	// OAuth client secret
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// If this IDP allows the user to sign in
	// If this IDP allows the user to sign in
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// ID of the IDP. Possible values include:
	// ID of the IDP. Possible values include:
	//
	// * 'apple.com'
	//
	// * 'facebook.com'
	//
	// * 'gc.apple.com'
	//
	// * 'github.com'
	//
	// * 'google.com'
	//
	// * 'linkedin.com'
	//
	// * 'microsoft.com'
	//
	// * 'playgames.google.com'
	//
	// * 'twitter.com'
	//
	// * 'yahoo.com'
	// +kubebuilder:validation:Required
	IdPID *string `json:"idpId" tf:"idp_id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the tenant where this DefaultSupportedIdpConfig resource exists
	// The name of the tenant where this DefaultSupportedIdpConfig resource exists
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/identityplatform/v1beta1.Tenant
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// +kubebuilder:validation:Optional
	TenantRef *v1.Reference `json:"tenantRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TenantSelector *v1.Selector `json:"tenantSelector,omitempty" tf:"-"`
}

// TenantDefaultSupportedIdPConfigSpec defines the desired state of TenantDefaultSupportedIdPConfig
type TenantDefaultSupportedIdPConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TenantDefaultSupportedIdPConfigParameters `json:"forProvider"`
}

// TenantDefaultSupportedIdPConfigStatus defines the observed state of TenantDefaultSupportedIdPConfig.
type TenantDefaultSupportedIdPConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TenantDefaultSupportedIdPConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TenantDefaultSupportedIdPConfig is the Schema for the TenantDefaultSupportedIdPConfigs API. Configurations options for the tenant for authenticating with a the standard set of Identity Toolkit-trusted IDPs.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TenantDefaultSupportedIdPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TenantDefaultSupportedIdPConfigSpec   `json:"spec"`
	Status            TenantDefaultSupportedIdPConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantDefaultSupportedIdPConfigList contains a list of TenantDefaultSupportedIdPConfigs
type TenantDefaultSupportedIdPConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TenantDefaultSupportedIdPConfig `json:"items"`
}

// Repository type metadata.
var (
	TenantDefaultSupportedIdPConfig_Kind             = "TenantDefaultSupportedIdPConfig"
	TenantDefaultSupportedIdPConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TenantDefaultSupportedIdPConfig_Kind}.String()
	TenantDefaultSupportedIdPConfig_KindAPIVersion   = TenantDefaultSupportedIdPConfig_Kind + "." + CRDGroupVersion.String()
	TenantDefaultSupportedIdPConfig_GroupVersionKind = CRDGroupVersion.WithKind(TenantDefaultSupportedIdPConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&TenantDefaultSupportedIdPConfig{}, &TenantDefaultSupportedIdPConfigList{})
}

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedIPRangeObservation) DeepCopyInto(out *AllowedIPRangeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedIPRangeObservation.
func (in *AllowedIPRangeObservation) DeepCopy() *AllowedIPRangeObservation {
	if in == nil {
		return nil
	}
	out := new(AllowedIPRangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedIPRangeParameters) DeepCopyInto(out *AllowedIPRangeParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedIPRangeParameters.
func (in *AllowedIPRangeParameters) DeepCopy() *AllowedIPRangeParameters {
	if in == nil {
		return nil
	}
	out := new(AllowedIPRangeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigObservation) DeepCopyInto(out *ConfigObservation) {
	*out = *in
	if in.AirflowURI != nil {
		in, out := &in.AirflowURI, &out.AirflowURI
		*out = new(string)
		**out = **in
	}
	if in.DagGcsPrefix != nil {
		in, out := &in.DagGcsPrefix, &out.DagGcsPrefix
		*out = new(string)
		**out = **in
	}
	if in.GkeCluster != nil {
		in, out := &in.GkeCluster, &out.GkeCluster
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigObservation.
func (in *ConfigObservation) DeepCopy() *ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigParameters) DeepCopyInto(out *ConfigParameters) {
	*out = *in
	if in.DatabaseConfig != nil {
		in, out := &in.DatabaseConfig, &out.DatabaseConfig
		*out = make([]DatabaseConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = make([]EncryptionConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvironmentSize != nil {
		in, out := &in.EnvironmentSize, &out.EnvironmentSize
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = make([]MaintenanceWindowParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = make([]NodeConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(float64)
		**out = **in
	}
	if in.PrivateEnvironmentConfig != nil {
		in, out := &in.PrivateEnvironmentConfig, &out.PrivateEnvironmentConfig
		*out = make([]PrivateEnvironmentConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SoftwareConfig != nil {
		in, out := &in.SoftwareConfig, &out.SoftwareConfig
		*out = make([]SoftwareConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebServerConfig != nil {
		in, out := &in.WebServerConfig, &out.WebServerConfig
		*out = make([]WebServerConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebServerNetworkAccessControl != nil {
		in, out := &in.WebServerNetworkAccessControl, &out.WebServerNetworkAccessControl
		*out = make([]WebServerNetworkAccessControlParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkloadsConfig != nil {
		in, out := &in.WorkloadsConfig, &out.WorkloadsConfig
		*out = make([]WorkloadsConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigParameters.
func (in *ConfigParameters) DeepCopy() *ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseConfigObservation) DeepCopyInto(out *DatabaseConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseConfigObservation.
func (in *DatabaseConfigObservation) DeepCopy() *DatabaseConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DatabaseConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseConfigParameters) DeepCopyInto(out *DatabaseConfigParameters) {
	*out = *in
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseConfigParameters.
func (in *DatabaseConfigParameters) DeepCopy() *DatabaseConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DatabaseConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigObservation) DeepCopyInto(out *EncryptionConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigObservation.
func (in *EncryptionConfigObservation) DeepCopy() *EncryptionConfigObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigParameters) DeepCopyInto(out *EncryptionConfigParameters) {
	*out = *in
	if in.KMSKeyName != nil {
		in, out := &in.KMSKeyName, &out.KMSKeyName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigParameters.
func (in *EncryptionConfigParameters) DeepCopy() *EncryptionConfigParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentObservation) DeepCopyInto(out *EnvironmentObservation) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentObservation.
func (in *EnvironmentObservation) DeepCopy() *EnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentParameters) DeepCopyInto(out *EnvironmentParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectSelector != nil {
		in, out := &in.ProjectSelector, &out.ProjectSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentParameters.
func (in *EnvironmentParameters) DeepCopy() *EnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAllocationPolicyObservation) DeepCopyInto(out *IPAllocationPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAllocationPolicyObservation.
func (in *IPAllocationPolicyObservation) DeepCopy() *IPAllocationPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(IPAllocationPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAllocationPolicyParameters) DeepCopyInto(out *IPAllocationPolicyParameters) {
	*out = *in
	if in.ClusterIPv4CidrBlock != nil {
		in, out := &in.ClusterIPv4CidrBlock, &out.ClusterIPv4CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.ClusterSecondaryRangeName != nil {
		in, out := &in.ClusterSecondaryRangeName, &out.ClusterSecondaryRangeName
		*out = new(string)
		**out = **in
	}
	if in.ServicesIPv4CidrBlock != nil {
		in, out := &in.ServicesIPv4CidrBlock, &out.ServicesIPv4CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.ServicesSecondaryRangeName != nil {
		in, out := &in.ServicesSecondaryRangeName, &out.ServicesSecondaryRangeName
		*out = new(string)
		**out = **in
	}
	if in.UseIPAliases != nil {
		in, out := &in.UseIPAliases, &out.UseIPAliases
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAllocationPolicyParameters.
func (in *IPAllocationPolicyParameters) DeepCopy() *IPAllocationPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(IPAllocationPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowObservation) DeepCopyInto(out *MaintenanceWindowObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowObservation.
func (in *MaintenanceWindowObservation) DeepCopy() *MaintenanceWindowObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowParameters) DeepCopyInto(out *MaintenanceWindowParameters) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowParameters.
func (in *MaintenanceWindowParameters) DeepCopy() *MaintenanceWindowParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigObservation) DeepCopyInto(out *NodeConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigObservation.
func (in *NodeConfigObservation) DeepCopy() *NodeConfigObservation {
	if in == nil {
		return nil
	}
	out := new(NodeConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigParameters) DeepCopyInto(out *NodeConfigParameters) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(float64)
		**out = **in
	}
	if in.IPAllocationPolicy != nil {
		in, out := &in.IPAllocationPolicy, &out.IPAllocationPolicy
		*out = make([]IPAllocationPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkSelector != nil {
		in, out := &in.NetworkSelector, &out.NetworkSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.OAuthScopes != nil {
		in, out := &in.OAuthScopes, &out.OAuthScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountSelector != nil {
		in, out := &in.ServiceAccountSelector, &out.ServiceAccountSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Subnetwork != nil {
		in, out := &in.Subnetwork, &out.Subnetwork
		*out = new(string)
		**out = **in
	}
	if in.SubnetworkRef != nil {
		in, out := &in.SubnetworkRef, &out.SubnetworkRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetworkSelector != nil {
		in, out := &in.SubnetworkSelector, &out.SubnetworkSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigParameters.
func (in *NodeConfigParameters) DeepCopy() *NodeConfigParameters {
	if in == nil {
		return nil
	}
	out := new(NodeConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEnvironmentConfigObservation) DeepCopyInto(out *PrivateEnvironmentConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEnvironmentConfigObservation.
func (in *PrivateEnvironmentConfigObservation) DeepCopy() *PrivateEnvironmentConfigObservation {
	if in == nil {
		return nil
	}
	out := new(PrivateEnvironmentConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEnvironmentConfigParameters) DeepCopyInto(out *PrivateEnvironmentConfigParameters) {
	*out = *in
	if in.CloudComposerConnectionSubnetwork != nil {
		in, out := &in.CloudComposerConnectionSubnetwork, &out.CloudComposerConnectionSubnetwork
		*out = new(string)
		**out = **in
	}
	if in.CloudComposerNetworkIPv4CidrBlock != nil {
		in, out := &in.CloudComposerNetworkIPv4CidrBlock, &out.CloudComposerNetworkIPv4CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.CloudSQLIPv4CidrBlock != nil {
		in, out := &in.CloudSQLIPv4CidrBlock, &out.CloudSQLIPv4CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.EnablePrivateEndpoint != nil {
		in, out := &in.EnablePrivateEndpoint, &out.EnablePrivateEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.MasterIPv4CidrBlock != nil {
		in, out := &in.MasterIPv4CidrBlock, &out.MasterIPv4CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.WebServerIPv4CidrBlock != nil {
		in, out := &in.WebServerIPv4CidrBlock, &out.WebServerIPv4CidrBlock
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEnvironmentConfigParameters.
func (in *PrivateEnvironmentConfigParameters) DeepCopy() *PrivateEnvironmentConfigParameters {
	if in == nil {
		return nil
	}
	out := new(PrivateEnvironmentConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerObservation) DeepCopyInto(out *SchedulerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerObservation.
func (in *SchedulerObservation) DeepCopy() *SchedulerObservation {
	if in == nil {
		return nil
	}
	out := new(SchedulerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerParameters) DeepCopyInto(out *SchedulerParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(float64)
		**out = **in
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(float64)
		**out = **in
	}
	if in.MemoryGb != nil {
		in, out := &in.MemoryGb, &out.MemoryGb
		*out = new(float64)
		**out = **in
	}
	if in.StorageGb != nil {
		in, out := &in.StorageGb, &out.StorageGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerParameters.
func (in *SchedulerParameters) DeepCopy() *SchedulerParameters {
	if in == nil {
		return nil
	}
	out := new(SchedulerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SoftwareConfigObservation) DeepCopyInto(out *SoftwareConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SoftwareConfigObservation.
func (in *SoftwareConfigObservation) DeepCopy() *SoftwareConfigObservation {
	if in == nil {
		return nil
	}
	out := new(SoftwareConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SoftwareConfigParameters) DeepCopyInto(out *SoftwareConfigParameters) {
	*out = *in
	if in.AirflowConfigOverrides != nil {
		in, out := &in.AirflowConfigOverrides, &out.AirflowConfigOverrides
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.EnvVariables != nil {
		in, out := &in.EnvVariables, &out.EnvVariables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ImageVersion != nil {
		in, out := &in.ImageVersion, &out.ImageVersion
		*out = new(string)
		**out = **in
	}
	if in.PypiPackages != nil {
		in, out := &in.PypiPackages, &out.PypiPackages
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PythonVersion != nil {
		in, out := &in.PythonVersion, &out.PythonVersion
		*out = new(string)
		**out = **in
	}
	if in.SchedulerCount != nil {
		in, out := &in.SchedulerCount, &out.SchedulerCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SoftwareConfigParameters.
func (in *SoftwareConfigParameters) DeepCopy() *SoftwareConfigParameters {
	if in == nil {
		return nil
	}
	out := new(SoftwareConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerConfigObservation) DeepCopyInto(out *WebServerConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerConfigObservation.
func (in *WebServerConfigObservation) DeepCopy() *WebServerConfigObservation {
	if in == nil {
		return nil
	}
	out := new(WebServerConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerConfigParameters) DeepCopyInto(out *WebServerConfigParameters) {
	*out = *in
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerConfigParameters.
func (in *WebServerConfigParameters) DeepCopy() *WebServerConfigParameters {
	if in == nil {
		return nil
	}
	out := new(WebServerConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerNetworkAccessControlObservation) DeepCopyInto(out *WebServerNetworkAccessControlObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerNetworkAccessControlObservation.
func (in *WebServerNetworkAccessControlObservation) DeepCopy() *WebServerNetworkAccessControlObservation {
	if in == nil {
		return nil
	}
	out := new(WebServerNetworkAccessControlObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerNetworkAccessControlParameters) DeepCopyInto(out *WebServerNetworkAccessControlParameters) {
	*out = *in
	if in.AllowedIPRange != nil {
		in, out := &in.AllowedIPRange, &out.AllowedIPRange
		*out = make([]AllowedIPRangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerNetworkAccessControlParameters.
func (in *WebServerNetworkAccessControlParameters) DeepCopy() *WebServerNetworkAccessControlParameters {
	if in == nil {
		return nil
	}
	out := new(WebServerNetworkAccessControlParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerObservation) DeepCopyInto(out *WebServerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerObservation.
func (in *WebServerObservation) DeepCopy() *WebServerObservation {
	if in == nil {
		return nil
	}
	out := new(WebServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebServerParameters) DeepCopyInto(out *WebServerParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(float64)
		**out = **in
	}
	if in.MemoryGb != nil {
		in, out := &in.MemoryGb, &out.MemoryGb
		*out = new(float64)
		**out = **in
	}
	if in.StorageGb != nil {
		in, out := &in.StorageGb, &out.StorageGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebServerParameters.
func (in *WebServerParameters) DeepCopy() *WebServerParameters {
	if in == nil {
		return nil
	}
	out := new(WebServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerObservation) DeepCopyInto(out *WorkerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerObservation.
func (in *WorkerObservation) DeepCopy() *WorkerObservation {
	if in == nil {
		return nil
	}
	out := new(WorkerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerParameters) DeepCopyInto(out *WorkerParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(float64)
		**out = **in
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(float64)
		**out = **in
	}
	if in.MemoryGb != nil {
		in, out := &in.MemoryGb, &out.MemoryGb
		*out = new(float64)
		**out = **in
	}
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(float64)
		**out = **in
	}
	if in.StorageGb != nil {
		in, out := &in.StorageGb, &out.StorageGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerParameters.
func (in *WorkerParameters) DeepCopy() *WorkerParameters {
	if in == nil {
		return nil
	}
	out := new(WorkerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadsConfigObservation) DeepCopyInto(out *WorkloadsConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadsConfigObservation.
func (in *WorkloadsConfigObservation) DeepCopy() *WorkloadsConfigObservation {
	if in == nil {
		return nil
	}
	out := new(WorkloadsConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadsConfigParameters) DeepCopyInto(out *WorkloadsConfigParameters) {
	*out = *in
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = make([]SchedulerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebServer != nil {
		in, out := &in.WebServer, &out.WebServer
		*out = make([]WebServerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = make([]WorkerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadsConfigParameters.
func (in *WorkloadsConfigParameters) DeepCopy() *WorkloadsConfigParameters {
	if in == nil {
		return nil
	}
	out := new(WorkloadsConfigParameters)
	in.DeepCopyInto(out)
	return out
}
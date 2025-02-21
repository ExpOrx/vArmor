//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfile) DeepCopyInto(out *ArmorProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfile.
func (in *ArmorProfile) DeepCopy() *ArmorProfile {
	if in == nil {
		return nil
	}
	out := new(ArmorProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmorProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileCondition) DeepCopyInto(out *ArmorProfileCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileCondition.
func (in *ArmorProfileCondition) DeepCopy() *ArmorProfileCondition {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileList) DeepCopyInto(out *ArmorProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmorProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileList.
func (in *ArmorProfileList) DeepCopy() *ArmorProfileList {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmorProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModel) DeepCopyInto(out *ArmorProfileModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModel.
func (in *ArmorProfileModel) DeepCopy() *ArmorProfileModel {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmorProfileModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModelCondition) DeepCopyInto(out *ArmorProfileModelCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModelCondition.
func (in *ArmorProfileModelCondition) DeepCopy() *ArmorProfileModelCondition {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModelCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModelList) DeepCopyInto(out *ArmorProfileModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmorProfileModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModelList.
func (in *ArmorProfileModelList) DeepCopy() *ArmorProfileModelList {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmorProfileModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModelSpec) DeepCopyInto(out *ArmorProfileModelSpec) {
	*out = *in
	in.DynamicResult.DeepCopyInto(&out.DynamicResult)
	out.StaticResult = in.StaticResult
	in.Profile.DeepCopyInto(&out.Profile)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModelSpec.
func (in *ArmorProfileModelSpec) DeepCopy() *ArmorProfileModelSpec {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileModelStatus) DeepCopyInto(out *ArmorProfileModelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ArmorProfileModelCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileModelStatus.
func (in *ArmorProfileModelStatus) DeepCopy() *ArmorProfileModelStatus {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileSpec) DeepCopyInto(out *ArmorProfileSpec) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	in.Profile.DeepCopyInto(&out.Profile)
	out.BehaviorModeling = in.BehaviorModeling
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileSpec.
func (in *ArmorProfileSpec) DeepCopy() *ArmorProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmorProfileStatus) DeepCopyInto(out *ArmorProfileStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ArmorProfileCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmorProfileStatus.
func (in *ArmorProfileStatus) DeepCopy() *ArmorProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ArmorProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttackProtectionRules) DeepCopyInto(out *AttackProtectionRules) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttackProtectionRules.
func (in *AttackProtectionRules) DeepCopy() *AttackProtectionRules {
	if in == nil {
		return nil
	}
	out := new(AttackProtectionRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BehaviorModeling) DeepCopyInto(out *BehaviorModeling) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BehaviorModeling.
func (in *BehaviorModeling) DeepCopy() *BehaviorModeling {
	if in == nil {
		return nil
	}
	out := new(BehaviorModeling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfContent) DeepCopyInto(out *BpfContent) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]FileContent, len(*in))
		copy(*out, *in)
	}
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make([]FileContent, len(*in))
		copy(*out, *in)
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworkContent, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfContent.
func (in *BpfContent) DeepCopy() *BpfContent {
	if in == nil {
		return nil
	}
	out := new(BpfContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfRawRules) DeepCopyInto(out *BpfRawRules) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]FileRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make([]FileRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Network.DeepCopyInto(&out.Network)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfRawRules.
func (in *BpfRawRules) DeepCopy() *BpfRawRules {
	if in == nil {
		return nil
	}
	out := new(BpfRawRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapabilitiesCapabilityType) DeepCopyInto(out *CapabilitiesCapabilityType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapabilitiesCapabilityType.
func (in *CapabilitiesCapabilityType) DeepCopy() *CapabilitiesCapabilityType {
	if in == nil {
		return nil
	}
	out := new(CapabilitiesCapabilityType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapabilitiesType) DeepCopyInto(out *CapabilitiesType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchCapabilities != nil {
		in, out := &in.MatchCapabilities, &out.MatchCapabilities
		*out = make([]CapabilitiesCapabilityType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapabilitiesType.
func (in *CapabilitiesType) DeepCopy() *CapabilitiesType {
	if in == nil {
		return nil
	}
	out := new(CapabilitiesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomPolicy) DeepCopyInto(out *CustomPolicy) {
	*out = *in
	in.Process.DeepCopyInto(&out.Process)
	in.File.DeepCopyInto(&out.File)
	in.Network.DeepCopyInto(&out.Network)
	in.Capabilities.DeepCopyInto(&out.Capabilities)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomPolicy.
func (in *CustomPolicy) DeepCopy() *CustomPolicy {
	if in == nil {
		return nil
	}
	out := new(CustomPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefenseInDepth) DeepCopyInto(out *DefenseInDepth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefenseInDepth.
func (in *DefenseInDepth) DeepCopy() *DefenseInDepth {
	if in == nil {
		return nil
	}
	out := new(DefenseInDepth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicResult) DeepCopyInto(out *DynamicResult) {
	*out = *in
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Executions != nil {
		in, out := &in.Executions, &out.Executions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]Network, len(*in))
		copy(*out, *in)
	}
	if in.Ptraces != nil {
		in, out := &in.Ptraces, &out.Ptraces
		*out = make([]Ptrace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Signals != nil {
		in, out := &in.Signals, &out.Signals
		*out = make([]Signal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Unhandled != nil {
		in, out := &in.Unhandled, &out.Unhandled
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicResult.
func (in *DynamicResult) DeepCopy() *DynamicResult {
	if in == nil {
		return nil
	}
	out := new(DynamicResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhanceProtect) DeepCopyInto(out *EnhanceProtect) {
	*out = *in
	if in.HardeningRules != nil {
		in, out := &in.HardeningRules, &out.HardeningRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AttackProtectionRules != nil {
		in, out := &in.AttackProtectionRules, &out.AttackProtectionRules
		*out = make([]AttackProtectionRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VulMitigationRules != nil {
		in, out := &in.VulMitigationRules, &out.VulMitigationRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AppArmorRawRules != nil {
		in, out := &in.AppArmorRawRules, &out.AppArmorRawRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.BpfRawRules.DeepCopyInto(&out.BpfRawRules)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhanceProtect.
func (in *EnhanceProtect) DeepCopy() *EnhanceProtect {
	if in == nil {
		return nil
	}
	out := new(EnhanceProtect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileContent) DeepCopyInto(out *FileContent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileContent.
func (in *FileContent) DeepCopy() *FileContent {
	if in == nil {
		return nil
	}
	out := new(FileContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileDirectoryType) DeepCopyInto(out *FileDirectoryType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileDirectoryType.
func (in *FileDirectoryType) DeepCopy() *FileDirectoryType {
	if in == nil {
		return nil
	}
	out := new(FileDirectoryType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilePathType) DeepCopyInto(out *FilePathType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilePathType.
func (in *FilePathType) DeepCopy() *FilePathType {
	if in == nil {
		return nil
	}
	out := new(FilePathType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilePatternType) DeepCopyInto(out *FilePatternType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilePatternType.
func (in *FilePatternType) DeepCopy() *FilePatternType {
	if in == nil {
		return nil
	}
	out := new(FilePatternType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileRule) DeepCopyInto(out *FileRule) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileRule.
func (in *FileRule) DeepCopy() *FileRule {
	if in == nil {
		return nil
	}
	out := new(FileRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileType) DeepCopyInto(out *FileType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchPaths != nil {
		in, out := &in.MatchPaths, &out.MatchPaths
		*out = make([]FilePathType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchDirectories != nil {
		in, out := &in.MatchDirectories, &out.MatchDirectories
		*out = make([]FileDirectoryType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchPatterns != nil {
		in, out := &in.MatchPatterns, &out.MatchPatterns
		*out = make([]FilePatternType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileType.
func (in *FileType) DeepCopy() *FileType {
	if in == nil {
		return nil
	}
	out := new(FileType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchSourceType) DeepCopyInto(out *MatchSourceType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchSourceType.
func (in *MatchSourceType) DeepCopy() *MatchSourceType {
	if in == nil {
		return nil
	}
	out := new(MatchSourceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkContent) DeepCopyInto(out *NetworkContent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkContent.
func (in *NetworkContent) DeepCopy() *NetworkContent {
	if in == nil {
		return nil
	}
	out := new(NetworkContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkEgressRule) DeepCopyInto(out *NetworkEgressRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkEgressRule.
func (in *NetworkEgressRule) DeepCopy() *NetworkEgressRule {
	if in == nil {
		return nil
	}
	out := new(NetworkEgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkProtocolType) DeepCopyInto(out *NetworkProtocolType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkProtocolType.
func (in *NetworkProtocolType) DeepCopy() *NetworkProtocolType {
	if in == nil {
		return nil
	}
	out := new(NetworkProtocolType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRule) DeepCopyInto(out *NetworkRule) {
	*out = *in
	if in.Egresses != nil {
		in, out := &in.Egresses, &out.Egresses
		*out = make([]NetworkEgressRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRule.
func (in *NetworkRule) DeepCopy() *NetworkRule {
	if in == nil {
		return nil
	}
	out := new(NetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkType) DeepCopyInto(out *NetworkType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchProtocols != nil {
		in, out := &in.MatchProtocols, &out.MatchProtocols
		*out = make([]NetworkProtocolType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkType.
func (in *NetworkType) DeepCopy() *NetworkType {
	if in == nil {
		return nil
	}
	out := new(NetworkType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	in.EnhanceProtect.DeepCopyInto(&out.EnhanceProtect)
	in.CustomPolicy.DeepCopyInto(&out.CustomPolicy)
	out.DefenseInDepth = in.DefenseInDepth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessDirectoryType) DeepCopyInto(out *ProcessDirectoryType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessDirectoryType.
func (in *ProcessDirectoryType) DeepCopy() *ProcessDirectoryType {
	if in == nil {
		return nil
	}
	out := new(ProcessDirectoryType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessPathType) DeepCopyInto(out *ProcessPathType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessPathType.
func (in *ProcessPathType) DeepCopy() *ProcessPathType {
	if in == nil {
		return nil
	}
	out := new(ProcessPathType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessPatternType) DeepCopyInto(out *ProcessPatternType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessPatternType.
func (in *ProcessPatternType) DeepCopy() *ProcessPatternType {
	if in == nil {
		return nil
	}
	out := new(ProcessPatternType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessType) DeepCopyInto(out *ProcessType) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchPaths != nil {
		in, out := &in.MatchPaths, &out.MatchPaths
		*out = make([]ProcessPathType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchDirectories != nil {
		in, out := &in.MatchDirectories, &out.MatchDirectories
		*out = make([]ProcessDirectoryType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchPatterns != nil {
		in, out := &in.MatchPatterns, &out.MatchPatterns
		*out = make([]ProcessPatternType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessType.
func (in *ProcessType) DeepCopy() *ProcessType {
	if in == nil {
		return nil
	}
	out := new(ProcessType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Profile) DeepCopyInto(out *Profile) {
	*out = *in
	in.BpfContent.DeepCopyInto(&out.BpfContent)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Profile.
func (in *Profile) DeepCopy() *Profile {
	if in == nil {
		return nil
	}
	out := new(Profile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ptrace) DeepCopyInto(out *Ptrace) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ptrace.
func (in *Ptrace) DeepCopy() *Ptrace {
	if in == nil {
		return nil
	}
	out := new(Ptrace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Signal) DeepCopyInto(out *Signal) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Signals != nil {
		in, out := &in.Signals, &out.Signals
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Signal.
func (in *Signal) DeepCopy() *Signal {
	if in == nil {
		return nil
	}
	out := new(Signal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticResult) DeepCopyInto(out *StaticResult) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticResult.
func (in *StaticResult) DeepCopy() *StaticResult {
	if in == nil {
		return nil
	}
	out := new(StaticResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicy) DeepCopyInto(out *VarmorPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicy.
func (in *VarmorPolicy) DeepCopy() *VarmorPolicy {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VarmorPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicyCondition) DeepCopyInto(out *VarmorPolicyCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicyCondition.
func (in *VarmorPolicyCondition) DeepCopy() *VarmorPolicyCondition {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicyList) DeepCopyInto(out *VarmorPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VarmorPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicyList.
func (in *VarmorPolicyList) DeepCopy() *VarmorPolicyList {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VarmorPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicySpec) DeepCopyInto(out *VarmorPolicySpec) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	in.Policy.DeepCopyInto(&out.Policy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicySpec.
func (in *VarmorPolicySpec) DeepCopy() *VarmorPolicySpec {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarmorPolicyStatus) DeepCopyInto(out *VarmorPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VarmorPolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarmorPolicyStatus.
func (in *VarmorPolicyStatus) DeepCopy() *VarmorPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(VarmorPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

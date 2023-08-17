//go:build !ignore_autogenerated

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

package controlplane

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericControlPlane) DeepCopyInto(out *GenericControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericControlPlane.
func (in *GenericControlPlane) DeepCopy() *GenericControlPlane {
	if in == nil {
		return nil
	}
	out := new(GenericControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GenericControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericControlPlaneList) DeepCopyInto(out *GenericControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GenericControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericControlPlaneList.
func (in *GenericControlPlaneList) DeepCopy() *GenericControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(GenericControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GenericControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericControlPlaneSpec) DeepCopyInto(out *GenericControlPlaneSpec) {
	*out = *in
	out.MachineTemplate = in.MachineTemplate
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericControlPlaneSpec.
func (in *GenericControlPlaneSpec) DeepCopy() *GenericControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(GenericControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericMachineTemplate) DeepCopyInto(out *GenericMachineTemplate) {
	*out = *in
	out.InfrastructureRef = in.InfrastructureRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericMachineTemplate.
func (in *GenericMachineTemplate) DeepCopy() *GenericMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(GenericMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

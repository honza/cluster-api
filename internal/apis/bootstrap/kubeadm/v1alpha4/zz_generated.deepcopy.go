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

package v1alpha4

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1alpha4 "sigs.k8s.io/cluster-api/internal/apis/core/v1alpha4"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIEndpoint) DeepCopyInto(out *APIEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIEndpoint.
func (in *APIEndpoint) DeepCopy() *APIEndpoint {
	if in == nil {
		return nil
	}
	out := new(APIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServer) DeepCopyInto(out *APIServer) {
	*out = *in
	in.ControlPlaneComponent.DeepCopyInto(&out.ControlPlaneComponent)
	if in.CertSANs != nil {
		in, out := &in.CertSANs, &out.CertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeoutForControlPlane != nil {
		in, out := &in.TimeoutForControlPlane, &out.TimeoutForControlPlane
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServer.
func (in *APIServer) DeepCopy() *APIServer {
	if in == nil {
		return nil
	}
	out := new(APIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapToken) DeepCopyInto(out *BootstrapToken) {
	*out = *in
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(BootstrapTokenString)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Expires != nil {
		in, out := &in.Expires, &out.Expires
		*out = (*in).DeepCopy()
	}
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapToken.
func (in *BootstrapToken) DeepCopy() *BootstrapToken {
	if in == nil {
		return nil
	}
	out := new(BootstrapToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapTokenDiscovery) DeepCopyInto(out *BootstrapTokenDiscovery) {
	*out = *in
	if in.CACertHashes != nil {
		in, out := &in.CACertHashes, &out.CACertHashes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapTokenDiscovery.
func (in *BootstrapTokenDiscovery) DeepCopy() *BootstrapTokenDiscovery {
	if in == nil {
		return nil
	}
	out := new(BootstrapTokenDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapTokenString) DeepCopyInto(out *BootstrapTokenString) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapTokenString.
func (in *BootstrapTokenString) DeepCopy() *BootstrapTokenString {
	if in == nil {
		return nil
	}
	out := new(BootstrapTokenString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfiguration) DeepCopyInto(out *ClusterConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Etcd.DeepCopyInto(&out.Etcd)
	out.Networking = in.Networking
	in.APIServer.DeepCopyInto(&out.APIServer)
	in.ControllerManager.DeepCopyInto(&out.ControllerManager)
	in.Scheduler.DeepCopyInto(&out.Scheduler)
	out.DNS = in.DNS
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfiguration.
func (in *ClusterConfiguration) DeepCopy() *ClusterConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClusterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]APIEndpoint, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneComponent) DeepCopyInto(out *ControlPlaneComponent) {
	*out = *in
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]HostPathMount, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneComponent.
func (in *ControlPlaneComponent) DeepCopy() *ControlPlaneComponent {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNS) DeepCopyInto(out *DNS) {
	*out = *in
	out.ImageMeta = in.ImageMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNS.
func (in *DNS) DeepCopy() *DNS {
	if in == nil {
		return nil
	}
	out := new(DNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Discovery) DeepCopyInto(out *Discovery) {
	*out = *in
	if in.BootstrapToken != nil {
		in, out := &in.BootstrapToken, &out.BootstrapToken
		*out = new(BootstrapTokenDiscovery)
		(*in).DeepCopyInto(*out)
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileDiscovery)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Discovery.
func (in *Discovery) DeepCopy() *Discovery {
	if in == nil {
		return nil
	}
	out := new(Discovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSetup) DeepCopyInto(out *DiskSetup) {
	*out = *in
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = make([]Partition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filesystems != nil {
		in, out := &in.Filesystems, &out.Filesystems
		*out = make([]Filesystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSetup.
func (in *DiskSetup) DeepCopy() *DiskSetup {
	if in == nil {
		return nil
	}
	out := new(DiskSetup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalEtcd)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalEtcd)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalEtcd) DeepCopyInto(out *ExternalEtcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalEtcd.
func (in *ExternalEtcd) DeepCopy() *ExternalEtcd {
	if in == nil {
		return nil
	}
	out := new(ExternalEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	if in.ContentFrom != nil {
		in, out := &in.ContentFrom, &out.ContentFrom
		*out = new(FileSource)
		**out = **in
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
func (in *FileDiscovery) DeepCopyInto(out *FileDiscovery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileDiscovery.
func (in *FileDiscovery) DeepCopy() *FileDiscovery {
	if in == nil {
		return nil
	}
	out := new(FileDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSource) DeepCopyInto(out *FileSource) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSource.
func (in *FileSource) DeepCopy() *FileSource {
	if in == nil {
		return nil
	}
	out := new(FileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filesystem) DeepCopyInto(out *Filesystem) {
	*out = *in
	if in.Partition != nil {
		in, out := &in.Partition, &out.Partition
		*out = new(string)
		**out = **in
	}
	if in.Overwrite != nil {
		in, out := &in.Overwrite, &out.Overwrite
		*out = new(bool)
		**out = **in
	}
	if in.ReplaceFS != nil {
		in, out := &in.ReplaceFS, &out.ReplaceFS
		*out = new(string)
		**out = **in
	}
	if in.ExtraOpts != nil {
		in, out := &in.ExtraOpts, &out.ExtraOpts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filesystem.
func (in *Filesystem) DeepCopy() *Filesystem {
	if in == nil {
		return nil
	}
	out := new(Filesystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPathMount) DeepCopyInto(out *HostPathMount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPathMount.
func (in *HostPathMount) DeepCopy() *HostPathMount {
	if in == nil {
		return nil
	}
	out := new(HostPathMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageMeta) DeepCopyInto(out *ImageMeta) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageMeta.
func (in *ImageMeta) DeepCopy() *ImageMeta {
	if in == nil {
		return nil
	}
	out := new(ImageMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitConfiguration) DeepCopyInto(out *InitConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.BootstrapTokens != nil {
		in, out := &in.BootstrapTokens, &out.BootstrapTokens
		*out = make([]BootstrapToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.NodeRegistration.DeepCopyInto(&out.NodeRegistration)
	out.LocalAPIEndpoint = in.LocalAPIEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitConfiguration.
func (in *InitConfiguration) DeepCopy() *InitConfiguration {
	if in == nil {
		return nil
	}
	out := new(InitConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InitConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JoinConfiguration) DeepCopyInto(out *JoinConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.NodeRegistration.DeepCopyInto(&out.NodeRegistration)
	in.Discovery.DeepCopyInto(&out.Discovery)
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(JoinControlPlane)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JoinConfiguration.
func (in *JoinConfiguration) DeepCopy() *JoinConfiguration {
	if in == nil {
		return nil
	}
	out := new(JoinConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JoinConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JoinControlPlane) DeepCopyInto(out *JoinControlPlane) {
	*out = *in
	out.LocalAPIEndpoint = in.LocalAPIEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JoinControlPlane.
func (in *JoinControlPlane) DeepCopy() *JoinControlPlane {
	if in == nil {
		return nil
	}
	out := new(JoinControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfig) DeepCopyInto(out *KubeadmConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfig.
func (in *KubeadmConfig) DeepCopy() *KubeadmConfig {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfigList) DeepCopyInto(out *KubeadmConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfigList.
func (in *KubeadmConfigList) DeepCopy() *KubeadmConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfigSpec) DeepCopyInto(out *KubeadmConfigSpec) {
	*out = *in
	if in.ClusterConfiguration != nil {
		in, out := &in.ClusterConfiguration, &out.ClusterConfiguration
		*out = new(ClusterConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.InitConfiguration != nil {
		in, out := &in.InitConfiguration, &out.InitConfiguration
		*out = new(InitConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.JoinConfiguration != nil {
		in, out := &in.JoinConfiguration, &out.JoinConfiguration
		*out = new(JoinConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DiskSetup != nil {
		in, out := &in.DiskSetup, &out.DiskSetup
		*out = new(DiskSetup)
		(*in).DeepCopyInto(*out)
	}
	if in.Mounts != nil {
		in, out := &in.Mounts, &out.Mounts
		*out = make([]MountPoints, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(MountPoints, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.PreKubeadmCommands != nil {
		in, out := &in.PreKubeadmCommands, &out.PreKubeadmCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PostKubeadmCommands != nil {
		in, out := &in.PostKubeadmCommands, &out.PostKubeadmCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NTP != nil {
		in, out := &in.NTP, &out.NTP
		*out = new(NTP)
		(*in).DeepCopyInto(*out)
	}
	if in.Verbosity != nil {
		in, out := &in.Verbosity, &out.Verbosity
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfigSpec.
func (in *KubeadmConfigSpec) DeepCopy() *KubeadmConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfigStatus) DeepCopyInto(out *KubeadmConfigStatus) {
	*out = *in
	if in.DataSecretName != nil {
		in, out := &in.DataSecretName, &out.DataSecretName
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfigStatus.
func (in *KubeadmConfigStatus) DeepCopy() *KubeadmConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfigTemplate) DeepCopyInto(out *KubeadmConfigTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfigTemplate.
func (in *KubeadmConfigTemplate) DeepCopy() *KubeadmConfigTemplate {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfigTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmConfigTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfigTemplateList) DeepCopyInto(out *KubeadmConfigTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmConfigTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfigTemplateList.
func (in *KubeadmConfigTemplateList) DeepCopy() *KubeadmConfigTemplateList {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfigTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeadmConfigTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfigTemplateResource) DeepCopyInto(out *KubeadmConfigTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfigTemplateResource.
func (in *KubeadmConfigTemplateResource) DeepCopy() *KubeadmConfigTemplateResource {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfigTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeadmConfigTemplateSpec) DeepCopyInto(out *KubeadmConfigTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeadmConfigTemplateSpec.
func (in *KubeadmConfigTemplateSpec) DeepCopy() *KubeadmConfigTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(KubeadmConfigTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalEtcd) DeepCopyInto(out *LocalEtcd) {
	*out = *in
	out.ImageMeta = in.ImageMeta
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServerCertSANs != nil {
		in, out := &in.ServerCertSANs, &out.ServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PeerCertSANs != nil {
		in, out := &in.PeerCertSANs, &out.PeerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalEtcd.
func (in *LocalEtcd) DeepCopy() *LocalEtcd {
	if in == nil {
		return nil
	}
	out := new(LocalEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MountPoints) DeepCopyInto(out *MountPoints) {
	{
		in := &in
		*out = make(MountPoints, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountPoints.
func (in MountPoints) DeepCopy() MountPoints {
	if in == nil {
		return nil
	}
	out := new(MountPoints)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NTP) DeepCopyInto(out *NTP) {
	*out = *in
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NTP.
func (in *NTP) DeepCopy() *NTP {
	if in == nil {
		return nil
	}
	out := new(NTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (in *Networking) DeepCopy() *Networking {
	if in == nil {
		return nil
	}
	out := new(Networking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRegistrationOptions) DeepCopyInto(out *NodeRegistrationOptions) {
	*out = *in
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]corev1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubeletExtraArgs != nil {
		in, out := &in.KubeletExtraArgs, &out.KubeletExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IgnorePreflightErrors != nil {
		in, out := &in.IgnorePreflightErrors, &out.IgnorePreflightErrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRegistrationOptions.
func (in *NodeRegistrationOptions) DeepCopy() *NodeRegistrationOptions {
	if in == nil {
		return nil
	}
	out := new(NodeRegistrationOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Partition) DeepCopyInto(out *Partition) {
	*out = *in
	if in.Overwrite != nil {
		in, out := &in.Overwrite, &out.Overwrite
		*out = new(bool)
		**out = **in
	}
	if in.TableType != nil {
		in, out := &in.TableType, &out.TableType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Partition.
func (in *Partition) DeepCopy() *Partition {
	if in == nil {
		return nil
	}
	out := new(Partition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretFileSource) DeepCopyInto(out *SecretFileSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretFileSource.
func (in *SecretFileSource) DeepCopy() *SecretFileSource {
	if in == nil {
		return nil
	}
	out := new(SecretFileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	if in.Gecos != nil {
		in, out := &in.Gecos, &out.Gecos
		*out = new(string)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = new(string)
		**out = **in
	}
	if in.HomeDir != nil {
		in, out := &in.HomeDir, &out.HomeDir
		*out = new(string)
		**out = **in
	}
	if in.Inactive != nil {
		in, out := &in.Inactive, &out.Inactive
		*out = new(bool)
		**out = **in
	}
	if in.Shell != nil {
		in, out := &in.Shell, &out.Shell
		*out = new(string)
		**out = **in
	}
	if in.Passwd != nil {
		in, out := &in.Passwd, &out.Passwd
		*out = new(string)
		**out = **in
	}
	if in.PrimaryGroup != nil {
		in, out := &in.PrimaryGroup, &out.PrimaryGroup
		*out = new(string)
		**out = **in
	}
	if in.LockPassword != nil {
		in, out := &in.LockPassword, &out.LockPassword
		*out = new(bool)
		**out = **in
	}
	if in.Sudo != nil {
		in, out := &in.Sudo, &out.Sudo
		*out = new(string)
		**out = **in
	}
	if in.SSHAuthorizedKeys != nil {
		in, out := &in.SSHAuthorizedKeys, &out.SSHAuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

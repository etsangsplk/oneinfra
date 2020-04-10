// +build !ignore_autogenerated

/*
Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>

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

package v1alpha1

import (
	commonv1alpha1 "github.com/oneinfra/oneinfra/apis/common/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ClusterFileMap) DeepCopyInto(out *ClusterFileMap) {
	{
		in := &in
		*out = make(ClusterFileMap, len(*in))
		for key, val := range *in {
			var outVal map[string]FileMap
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ComponentFileMap, len(*in))
				for key, val := range *in {
					var outVal map[string]string
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = make(FileMap, len(*in))
						for key, val := range *in {
							(*out)[key] = val
						}
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFileMap.
func (in ClusterFileMap) DeepCopy() ClusterFileMap {
	if in == nil {
		return nil
	}
	out := new(ClusterFileMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ComponentFileMap) DeepCopyInto(out *ComponentFileMap) {
	{
		in := &in
		*out = make(ComponentFileMap, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(FileMap, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentFileMap.
func (in ComponentFileMap) DeepCopy() ComponentFileMap {
	if in == nil {
		return nil
	}
	out := new(ComponentFileMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FileMap) DeepCopyInto(out *FileMap) {
	{
		in := &in
		*out = make(FileMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileMap.
func (in FileMap) DeepCopy() FileMap {
	if in == nil {
		return nil
	}
	out := new(FileMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hypervisor) DeepCopyInto(out *Hypervisor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hypervisor.
func (in *Hypervisor) DeepCopy() *Hypervisor {
	if in == nil {
		return nil
	}
	out := new(Hypervisor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hypervisor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypervisorList) DeepCopyInto(out *HypervisorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hypervisor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypervisorList.
func (in *HypervisorList) DeepCopy() *HypervisorList {
	if in == nil {
		return nil
	}
	out := new(HypervisorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HypervisorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypervisorPortAllocation) DeepCopyInto(out *HypervisorPortAllocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypervisorPortAllocation.
func (in *HypervisorPortAllocation) DeepCopy() *HypervisorPortAllocation {
	if in == nil {
		return nil
	}
	out := new(HypervisorPortAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypervisorPortRange) DeepCopyInto(out *HypervisorPortRange) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypervisorPortRange.
func (in *HypervisorPortRange) DeepCopy() *HypervisorPortRange {
	if in == nil {
		return nil
	}
	out := new(HypervisorPortRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypervisorSpec) DeepCopyInto(out *HypervisorSpec) {
	*out = *in
	if in.LocalCRIEndpoint != nil {
		in, out := &in.LocalCRIEndpoint, &out.LocalCRIEndpoint
		*out = new(LocalHypervisorCRIEndpoint)
		**out = **in
	}
	if in.RemoteCRIEndpoint != nil {
		in, out := &in.RemoteCRIEndpoint, &out.RemoteCRIEndpoint
		*out = new(RemoteHypervisorCRIEndpoint)
		(*in).DeepCopyInto(*out)
	}
	out.PortRange = in.PortRange
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypervisorSpec.
func (in *HypervisorSpec) DeepCopy() *HypervisorSpec {
	if in == nil {
		return nil
	}
	out := new(HypervisorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypervisorStatus) DeepCopyInto(out *HypervisorStatus) {
	*out = *in
	if in.AllocatedPorts != nil {
		in, out := &in.AllocatedPorts, &out.AllocatedPorts
		*out = make([]HypervisorPortAllocation, len(*in))
		copy(*out, *in)
	}
	if in.FreedPorts != nil {
		in, out := &in.FreedPorts, &out.FreedPorts
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make(ClusterFileMap, len(*in))
		for key, val := range *in {
			var outVal map[string]FileMap
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ComponentFileMap, len(*in))
				for key, val := range *in {
					var outVal map[string]string
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = make(FileMap, len(*in))
						for key, val := range *in {
							(*out)[key] = val
						}
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypervisorStatus.
func (in *HypervisorStatus) DeepCopy() *HypervisorStatus {
	if in == nil {
		return nil
	}
	out := new(HypervisorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalHypervisorCRIEndpoint) DeepCopyInto(out *LocalHypervisorCRIEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalHypervisorCRIEndpoint.
func (in *LocalHypervisorCRIEndpoint) DeepCopy() *LocalHypervisorCRIEndpoint {
	if in == nil {
		return nil
	}
	out := new(LocalHypervisorCRIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteHypervisorCRIEndpoint) DeepCopyInto(out *RemoteHypervisorCRIEndpoint) {
	*out = *in
	if in.ClientCertificate != nil {
		in, out := &in.ClientCertificate, &out.ClientCertificate
		*out = new(commonv1alpha1.Certificate)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteHypervisorCRIEndpoint.
func (in *RemoteHypervisorCRIEndpoint) DeepCopy() *RemoteHypervisorCRIEndpoint {
	if in == nil {
		return nil
	}
	out := new(RemoteHypervisorCRIEndpoint)
	in.DeepCopyInto(out)
	return out
}

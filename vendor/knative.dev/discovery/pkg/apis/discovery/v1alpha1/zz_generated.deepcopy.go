// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDuckType) DeepCopyInto(out *ClusterDuckType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDuckType.
func (in *ClusterDuckType) DeepCopy() *ClusterDuckType {
	if in == nil {
		return nil
	}
	out := new(ClusterDuckType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDuckType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDuckTypeList) DeepCopyInto(out *ClusterDuckTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterDuckType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDuckTypeList.
func (in *ClusterDuckTypeList) DeepCopy() *ClusterDuckTypeList {
	if in == nil {
		return nil
	}
	out := new(ClusterDuckTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDuckTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDuckTypeSpec) DeepCopyInto(out *ClusterDuckTypeSpec) {
	*out = *in
	out.Names = in.Names
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]DuckVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]CustomResourceDefinitionSelector, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDuckTypeSpec.
func (in *ClusterDuckTypeSpec) DeepCopy() *ClusterDuckTypeSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterDuckTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDuckTypeStatus) DeepCopyInto(out *ClusterDuckTypeStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Ducks != nil {
		in, out := &in.Ducks, &out.Ducks
		*out = make(map[string][]ResourceMeta, len(*in))
		for key, val := range *in {
			var outVal []ResourceMeta
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]ResourceMeta, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDuckTypeStatus.
func (in *ClusterDuckTypeStatus) DeepCopy() *ClusterDuckTypeStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterDuckTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionSelector) DeepCopyInto(out *CustomResourceDefinitionSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionSelector.
func (in *CustomResourceDefinitionSelector) DeepCopy() *CustomResourceDefinitionSelector {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckTypeNames) DeepCopyInto(out *DuckTypeNames) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckTypeNames.
func (in *DuckTypeNames) DeepCopy() *DuckTypeNames {
	if in == nil {
		return nil
	}
	out := new(DuckTypeNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckVersion) DeepCopyInto(out *DuckVersion) {
	*out = *in
	if in.Refs != nil {
		in, out := &in.Refs, &out.Refs
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalPrinterColumns != nil {
		in, out := &in.AdditionalPrinterColumns, &out.AdditionalPrinterColumns
		*out = make([]v1.CustomResourceColumnDefinition, len(*in))
		copy(*out, *in)
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(v1.CustomResourceValidation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckVersion.
func (in *DuckVersion) DeepCopy() *DuckVersion {
	if in == nil {
		return nil
	}
	out := new(DuckVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMeta) DeepCopyInto(out *ResourceMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMeta.
func (in *ResourceMeta) DeepCopy() *ResourceMeta {
	if in == nil {
		return nil
	}
	out := new(ResourceMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}

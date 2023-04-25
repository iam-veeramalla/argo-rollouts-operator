//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoRollout) DeepCopyInto(out *ArgoRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoRollout.
func (in *ArgoRollout) DeepCopy() *ArgoRollout {
	if in == nil {
		return nil
	}
	out := new(ArgoRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArgoRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoRolloutList) DeepCopyInto(out *ArgoRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArgoRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoRolloutList.
func (in *ArgoRolloutList) DeepCopy() *ArgoRolloutList {
	if in == nil {
		return nil
	}
	out := new(ArgoRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArgoRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoRolloutSpec) DeepCopyInto(out *ArgoRolloutSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraCommandArgs != nil {
		in, out := &in.ExtraCommandArgs, &out.ExtraCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodePlacement != nil {
		in, out := &in.NodePlacement, &out.NodePlacement
		*out = new(ArgoRolloutsNodePlacementSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoRolloutSpec.
func (in *ArgoRolloutSpec) DeepCopy() *ArgoRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoRolloutStatus) DeepCopyInto(out *ArgoRolloutStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoRolloutStatus.
func (in *ArgoRolloutStatus) DeepCopy() *ArgoRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(ArgoRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoRolloutsNodePlacementSpec) DeepCopyInto(out *ArgoRolloutsNodePlacementSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoRolloutsNodePlacementSpec.
func (in *ArgoRolloutsNodePlacementSpec) DeepCopy() *ArgoRolloutsNodePlacementSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoRolloutsNodePlacementSpec)
	in.DeepCopyInto(out)
	return out
}
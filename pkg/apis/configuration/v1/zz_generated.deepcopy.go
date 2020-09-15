// +build !ignore_autogenerated

/*
Copyright 2018 The Kong Authors.

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

package v1

import (
	kong "github.com/kong/go-kong/kong"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSource) DeepCopyInto(out *ConfigSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.SecretValue = in.SecretValue
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSource.
func (in *ConfigSource) DeepCopy() *ConfigSource {
	if in == nil {
		return nil
	}
	out := new(ConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongClusterPlugin) DeepCopyInto(out *KongClusterPlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
	out.ConfigFrom = in.ConfigFrom
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongClusterPlugin.
func (in *KongClusterPlugin) DeepCopy() *KongClusterPlugin {
	if in == nil {
		return nil
	}
	out := new(KongClusterPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongClusterPlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongClusterPluginList) DeepCopyInto(out *KongClusterPluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongClusterPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongClusterPluginList.
func (in *KongClusterPluginList) DeepCopy() *KongClusterPluginList {
	if in == nil {
		return nil
	}
	out := new(KongClusterPluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongClusterPluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongConsumer) DeepCopyInto(out *KongConsumer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongConsumer.
func (in *KongConsumer) DeepCopy() *KongConsumer {
	if in == nil {
		return nil
	}
	out := new(KongConsumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongConsumer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongConsumerList) DeepCopyInto(out *KongConsumerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongConsumer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongConsumerList.
func (in *KongConsumerList) DeepCopy() *KongConsumerList {
	if in == nil {
		return nil
	}
	out := new(KongConsumerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongConsumerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongCredential.
func (in *KongCredential) DeepCopy() *KongCredential {
	if in == nil {
		return nil
	}
	out := new(KongCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongCredentialList) DeepCopyInto(out *KongCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongCredentialList.
func (in *KongCredentialList) DeepCopy() *KongCredentialList {
	if in == nil {
		return nil
	}
	out := new(KongCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngress) DeepCopyInto(out *KongIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Upstream != nil {
		in, out := &in.Upstream, &out.Upstream
		*out = new(kong.Upstream)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(kong.Service)
		(*in).DeepCopyInto(*out)
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(kong.Route)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngress.
func (in *KongIngress) DeepCopy() *KongIngress {
	if in == nil {
		return nil
	}
	out := new(KongIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongIngressList) DeepCopyInto(out *KongIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongIngressList.
func (in *KongIngressList) DeepCopy() *KongIngressList {
	if in == nil {
		return nil
	}
	out := new(KongIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongPlugin) DeepCopyInto(out *KongPlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
	out.ConfigFrom = in.ConfigFrom
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongPlugin.
func (in *KongPlugin) DeepCopy() *KongPlugin {
	if in == nil {
		return nil
	}
	out := new(KongPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongPlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongPluginList) DeepCopyInto(out *KongPluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongPluginList.
func (in *KongPluginList) DeepCopy() *KongPluginList {
	if in == nil {
		return nil
	}
	out := new(KongPluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongPluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedConfigSource) DeepCopyInto(out *NamespacedConfigSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.SecretValue = in.SecretValue
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedConfigSource.
func (in *NamespacedConfigSource) DeepCopy() *NamespacedConfigSource {
	if in == nil {
		return nil
	}
	out := new(NamespacedConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedConfigSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedSecretValueFromSource) DeepCopyInto(out *NamespacedSecretValueFromSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedSecretValueFromSource.
func (in *NamespacedSecretValueFromSource) DeepCopy() *NamespacedSecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(NamespacedSecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedSecretValueFromSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretValueFromSource) DeepCopyInto(out *SecretValueFromSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretValueFromSource.
func (in *SecretValueFromSource) DeepCopy() *SecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(SecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretValueFromSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

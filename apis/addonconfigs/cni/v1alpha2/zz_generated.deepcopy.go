//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Antrea) DeepCopyInto(out *Antrea) {
	*out = *in
	in.AntreaConfigDataValue.DeepCopyInto(&out.AntreaConfigDataValue)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Antrea.
func (in *Antrea) DeepCopy() *Antrea {
	if in == nil {
		return nil
	}
	out := new(Antrea)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaConfig) DeepCopyInto(out *AntreaConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaConfig.
func (in *AntreaConfig) DeepCopy() *AntreaConfig {
	if in == nil {
		return nil
	}
	out := new(AntreaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaConfigDataValue) DeepCopyInto(out *AntreaConfigDataValue) {
	*out = *in
	in.Egress.DeepCopyInto(&out.Egress)
	out.NodePortLocal = in.NodePortLocal
	in.AntreaProxy.DeepCopyInto(&out.AntreaProxy)
	out.AntreaFlowExporter = in.AntreaFlowExporter
	out.Multicast = in.Multicast
	out.MultiCluster = in.MultiCluster
	if in.TransportInterfaceCIDRs != nil {
		in, out := &in.TransportInterfaceCIDRs, &out.TransportInterfaceCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MulticastInterfaces != nil {
		in, out := &in.MulticastInterfaces, &out.MulticastInterfaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.WireGuard = in.WireGuard
	out.FeatureGates = in.FeatureGates
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaConfigDataValue.
func (in *AntreaConfigDataValue) DeepCopy() *AntreaConfigDataValue {
	if in == nil {
		return nil
	}
	out := new(AntreaConfigDataValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaConfigList) DeepCopyInto(out *AntreaConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AntreaConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaConfigList.
func (in *AntreaConfigList) DeepCopy() *AntreaConfigList {
	if in == nil {
		return nil
	}
	out := new(AntreaConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AntreaConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaConfigSpec) DeepCopyInto(out *AntreaConfigSpec) {
	*out = *in
	in.Antrea.DeepCopyInto(&out.Antrea)
	in.AntreaNsx.DeepCopyInto(&out.AntreaNsx)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaConfigSpec.
func (in *AntreaConfigSpec) DeepCopy() *AntreaConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AntreaConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaConfigStatus) DeepCopyInto(out *AntreaConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaConfigStatus.
func (in *AntreaConfigStatus) DeepCopy() *AntreaConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AntreaConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaEgress) DeepCopyInto(out *AntreaEgress) {
	*out = *in
	if in.EgressExceptCIDRs != nil {
		in, out := &in.EgressExceptCIDRs, &out.EgressExceptCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaEgress.
func (in *AntreaEgress) DeepCopy() *AntreaEgress {
	if in == nil {
		return nil
	}
	out := new(AntreaEgress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaFeatureGates) DeepCopyInto(out *AntreaFeatureGates) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaFeatureGates.
func (in *AntreaFeatureGates) DeepCopy() *AntreaFeatureGates {
	if in == nil {
		return nil
	}
	out := new(AntreaFeatureGates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaFlowExporter) DeepCopyInto(out *AntreaFlowExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaFlowExporter.
func (in *AntreaFlowExporter) DeepCopy() *AntreaFlowExporter {
	if in == nil {
		return nil
	}
	out := new(AntreaFlowExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaMultiCluster) DeepCopyInto(out *AntreaMultiCluster) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaMultiCluster.
func (in *AntreaMultiCluster) DeepCopy() *AntreaMultiCluster {
	if in == nil {
		return nil
	}
	out := new(AntreaMultiCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaMulticast) DeepCopyInto(out *AntreaMulticast) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaMulticast.
func (in *AntreaMulticast) DeepCopy() *AntreaMulticast {
	if in == nil {
		return nil
	}
	out := new(AntreaMulticast)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaNodePortLocal) DeepCopyInto(out *AntreaNodePortLocal) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaNodePortLocal.
func (in *AntreaNodePortLocal) DeepCopy() *AntreaNodePortLocal {
	if in == nil {
		return nil
	}
	out := new(AntreaNodePortLocal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaNsx) DeepCopyInto(out *AntreaNsx) {
	*out = *in
	in.BootstrapFrom.DeepCopyInto(&out.BootstrapFrom)
	out.AntreaNsxConfig = in.AntreaNsxConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaNsx.
func (in *AntreaNsx) DeepCopy() *AntreaNsx {
	if in == nil {
		return nil
	}
	out := new(AntreaNsx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaNsxBootstrapFrom) DeepCopyInto(out *AntreaNsxBootstrapFrom) {
	*out = *in
	if in.ProviderRef != nil {
		in, out := &in.ProviderRef, &out.ProviderRef
		*out = new(AntreaNsxProvider)
		**out = **in
	}
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(AntreaNsxInline)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaNsxBootstrapFrom.
func (in *AntreaNsxBootstrapFrom) DeepCopy() *AntreaNsxBootstrapFrom {
	if in == nil {
		return nil
	}
	out := new(AntreaNsxBootstrapFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaNsxConfig) DeepCopyInto(out *AntreaNsxConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaNsxConfig.
func (in *AntreaNsxConfig) DeepCopy() *AntreaNsxConfig {
	if in == nil {
		return nil
	}
	out := new(AntreaNsxConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaNsxInline) DeepCopyInto(out *AntreaNsxInline) {
	*out = *in
	if in.NsxManagers != nil {
		in, out := &in.NsxManagers, &out.NsxManagers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaNsxInline.
func (in *AntreaNsxInline) DeepCopy() *AntreaNsxInline {
	if in == nil {
		return nil
	}
	out := new(AntreaNsxInline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaNsxProvider) DeepCopyInto(out *AntreaNsxProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaNsxProvider.
func (in *AntreaNsxProvider) DeepCopy() *AntreaNsxProvider {
	if in == nil {
		return nil
	}
	out := new(AntreaNsxProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaProxy) DeepCopyInto(out *AntreaProxy) {
	*out = *in
	if in.NodePortAddresses != nil {
		in, out := &in.NodePortAddresses, &out.NodePortAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SkipServices != nil {
		in, out := &in.SkipServices, &out.SkipServices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaProxy.
func (in *AntreaProxy) DeepCopy() *AntreaProxy {
	if in == nil {
		return nil
	}
	out := new(AntreaProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AntreaProxyNodePortAddress) DeepCopyInto(out *AntreaProxyNodePortAddress) {
	{
		in := &in
		*out = make(AntreaProxyNodePortAddress, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaProxyNodePortAddress.
func (in AntreaProxyNodePortAddress) DeepCopy() AntreaProxyNodePortAddress {
	if in == nil {
		return nil
	}
	out := new(AntreaProxyNodePortAddress)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AntreaWireGuard) DeepCopyInto(out *AntreaWireGuard) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AntreaWireGuard.
func (in *AntreaWireGuard) DeepCopy() *AntreaWireGuard {
	if in == nil {
		return nil
	}
	out := new(AntreaWireGuard)
	in.DeepCopyInto(out)
	return out
}
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Password.DeepCopyInto(&out.Password)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Datasource) DeepCopyInto(out *Datasource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datasource.
func (in *Datasource) DeepCopy() *Datasource {
	if in == nil {
		return nil
	}
	out := new(Datasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Datasource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceList) DeepCopyInto(out *DatasourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Datasource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceList.
func (in *DatasourceList) DeepCopy() *DatasourceList {
	if in == nil {
		return nil
	}
	out := new(DatasourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceSpec) DeepCopyInto(out *DatasourceSpec) {
	*out = *in
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusDatasource)
		(*in).DeepCopyInto(*out)
	}
	if in.Stackdriver != nil {
		in, out := &in.Stackdriver, &out.Stackdriver
		*out = new(StackdriverDatasource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceSpec.
func (in *DatasourceSpec) DeepCopy() *DatasourceSpec {
	if in == nil {
		return nil
	}
	out := new(DatasourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceStatus) DeepCopyInto(out *DatasourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceStatus.
func (in *DatasourceStatus) DeepCopy() *DatasourceStatus {
	if in == nil {
		return nil
	}
	out := new(DatasourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusDatasource) DeepCopyInto(out *PrometheusDatasource) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.ForwardOauth != nil {
		in, out := &in.ForwardOauth, &out.ForwardOauth
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCredentials != nil {
		in, out := &in.ForwardCredentials, &out.ForwardCredentials
		*out = new(bool)
		**out = **in
	}
	if in.SkipTLSVerify != nil {
		in, out := &in.SkipTLSVerify, &out.SkipTLSVerify
		*out = new(bool)
		**out = **in
	}
	if in.ForwardCookies != nil {
		in, out := &in.ForwardCookies, &out.ForwardCookies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = new(ValueOrRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusDatasource.
func (in *PrometheusDatasource) DeepCopy() *PrometheusDatasource {
	if in == nil {
		return nil
	}
	out := new(PrometheusDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackdriverDatasource) DeepCopyInto(out *StackdriverDatasource) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	if in.JWTAuthentication != nil {
		in, out := &in.JWTAuthentication, &out.JWTAuthentication
		*out = new(ValueOrRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackdriverDatasource.
func (in *StackdriverDatasource) DeepCopy() *StackdriverDatasource {
	if in == nil {
		return nil
	}
	out := new(StackdriverDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueOrRef) DeepCopyInto(out *ValueOrRef) {
	*out = *in
	if in.ValueRef != nil {
		in, out := &in.ValueRef, &out.ValueRef
		*out = new(ValueRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueOrRef.
func (in *ValueOrRef) DeepCopy() *ValueOrRef {
	if in == nil {
		return nil
	}
	out := new(ValueOrRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueRef) DeepCopyInto(out *ValueRef) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueRef.
func (in *ValueRef) DeepCopy() *ValueRef {
	if in == nil {
		return nil
	}
	out := new(ValueRef)
	in.DeepCopyInto(out)
	return out
}

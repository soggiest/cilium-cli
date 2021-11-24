// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	"context"
	"time"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumNetworkPoliciesGetter has a method to return a CiliumNetworkPolicyInterface.
// A group's client should implement this interface.
type CiliumNetworkPoliciesGetter interface {
	CiliumNetworkPolicies(namespace string) CiliumNetworkPolicyInterface
}

// CiliumNetworkPolicyInterface has methods to work with CiliumNetworkPolicy resources.
type CiliumNetworkPolicyInterface interface {
	Create(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.CreateOptions) (*v2.CiliumNetworkPolicy, error)
	Update(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.UpdateOptions) (*v2.CiliumNetworkPolicy, error)
	UpdateStatus(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.UpdateOptions) (*v2.CiliumNetworkPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.CiliumNetworkPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.CiliumNetworkPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumNetworkPolicy, err error)
	CiliumNetworkPolicyExpansion
}

// ciliumNetworkPolicies implements CiliumNetworkPolicyInterface
type ciliumNetworkPolicies struct {
	client rest.Interface
	ns     string
}

// newCiliumNetworkPolicies returns a CiliumNetworkPolicies
func newCiliumNetworkPolicies(c *CiliumV2Client, namespace string) *ciliumNetworkPolicies {
	return &ciliumNetworkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ciliumNetworkPolicy, and returns the corresponding ciliumNetworkPolicy object, and an error if there is any.
func (c *ciliumNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumNetworkPolicy, err error) {
	result = &v2.CiliumNetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumNetworkPolicies that match those selectors.
func (c *ciliumNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumNetworkPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.CiliumNetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumNetworkPolicies.
func (c *ciliumNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumNetworkPolicy and creates it.  Returns the server's representation of the ciliumNetworkPolicy, and an error, if there is any.
func (c *ciliumNetworkPolicies) Create(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.CreateOptions) (result *v2.CiliumNetworkPolicy, err error) {
	result = &v2.CiliumNetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumNetworkPolicy and updates it. Returns the server's representation of the ciliumNetworkPolicy, and an error, if there is any.
func (c *ciliumNetworkPolicies) Update(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.UpdateOptions) (result *v2.CiliumNetworkPolicy, err error) {
	result = &v2.CiliumNetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		Name(ciliumNetworkPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *ciliumNetworkPolicies) UpdateStatus(ctx context.Context, ciliumNetworkPolicy *v2.CiliumNetworkPolicy, opts v1.UpdateOptions) (result *v2.CiliumNetworkPolicy, err error) {
	result = &v2.CiliumNetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		Name(ciliumNetworkPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *ciliumNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumNetworkPolicy.
func (c *ciliumNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumNetworkPolicy, err error) {
	result = &v2.CiliumNetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ciliumnetworkpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

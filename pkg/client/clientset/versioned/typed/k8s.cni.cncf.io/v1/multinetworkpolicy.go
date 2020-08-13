/*
Copyright 2020 The Kubernetes Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/k8snetworkplumbingwg/multi-networkpolicy/pkg/apis/k8s.cni.cncf.io/v1"
	scheme "github.com/k8snetworkplumbingwg/multi-networkpolicy/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MultiNetworkPoliciesGetter has a method to return a MultiNetworkPolicyInterface.
// A group's client should implement this interface.
type MultiNetworkPoliciesGetter interface {
	MultiNetworkPolicies(namespace string) MultiNetworkPolicyInterface
}

// MultiNetworkPolicyInterface has methods to work with MultiNetworkPolicy resources.
type MultiNetworkPolicyInterface interface {
	Create(ctx context.Context, multiNetworkPolicy *v1.MultiNetworkPolicy, opts metav1.CreateOptions) (*v1.MultiNetworkPolicy, error)
	Update(ctx context.Context, multiNetworkPolicy *v1.MultiNetworkPolicy, opts metav1.UpdateOptions) (*v1.MultiNetworkPolicy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.MultiNetworkPolicy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MultiNetworkPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MultiNetworkPolicy, err error)
	MultiNetworkPolicyExpansion
}

// multiNetworkPolicies implements MultiNetworkPolicyInterface
type multiNetworkPolicies struct {
	client rest.Interface
	ns     string
}

// newMultiNetworkPolicies returns a MultiNetworkPolicies
func newMultiNetworkPolicies(c *K8sCniCncfIoV1Client, namespace string) *multiNetworkPolicies {
	return &multiNetworkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the multiNetworkPolicy, and returns the corresponding multiNetworkPolicy object, and an error if there is any.
func (c *multiNetworkPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MultiNetworkPolicy, err error) {
	result = &v1.MultiNetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multi-networkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MultiNetworkPolicies that match those selectors.
func (c *multiNetworkPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MultiNetworkPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MultiNetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multi-networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested multiNetworkPolicies.
func (c *multiNetworkPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("multi-networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a multiNetworkPolicy and creates it.  Returns the server's representation of the multiNetworkPolicy, and an error, if there is any.
func (c *multiNetworkPolicies) Create(ctx context.Context, multiNetworkPolicy *v1.MultiNetworkPolicy, opts metav1.CreateOptions) (result *v1.MultiNetworkPolicy, err error) {
	result = &v1.MultiNetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("multi-networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a multiNetworkPolicy and updates it. Returns the server's representation of the multiNetworkPolicy, and an error, if there is any.
func (c *multiNetworkPolicies) Update(ctx context.Context, multiNetworkPolicy *v1.MultiNetworkPolicy, opts metav1.UpdateOptions) (result *v1.MultiNetworkPolicy, err error) {
	result = &v1.MultiNetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multi-networkpolicies").
		Name(multiNetworkPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the multiNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *multiNetworkPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multi-networkpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *multiNetworkPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multi-networkpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched multiNetworkPolicy.
func (c *multiNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MultiNetworkPolicy, err error) {
	result = &v1.MultiNetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("multi-networkpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

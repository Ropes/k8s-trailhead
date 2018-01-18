/*
Copyright 2018 The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/ropes/k8s-trailhead/pkg/apis/trailhead/v1"
	scheme "github.com/ropes/k8s-trailhead/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SimpleSpecsGetter has a method to return a SimpleSpecInterface.
// A group's client should implement this interface.
type SimpleSpecsGetter interface {
	SimpleSpecs(namespace string) SimpleSpecInterface
}

// SimpleSpecInterface has methods to work with SimpleSpec resources.
type SimpleSpecInterface interface {
	Create(*v1.SimpleSpec) (*v1.SimpleSpec, error)
	Update(*v1.SimpleSpec) (*v1.SimpleSpec, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.SimpleSpec, error)
	List(opts meta_v1.ListOptions) (*v1.SimpleSpecList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SimpleSpec, err error)
	SimpleSpecExpansion
}

// simpleSpecs implements SimpleSpecInterface
type simpleSpecs struct {
	client rest.Interface
	ns     string
}

// newSimpleSpecs returns a SimpleSpecs
func newSimpleSpecs(c *TrailheadV1Client, namespace string) *simpleSpecs {
	return &simpleSpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the simpleSpec, and returns the corresponding simpleSpec object, and an error if there is any.
func (c *simpleSpecs) Get(name string, options meta_v1.GetOptions) (result *v1.SimpleSpec, err error) {
	result = &v1.SimpleSpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("simplespecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SimpleSpecs that match those selectors.
func (c *simpleSpecs) List(opts meta_v1.ListOptions) (result *v1.SimpleSpecList, err error) {
	result = &v1.SimpleSpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("simplespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested simpleSpecs.
func (c *simpleSpecs) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("simplespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a simpleSpec and creates it.  Returns the server's representation of the simpleSpec, and an error, if there is any.
func (c *simpleSpecs) Create(simpleSpec *v1.SimpleSpec) (result *v1.SimpleSpec, err error) {
	result = &v1.SimpleSpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("simplespecs").
		Body(simpleSpec).
		Do().
		Into(result)
	return
}

// Update takes the representation of a simpleSpec and updates it. Returns the server's representation of the simpleSpec, and an error, if there is any.
func (c *simpleSpecs) Update(simpleSpec *v1.SimpleSpec) (result *v1.SimpleSpec, err error) {
	result = &v1.SimpleSpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("simplespecs").
		Name(simpleSpec.Name).
		Body(simpleSpec).
		Do().
		Into(result)
	return
}

// Delete takes name of the simpleSpec and deletes it. Returns an error if one occurs.
func (c *simpleSpecs) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("simplespecs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *simpleSpecs) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("simplespecs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched simpleSpec.
func (c *simpleSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SimpleSpec, err error) {
	result = &v1.SimpleSpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("simplespecs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

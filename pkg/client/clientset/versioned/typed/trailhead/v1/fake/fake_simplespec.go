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

package fake

import (
	trailhead_v1 "github.com/ropes/k8s-trailhead/pkg/apis/trailhead/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSimpleSpecs implements SimpleSpecInterface
type FakeSimpleSpecs struct {
	Fake *FakeTrailheadV1
	ns   string
}

var simplespecsResource = schema.GroupVersionResource{Group: "trailhead", Version: "v1", Resource: "simplespecs"}

var simplespecsKind = schema.GroupVersionKind{Group: "trailhead", Version: "v1", Kind: "SimpleSpec"}

// Get takes name of the simpleSpec, and returns the corresponding simpleSpec object, and an error if there is any.
func (c *FakeSimpleSpecs) Get(name string, options v1.GetOptions) (result *trailhead_v1.SimpleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(simplespecsResource, c.ns, name), &trailhead_v1.SimpleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trailhead_v1.SimpleSpec), err
}

// List takes label and field selectors, and returns the list of SimpleSpecs that match those selectors.
func (c *FakeSimpleSpecs) List(opts v1.ListOptions) (result *trailhead_v1.SimpleSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(simplespecsResource, simplespecsKind, c.ns, opts), &trailhead_v1.SimpleSpecList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &trailhead_v1.SimpleSpecList{}
	for _, item := range obj.(*trailhead_v1.SimpleSpecList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested simpleSpecs.
func (c *FakeSimpleSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(simplespecsResource, c.ns, opts))

}

// Create takes the representation of a simpleSpec and creates it.  Returns the server's representation of the simpleSpec, and an error, if there is any.
func (c *FakeSimpleSpecs) Create(simpleSpec *trailhead_v1.SimpleSpec) (result *trailhead_v1.SimpleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(simplespecsResource, c.ns, simpleSpec), &trailhead_v1.SimpleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trailhead_v1.SimpleSpec), err
}

// Update takes the representation of a simpleSpec and updates it. Returns the server's representation of the simpleSpec, and an error, if there is any.
func (c *FakeSimpleSpecs) Update(simpleSpec *trailhead_v1.SimpleSpec) (result *trailhead_v1.SimpleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(simplespecsResource, c.ns, simpleSpec), &trailhead_v1.SimpleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trailhead_v1.SimpleSpec), err
}

// Delete takes name of the simpleSpec and deletes it. Returns an error if one occurs.
func (c *FakeSimpleSpecs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(simplespecsResource, c.ns, name), &trailhead_v1.SimpleSpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSimpleSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(simplespecsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &trailhead_v1.SimpleSpecList{})
	return err
}

// Patch applies the patch and returns the patched simpleSpec.
func (c *FakeSimpleSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *trailhead_v1.SimpleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(simplespecsResource, c.ns, name, data, subresources...), &trailhead_v1.SimpleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*trailhead_v1.SimpleSpec), err
}

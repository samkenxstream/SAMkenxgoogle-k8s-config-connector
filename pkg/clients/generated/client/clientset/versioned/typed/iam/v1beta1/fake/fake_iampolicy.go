// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIAMPolicies implements IAMPolicyInterface
type FakeIAMPolicies struct {
	Fake *FakeIamV1beta1
	ns   string
}

var iampoliciesResource = schema.GroupVersionResource{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Resource: "iampolicies"}

var iampoliciesKind = schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPolicy"}

// Get takes name of the iAMPolicy, and returns the corresponding iAMPolicy object, and an error if there is any.
func (c *FakeIAMPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IAMPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iampoliciesResource, c.ns, name), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}

// List takes label and field selectors, and returns the list of IAMPolicies that match those selectors.
func (c *FakeIAMPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IAMPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iampoliciesResource, iampoliciesKind, c.ns, opts), &v1beta1.IAMPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IAMPolicyList{ListMeta: obj.(*v1beta1.IAMPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.IAMPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iAMPolicies.
func (c *FakeIAMPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iampoliciesResource, c.ns, opts))

}

// Create takes the representation of a iAMPolicy and creates it.  Returns the server's representation of the iAMPolicy, and an error, if there is any.
func (c *FakeIAMPolicies) Create(ctx context.Context, iAMPolicy *v1beta1.IAMPolicy, opts v1.CreateOptions) (result *v1beta1.IAMPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iampoliciesResource, c.ns, iAMPolicy), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}

// Update takes the representation of a iAMPolicy and updates it. Returns the server's representation of the iAMPolicy, and an error, if there is any.
func (c *FakeIAMPolicies) Update(ctx context.Context, iAMPolicy *v1beta1.IAMPolicy, opts v1.UpdateOptions) (result *v1beta1.IAMPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iampoliciesResource, c.ns, iAMPolicy), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIAMPolicies) UpdateStatus(ctx context.Context, iAMPolicy *v1beta1.IAMPolicy, opts v1.UpdateOptions) (*v1beta1.IAMPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iampoliciesResource, "status", c.ns, iAMPolicy), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}

// Delete takes name of the iAMPolicy and deletes it. Returns an error if one occurs.
func (c *FakeIAMPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(iampoliciesResource, c.ns, name, opts), &v1beta1.IAMPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIAMPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iampoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IAMPolicyList{})
	return err
}

// Patch applies the patch and returns the patched iAMPolicy.
func (c *FakeIAMPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IAMPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iampoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.IAMPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMPolicy), err
}

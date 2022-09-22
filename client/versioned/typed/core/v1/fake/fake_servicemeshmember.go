// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	corev1 "maistra.io/api/core/v1"
)

// FakeServiceMeshMembers implements ServiceMeshMemberInterface
type FakeServiceMeshMembers struct {
	Fake *FakeCoreV1
	ns   string
}

var servicemeshmembersResource = schema.GroupVersionResource{Group: "maistra.io", Version: "v1", Resource: "servicemeshmembers"}

var servicemeshmembersKind = schema.GroupVersionKind{Group: "maistra.io", Version: "v1", Kind: "ServiceMeshMember"}

// Get takes name of the serviceMeshMember, and returns the corresponding serviceMeshMember object, and an error if there is any.
func (c *FakeServiceMeshMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.ServiceMeshMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicemeshmembersResource, c.ns, name), &corev1.ServiceMeshMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshMember), err
}

// List takes label and field selectors, and returns the list of ServiceMeshMembers that match those selectors.
func (c *FakeServiceMeshMembers) List(ctx context.Context, opts v1.ListOptions) (result *corev1.ServiceMeshMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicemeshmembersResource, servicemeshmembersKind, c.ns, opts), &corev1.ServiceMeshMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ServiceMeshMemberList{ListMeta: obj.(*corev1.ServiceMeshMemberList).ListMeta}
	for _, item := range obj.(*corev1.ServiceMeshMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceMeshMembers.
func (c *FakeServiceMeshMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicemeshmembersResource, c.ns, opts))

}

// Create takes the representation of a serviceMeshMember and creates it.  Returns the server's representation of the serviceMeshMember, and an error, if there is any.
func (c *FakeServiceMeshMembers) Create(ctx context.Context, serviceMeshMember *corev1.ServiceMeshMember, opts v1.CreateOptions) (result *corev1.ServiceMeshMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicemeshmembersResource, c.ns, serviceMeshMember), &corev1.ServiceMeshMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshMember), err
}

// Update takes the representation of a serviceMeshMember and updates it. Returns the server's representation of the serviceMeshMember, and an error, if there is any.
func (c *FakeServiceMeshMembers) Update(ctx context.Context, serviceMeshMember *corev1.ServiceMeshMember, opts v1.UpdateOptions) (result *corev1.ServiceMeshMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicemeshmembersResource, c.ns, serviceMeshMember), &corev1.ServiceMeshMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceMeshMembers) UpdateStatus(ctx context.Context, serviceMeshMember *corev1.ServiceMeshMember, opts v1.UpdateOptions) (*corev1.ServiceMeshMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicemeshmembersResource, "status", c.ns, serviceMeshMember), &corev1.ServiceMeshMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshMember), err
}

// Delete takes name of the serviceMeshMember and deletes it. Returns an error if one occurs.
func (c *FakeServiceMeshMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(servicemeshmembersResource, c.ns, name, opts), &corev1.ServiceMeshMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceMeshMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicemeshmembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.ServiceMeshMemberList{})
	return err
}

// Patch applies the patch and returns the patched serviceMeshMember.
func (c *FakeServiceMeshMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.ServiceMeshMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicemeshmembersResource, c.ns, name, pt, data, subresources...), &corev1.ServiceMeshMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceMeshMember), err
}

/*

Don't alter this file, it was generated.

*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "code.cloudfoundry.org/quarks-secret/pkg/kube/apis/quarkssecret/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeQuarksSecrets implements QuarksSecretInterface
type FakeQuarksSecrets struct {
	Fake *FakeQuarkssecretV1alpha1
	ns   string
}

var quarkssecretsResource = schema.GroupVersionResource{Group: "quarkssecret", Version: "v1alpha1", Resource: "quarkssecrets"}

var quarkssecretsKind = schema.GroupVersionKind{Group: "quarkssecret", Version: "v1alpha1", Kind: "QuarksSecret"}

// Get takes name of the quarksSecret, and returns the corresponding quarksSecret object, and an error if there is any.
func (c *FakeQuarksSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.QuarksSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(quarkssecretsResource, c.ns, name), &v1alpha1.QuarksSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuarksSecret), err
}

// List takes label and field selectors, and returns the list of QuarksSecrets that match those selectors.
func (c *FakeQuarksSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.QuarksSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(quarkssecretsResource, quarkssecretsKind, c.ns, opts), &v1alpha1.QuarksSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.QuarksSecretList{ListMeta: obj.(*v1alpha1.QuarksSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.QuarksSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested quarksSecrets.
func (c *FakeQuarksSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(quarkssecretsResource, c.ns, opts))

}

// Create takes the representation of a quarksSecret and creates it.  Returns the server's representation of the quarksSecret, and an error, if there is any.
func (c *FakeQuarksSecrets) Create(ctx context.Context, quarksSecret *v1alpha1.QuarksSecret, opts v1.CreateOptions) (result *v1alpha1.QuarksSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(quarkssecretsResource, c.ns, quarksSecret), &v1alpha1.QuarksSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuarksSecret), err
}

// Update takes the representation of a quarksSecret and updates it. Returns the server's representation of the quarksSecret, and an error, if there is any.
func (c *FakeQuarksSecrets) Update(ctx context.Context, quarksSecret *v1alpha1.QuarksSecret, opts v1.UpdateOptions) (result *v1alpha1.QuarksSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(quarkssecretsResource, c.ns, quarksSecret), &v1alpha1.QuarksSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuarksSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeQuarksSecrets) UpdateStatus(ctx context.Context, quarksSecret *v1alpha1.QuarksSecret, opts v1.UpdateOptions) (*v1alpha1.QuarksSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(quarkssecretsResource, "status", c.ns, quarksSecret), &v1alpha1.QuarksSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuarksSecret), err
}

// Delete takes name of the quarksSecret and deletes it. Returns an error if one occurs.
func (c *FakeQuarksSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(quarkssecretsResource, c.ns, name), &v1alpha1.QuarksSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQuarksSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(quarkssecretsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.QuarksSecretList{})
	return err
}

// Patch applies the patch and returns the patched quarksSecret.
func (c *FakeQuarksSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.QuarksSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(quarkssecretsResource, c.ns, name, pt, data, subresources...), &v1alpha1.QuarksSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QuarksSecret), err
}

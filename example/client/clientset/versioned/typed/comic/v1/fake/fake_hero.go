// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/slok/kube-code-generator/example/apis/comic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHeros implements HeroInterface
type FakeHeros struct {
	Fake *FakeComicV1
}

var herosResource = v1.SchemeGroupVersion.WithResource("heros")

var herosKind = v1.SchemeGroupVersion.WithKind("Hero")

// Get takes name of the hero, and returns the corresponding hero object, and an error if there is any.
func (c *FakeHeros) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Hero, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(herosResource, name), &v1.Hero{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Hero), err
}

// List takes label and field selectors, and returns the list of Heros that match those selectors.
func (c *FakeHeros) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HeroList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(herosResource, herosKind, opts), &v1.HeroList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.HeroList{ListMeta: obj.(*v1.HeroList).ListMeta}
	for _, item := range obj.(*v1.HeroList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested heros.
func (c *FakeHeros) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(herosResource, opts))
}

// Create takes the representation of a hero and creates it.  Returns the server's representation of the hero, and an error, if there is any.
func (c *FakeHeros) Create(ctx context.Context, hero *v1.Hero, opts metav1.CreateOptions) (result *v1.Hero, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(herosResource, hero), &v1.Hero{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Hero), err
}

// Update takes the representation of a hero and updates it. Returns the server's representation of the hero, and an error, if there is any.
func (c *FakeHeros) Update(ctx context.Context, hero *v1.Hero, opts metav1.UpdateOptions) (result *v1.Hero, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(herosResource, hero), &v1.Hero{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Hero), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHeros) UpdateStatus(ctx context.Context, hero *v1.Hero, opts metav1.UpdateOptions) (*v1.Hero, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(herosResource, "status", hero), &v1.Hero{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Hero), err
}

// Delete takes name of the hero and deletes it. Returns an error if one occurs.
func (c *FakeHeros) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(herosResource, name, opts), &v1.Hero{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHeros) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(herosResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.HeroList{})
	return err
}

// Patch applies the patch and returns the patched hero.
func (c *FakeHeros) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Hero, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(herosResource, name, pt, data, subresources...), &v1.Hero{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Hero), err
}

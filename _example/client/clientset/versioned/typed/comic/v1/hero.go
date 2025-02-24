// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	comicv1 "github.com/slok/kube-code-generator/example/apis/comic/v1"
	applyconfigurationcomicv1 "github.com/slok/kube-code-generator/example/client/applyconfiguration/comic/v1"
	scheme "github.com/slok/kube-code-generator/example/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// HerosGetter has a method to return a HeroInterface.
// A group's client should implement this interface.
type HerosGetter interface {
	Heros() HeroInterface
}

// HeroInterface has methods to work with Hero resources.
type HeroInterface interface {
	Create(ctx context.Context, hero *comicv1.Hero, opts metav1.CreateOptions) (*comicv1.Hero, error)
	Update(ctx context.Context, hero *comicv1.Hero, opts metav1.UpdateOptions) (*comicv1.Hero, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, hero *comicv1.Hero, opts metav1.UpdateOptions) (*comicv1.Hero, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*comicv1.Hero, error)
	List(ctx context.Context, opts metav1.ListOptions) (*comicv1.HeroList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *comicv1.Hero, err error)
	Apply(ctx context.Context, hero *applyconfigurationcomicv1.HeroApplyConfiguration, opts metav1.ApplyOptions) (result *comicv1.Hero, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, hero *applyconfigurationcomicv1.HeroApplyConfiguration, opts metav1.ApplyOptions) (result *comicv1.Hero, err error)
	HeroExpansion
}

// heros implements HeroInterface
type heros struct {
	*gentype.ClientWithListAndApply[*comicv1.Hero, *comicv1.HeroList, *applyconfigurationcomicv1.HeroApplyConfiguration]
}

// newHeros returns a Heros
func newHeros(c *ComicV1Client) *heros {
	return &heros{
		gentype.NewClientWithListAndApply[*comicv1.Hero, *comicv1.HeroList, *applyconfigurationcomicv1.HeroApplyConfiguration](
			"heros",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *comicv1.Hero { return &comicv1.Hero{} },
			func() *comicv1.HeroList { return &comicv1.HeroList{} },
		),
	}
}

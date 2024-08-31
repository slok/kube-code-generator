// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/slok/kube-code-generator/example/apis/comic/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// HeroLister helps list Heros.
// All objects returned here must be treated as read-only.
type HeroLister interface {
	// List lists all Heros in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Hero, err error)
	// Get retrieves the Hero from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Hero, error)
	HeroListerExpansion
}

// heroLister implements the HeroLister interface.
type heroLister struct {
	listers.ResourceIndexer[*v1.Hero]
}

// NewHeroLister returns a new HeroLister.
func NewHeroLister(indexer cache.Indexer) HeroLister {
	return &heroLister{listers.New[*v1.Hero](indexer, v1.Resource("hero"))}
}

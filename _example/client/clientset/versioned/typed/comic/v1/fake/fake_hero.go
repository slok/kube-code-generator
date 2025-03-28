// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/slok/kube-code-generator/example/apis/comic/v1"
	comicv1 "github.com/slok/kube-code-generator/example/client/applyconfiguration/comic/v1"
	typedcomicv1 "github.com/slok/kube-code-generator/example/client/clientset/versioned/typed/comic/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeHeros implements HeroInterface
type fakeHeros struct {
	*gentype.FakeClientWithListAndApply[*v1.Hero, *v1.HeroList, *comicv1.HeroApplyConfiguration]
	Fake *FakeComicV1
}

func newFakeHeros(fake *FakeComicV1) typedcomicv1.HeroInterface {
	return &fakeHeros{
		gentype.NewFakeClientWithListAndApply[*v1.Hero, *v1.HeroList, *comicv1.HeroApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("heros"),
			v1.SchemeGroupVersion.WithKind("Hero"),
			func() *v1.Hero { return &v1.Hero{} },
			func() *v1.HeroList { return &v1.HeroList{} },
			func(dst, src *v1.HeroList) { dst.ListMeta = src.ListMeta },
			func(list *v1.HeroList) []*v1.Hero { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.HeroList, items []*v1.Hero) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

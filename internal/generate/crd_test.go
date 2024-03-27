package generate_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/slok/kube-code-generator/internal/generate"
	"github.com/slok/kube-code-generator/internal/generate/generatemock"
	"github.com/slok/kube-code-generator/internal/log"
)

func TestCRDGenerator(t *testing.T) {
	tests := map[string]struct {
		exec   func(g *generate.CRDGenerator)
		expCmd string
	}{
		"Without options.": {
			exec:   func(g *generate.CRDGenerator) { _ = g.Run(context.TODO()) },
			expCmd: `controller-gen paths="" output:dir="" crd:crdVersions=v1`,
		},

		"Regular options.": {
			exec: func(g *generate.CRDGenerator) {
				_ = g.WithAPIsPath("./apis").
					WithOutputDir("./out").
					WithAllowDangerousTypes().
					WithIgnoreDescription().
					WithIgnoreUnexportedFields().
					Run(context.TODO())
			},
			expCmd: `controller-gen paths="./apis/..." output:dir="./out" crd:crdVersions=v1,allowDangerousTypes=true,maxDescLen=0,ignoreUnexportedFields=true`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			m := generatemock.NewBashExecutor(t)
			m.On("BashExec", mock.Anything, test.expCmd).Once().Return("", nil)

			g := generate.NewCRDGenerator(log.Noop, "", m)
			test.exec(g)

			m.AssertExpectations(t)
		})
	}
}

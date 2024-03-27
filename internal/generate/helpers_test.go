package generate_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/slok/kube-code-generator/internal/generate"
	"github.com/slok/kube-code-generator/internal/generate/generatemock"
	"github.com/slok/kube-code-generator/internal/log"
)

func TestHelpersGenerator(t *testing.T) {
	tests := map[string]struct {
		exec   func(g *generate.HelpersGenerator)
		expCmd string
	}{
		"Without options.": {
			exec:   func(g *generate.HelpersGenerator) { _ = g.Run(context.TODO()) },
			expCmd: `source kube_codegen.sh ; kube::codegen::gen_helpers  `,
		},

		"Regular options.": {
			exec: func(g *generate.HelpersGenerator) {
				_ = g.WithAPIsPath("./apis").
					WithBoilerplate("./boilerplate.txt").
					Run(context.TODO())
			},
			expCmd: `source kube_codegen.sh ; kube::codegen::gen_helpers --boilerplate ./boilerplate.txt ./apis`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			m := generatemock.NewBashExecutor(t)
			m.On("BashExec", mock.Anything, test.expCmd).Once().Return("", nil)

			g := generate.NewHelpersGenerator(log.Noop, "", m)
			test.exec(g)

			m.AssertExpectations(t)
		})
	}
}

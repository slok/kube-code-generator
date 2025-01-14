package generate_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/slok/kube-code-generator/internal/generate"
	"github.com/slok/kube-code-generator/internal/generate/generatemock"
	"github.com/slok/kube-code-generator/internal/log"
)

func TestClientGenerator(t *testing.T) {
	tests := map[string]struct {
		exec   func(g *generate.ClientGenerator)
		expCmd string
	}{
		"Without options.": {
			exec:   func(g *generate.ClientGenerator) { _ = g.Run(context.TODO()) },
			expCmd: `source kube_codegen.sh ; kube::codegen::gen_client  `,
		},

		"Regular options.": {
			exec: func(g *generate.ClientGenerator) {
				_ = g.WithAPIsPath("./apis").
					WithBoilerplate("./boilerplate.txt").
					WithOutputDir("./out").
					WithOutputPkg("my-pkg").
					WithApplyConfig().
					WithWatch().
					Run(context.TODO())
			},
			expCmd: `source kube_codegen.sh ; kube::codegen::gen_client --boilerplate ./boilerplate.txt --output-dir ./out --output-pkg my-pkg --with-applyconfig --with-watch ./apis`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			m := generatemock.NewBashExecutor(t)
			m.On("BashExec", mock.Anything, test.expCmd).Once().Return("", nil)

			g := generate.NewClientGenerator(log.Noop, "", m)
			test.exec(g)

			m.AssertExpectations(t)
		})
	}
}

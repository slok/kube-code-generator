package gomod_test

import (
	"testing"

	"github.com/slok/kube-code-generator/internal/util/gomod"
	"github.com/stretchr/testify/assert"
)

func TestGetGoModule(t *testing.T) {
	tests := map[string]struct {
		goDotMod      string
		expectedGoMod string
		expErr        bool
	}{
		"If not content, it should fail.": {
			goDotMod: "",
			expErr:   true,
		},

		"If there is content, but no go module, it should fail.": {
			goDotMod: `
go 1.22.5

require (
	github.com/alecthomas/kingpin/v2 v2.4.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.2
)

`,
			expErr: true,
		},

		"If there is content, and go module, it should return the module.": {
			goDotMod: `
module github.com/slok/kube-code-generator

go 1.22.5

require (
	github.com/alecthomas/kingpin/v2 v2.4.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)`,
			expectedGoMod: "github.com/slok/kube-code-generator",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			gotGoMod, err := gomod.GetGoModule(test.goDotMod)

			if test.expErr {
				assert.Error(err)
			} else if assert.NoError(err) {
				assert.Equal(test.expectedGoMod, gotGoMod)
			}
		})
	}
}

func TestGetGoPackageFromDir(t *testing.T) {
	tests := map[string]struct {
		goMod       string
		pkgDir      string
		expectedPkg string
	}{
		"Having a go module and a package dir, it should return the Go package.": {
			goMod:       "github.com/slok/kube-code-generator/example",
			pkgDir:      "./something/gen/otherthing",
			expectedPkg: "github.com/slok/kube-code-generator/example/something/gen/otherthing",
		},

		"Having a go module and a package dir, it should return the Go package (not relative dir prefix).": {
			goMod:       "github.com/slok/kube-code-generator/example",
			pkgDir:      "something/gen/otherthing",
			expectedPkg: "github.com/slok/kube-code-generator/example/something/gen/otherthing",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			gotGoPkg := gomod.GetGoPackageFromDir(test.goMod, test.pkgDir)
			assert.Equal(test.expectedPkg, gotGoPkg)

		})
	}
}

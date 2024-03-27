package gomod

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	goModuleRegexp = regexp.MustCompile(`(?m)^module ([^\s]+)$`)
)

// GetGoModule will return the go module declaration from a go mod file content.
func GetGoModule(goModFileContent string) (string, error) {
	match := goModuleRegexp.FindAllStringSubmatch(goModFileContent, 1)
	if len(match) < 1 || len(match[0]) < 2 {
		return "", fmt.Errorf(`could not find module declaration on "go.mod"`)
	}
	packageName := match[0][1]

	return packageName, nil
}

// GetImportPackageFromDir will return the go package based on a go project module and a relative directory.
func GetGoPackageFromDir(goModule, relativeDir string) string {
	pkg := strings.TrimSuffix(relativeDir, "/") + "/" // Ensure slash.
	pkg = filepath.Dir(pkg)
	if pkg != "." && pkg != "" {
		return goModule + "/" + pkg
	}

	return pkg

}

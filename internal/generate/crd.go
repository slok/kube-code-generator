package generate

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/slok/kube-code-generator/internal/log"
)

type CRDGenerator struct {
	controllerGenBin string
	apisPath         string
	outPath          string
	crdOptions       []string
	exec             BashExecutor
	logger           log.Logger
}

const strSeparator = string(filepath.Separator)

func NewCRDGenerator(logger log.Logger, controllerGenBin string, exec BashExecutor) *CRDGenerator {
	if exec == nil {
		exec = StdBashExecutor
	}

	if controllerGenBin == "" {
		controllerGenBin = "controller-gen"
	}

	return &CRDGenerator{
		controllerGenBin: controllerGenBin,
		crdOptions: []string{
			"crdVersions=v1", // Only one supported for now.
		},
		exec:   exec,
		logger: logger,
	}
}

func (g *CRDGenerator) WithAllowDangerousTypes() *CRDGenerator {
	g.crdOptions = append(g.crdOptions, "allowDangerousTypes=true")
	return g
}

func (g *CRDGenerator) WithIgnoreUnexportedFields() *CRDGenerator {
	g.crdOptions = append(g.crdOptions, "ignoreUnexportedFields=true")
	return g
}

func (g *CRDGenerator) WithIgnoreDescription() *CRDGenerator {
	g.crdOptions = append(g.crdOptions, "maxDescLen=0")
	return g
}

func (g *CRDGenerator) WithOutputDir(path string) *CRDGenerator {
	path = filepath.Clean(path)
	g.outPath = fmt.Sprintf(".%s%s", strSeparator, path) // We need `./` in front of it.
	return g
}

func (g *CRDGenerator) WithAPIsPath(path string) *CRDGenerator {
	path = filepath.Clean(path)
	g.apisPath = fmt.Sprintf(".%s%s%s...", strSeparator, path, strSeparator) // We need `./` in front of it.
	return g
}

func (g *CRDGenerator) Run(ctx context.Context) error {
	paths := fmt.Sprintf(`paths="%s"`, g.apisPath)
	outputDir := fmt.Sprintf(`output:dir="%s"`, g.outPath)
	crds := fmt.Sprintf("crd:%s", strings.Join(g.crdOptions, ","))
	bashCmd := fmt.Sprintf("%s %s %s %s", g.controllerGenBin, paths, outputDir, crds)

	g.logger.Debugf("Command executed: %s", bashCmd)
	out, err := g.exec.BashExec(ctx, bashCmd)
	if err != nil {
		return fmt.Errorf("error while executing bash script: %w", err)
	}
	g.logger.Debugf("Command output: %s", string(out))
	return nil
}

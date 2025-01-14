package generate

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/slok/kube-code-generator/internal/log"
)

type ClientGenerator struct {
	cmdArgs     []string
	codeGenPath string
	apisPath    string
	exec        BashExecutor
	logger      log.Logger
}

func NewClientGenerator(logger log.Logger, codeGenPath string, exec BashExecutor) *ClientGenerator {
	if exec == nil {
		exec = StdBashExecutor
	}

	return &ClientGenerator{
		codeGenPath: codeGenPath,
		exec:        exec,
		logger:      logger,
	}
}

func (g *ClientGenerator) WithWatch() *ClientGenerator {
	g.cmdArgs = append(g.cmdArgs, `--with-watch`)
	return g
}

func (g *ClientGenerator) WithOutputPkg(pkg string) *ClientGenerator {
	g.cmdArgs = append(g.cmdArgs, `--output-pkg`, pkg)
	return g
}

func (g *ClientGenerator) WithOutputDir(path string) *ClientGenerator {
	g.cmdArgs = append(g.cmdArgs, `--output-dir`, path)
	return g
}

func (g *ClientGenerator) WithBoilerplate(path string) *ClientGenerator {
	g.cmdArgs = append(g.cmdArgs, `--boilerplate`, path)
	return g
}

func (g *ClientGenerator) WithAPIsPath(path string) *ClientGenerator {
	g.apisPath = path
	return g
}

func (g *ClientGenerator) WithApplyConfig() *ClientGenerator {
	g.cmdArgs = append(g.cmdArgs, `--with-applyconfig`)
	return g
}

func (g *ClientGenerator) Run(ctx context.Context) error {
	kubeCodeGenSHPath := filepath.Join(g.codeGenPath, "kube_codegen.sh")
	bashCmd := fmt.Sprintf("source %s ; kube::codegen::gen_client %s %s", kubeCodeGenSHPath, strings.Join(g.cmdArgs, " "), g.apisPath)

	g.logger.Debugf("Command executed: %s", bashCmd)
	out, err := g.exec.BashExec(ctx, bashCmd)
	if err != nil {
		return fmt.Errorf("error while executing bash script: %w", err)
	}
	g.logger.Debugf("Command output: %s", string(out))

	return nil
}

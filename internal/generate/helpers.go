package generate

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/slok/kube-code-generator/internal/log"
)

type HelpersGenerator struct {
	cmdArgs     []string
	codeGenPath string
	apisPath    string
	exec        BashExecutor
	logger      log.Logger
}

func NewHelpersGenerator(logger log.Logger, codeGenPath string, exec BashExecutor) *HelpersGenerator {
	if exec == nil {
		exec = StdBashExecutor
	}

	return &HelpersGenerator{
		codeGenPath: codeGenPath,
		exec:        exec,
		logger:      logger,
	}
}

func (g *HelpersGenerator) WithBoilerplate(path string) *HelpersGenerator {
	g.cmdArgs = append(g.cmdArgs, `--boilerplate`, path)
	return g
}

func (g *HelpersGenerator) WithAPIsPath(path string) *HelpersGenerator {
	g.apisPath = path
	return g
}

func (g *HelpersGenerator) Run(ctx context.Context) error {
	kubeCodeGenSHPath := filepath.Join(g.codeGenPath, "kube_codegen.sh")
	bashCmd := fmt.Sprintf("source %s ; kube::codegen::gen_helpers %s %s", kubeCodeGenSHPath, strings.Join(g.cmdArgs, " "), g.apisPath)

	g.logger.Debugf("Command executed: %s", bashCmd)
	out, err := g.exec.BashExec(ctx, bashCmd)
	if err != nil {
		return fmt.Errorf("error while executing bash script: %w", err)
	}
	g.logger.Debugf("Command output: %s", string(out))

	return nil
}

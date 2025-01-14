package main

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	generate "github.com/slok/kube-code-generator/internal/generate"
	"github.com/slok/kube-code-generator/internal/info"
	"github.com/slok/kube-code-generator/internal/log"
	loglogrus "github.com/slok/kube-code-generator/internal/log/logrus"
	utilgomod "github.com/slok/kube-code-generator/internal/util/gomod"
)

func run(ctx context.Context, args []string, stdout, stderr io.Writer) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Load command flags and arguments.
	cmdCfg, err := NewCmdConfig(args)
	if err != nil {
		return fmt.Errorf("could not load command configuration: %w", err)
	}

	// If not logger disabled use logrus logger.
	logrusLog := logrus.New()
	logrusLog.Out = stderr // By default logger goes to stderr (so it can split stdout prints).
	logrusLogEntry := logrus.NewEntry(logrusLog)
	if cmdCfg.Debug {
		logrusLogEntry.Logger.SetLevel(logrus.DebugLevel)
	}
	logger := loglogrus.NewLogrus(logrusLogEntry).WithValues(log.Kv{
		"version": info.Version,
	})

	logger.Debugf("Debug level is enabled") // Will log only when debug enabled.

	if cmdCfg.GoCodeOutPath == "" && cmdCfg.CRDsOutPath == "" {
		return fmt.Errorf("at least a generated output path is required")
	}

	// Prepare.
	projectRootFS := os.DirFS(".")
	goMod, err := fs.ReadFile(projectRootFS, "go.mod")
	if err != nil {
		return fmt.Errorf(`error while reading "go.mod", you should execute this app from the project root: %w`, err)
	}

	goModule, err := utilgomod.GetGoModule(string(goMod))
	if err != nil {
		return fmt.Errorf("could not get go module: %w", err)
	}
	goCodeGenOutPkg := utilgomod.GetGoPackageFromDir(goModule, cmdCfg.GoCodeOutPath)
	logger.WithValues(log.Kv{"module": goCodeGenOutPkg}).Infof("Go generated code package inferred")

	if filepath.IsAbs(cmdCfg.APIsPath) {
		return fmt.Errorf("APIs path should be relative")
	}

	// Start autogeneration.
	err = generateGoCode(ctx, *cmdCfg, logger, goCodeGenOutPkg)
	if err != nil {
		return fmt.Errorf("could not generate Go code: %w", err)
	}

	err = generateCRDManifests(ctx, *cmdCfg, logger)
	if err != nil {
		return fmt.Errorf("could not generate CRDs: %w", err)
	}

	return nil
}

func generateGoCode(ctx context.Context, cmdCfg CmdConfig, logger log.Logger, genOutPkg string) error {
	if cmdCfg.GoCodeOutPath == "" {
		logger.Infof("Ignoring Go code generation")
		return nil
	}

	if filepath.IsAbs(cmdCfg.GoCodeOutPath) {
		return fmt.Errorf("Go generated code path should be relative")
	}

	err := os.MkdirAll(cmdCfg.GoCodeOutPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create directory for go generated code: %w", err)
	}

	// We will require a boilerplate always.
	boilerplatePath := cmdCfg.BoilerplatePath
	if cmdCfg.BoilerplatePath == "" {
		f, err := os.CreateTemp("", "kube-code-generator-boilerplate-")
		if err != nil {
			return fmt.Errorf("could not create boilerplate empty file")
		}
		f.Close()
		defer os.Remove(f.Name())
		boilerplatePath = f.Name()
	}

	logger.Infof("Generating Go code...")
	gen := generate.NewClientGenerator(logger, cmdCfg.CodeGenPath, generate.StdBashExecutor).
		WithWatch().
		WithBoilerplate(boilerplatePath).
		WithOutputPkg(genOutPkg).
		WithOutputDir(cmdCfg.GoCodeOutPath).
		WithAPIsPath(cmdCfg.APIsPath)

	if cmdCfg.EnableApplyConfigs {
		gen = gen.WithApplyConfig()
	}

	err = gen.Run(ctx)
	if err != nil {
		return fmt.Errorf("could not generate Go clients code: %w", err)
	}
	err = generate.NewHelpersGenerator(logger, cmdCfg.CodeGenPath, generate.StdBashExecutor).
		WithBoilerplate(boilerplatePath).
		WithAPIsPath(cmdCfg.APIsPath).Run(ctx)
	if err != nil {
		return fmt.Errorf("could not generate types Go helper code: %w", err)
	}

	return nil
}

func generateCRDManifests(ctx context.Context, cmdCfg CmdConfig, logger log.Logger) error {
	if cmdCfg.CRDsOutPath == "" {
		logger.Infof("Ignoring CRD manifest generation")
		return nil
	}

	if filepath.IsAbs(cmdCfg.CRDsOutPath) {
		return fmt.Errorf("crd manifests generated path should be relative")
	}

	err := os.MkdirAll(cmdCfg.CRDsOutPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create directory for CRD manifests generated content: %w", err)
	}

	logger.Infof("Generating CRDs...")
	err = generate.NewCRDGenerator(logger, cmdCfg.ControllerGenBin, generate.StdBashExecutor).
		WithAllowDangerousTypes().
		WithOutputDir(cmdCfg.CRDsOutPath).
		WithAPIsPath(cmdCfg.APIsPath).Run(ctx)
	if err != nil {
		return fmt.Errorf("could not generate Go types code: %w", err)
	}

	return nil
}

func main() {
	err := run(context.Background(), os.Args, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

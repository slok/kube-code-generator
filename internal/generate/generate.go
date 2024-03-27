package generate

import (
	"context"
	"fmt"
	"os/exec"
)

type BashExecutor interface {
	BashExec(ctx context.Context, bashCmd string) (string, error)
}

//go:generate mockery --case underscore --output generatemock --outpkg generatemock --name BashExecutor

// StdBashExecutor is an standard bash executor.
var StdBashExecutor = stdBashExecutor(false)

type stdBashExecutor bool

func (stdBashExecutor) BashExec(ctx context.Context, bashCmd string) (string, error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", bashCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(out))
	}

	return string(out), nil
}

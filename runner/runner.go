package runner

import (
	"context"
	"os/exec"
)

type Runner struct {
	executable string
	args       []string
	done       bool

	ctx context.Context
	cmd *exec.Cmd
}

func NewRunner(executable string, args []string) *Runner {
	return &Runner{
		executable: executable,
		args:       args,
		ctx:        context.Background(),
	}
}

func (r *Runner) Run() error {
	// exec.CommandContext(r.ctx, r.executable, r.args...)
	return nil
}

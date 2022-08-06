//go:build !debug
// +build !debug

package pocketbook

import (
	"context"
	"fmt"
	"os/exec"
)

func RefreshScanner(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "/bin/killall", "scanner.app")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}

package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var err error

	if len(cmd) == 0 {
		return 1
	}

	cmds := exec.Command(cmd[0], cmd[1:]...) // #nosec G204
	cmds.Stdin, cmds.Stdout, cmds.Stderr = os.Stdin, os.Stdout, os.Stderr

	for k, v := range env {
		if v.Value == "" {
			err = os.Unsetenv(k)
		} else {
			err = os.Setenv(k, v.Value)
		}

		if err != nil {
			return 1
		}
	}

	if err = cmds.Run(); err != nil {
		return cmds.ProcessState.ExitCode()
	}

	return 0
}

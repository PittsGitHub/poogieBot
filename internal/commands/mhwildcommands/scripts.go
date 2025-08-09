package mhwildcommands

import (
	"os/exec"
)

// RunUpdateScript executes the mhwilds update script and returns combined stdout/stderr output.
func RunUpdateScript(scriptPath string) (string, error) {
	cmd := exec.Command("bash", scriptPath)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

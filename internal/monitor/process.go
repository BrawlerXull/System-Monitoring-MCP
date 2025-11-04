package monitor

import (
	"errors"
	"os"
	"os/exec"

	"github.com/shirou/gopsutil/v3/process"
)

// ListProcessesTool lists all running processes with their PID, name, and CPU usage
func ListProcessesTool() ([]map[string]interface{}, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	for _, p := range procs {
		name, _ := p.Name()
		cpu, _ := p.CPUPercent()
		results = append(results, map[string]interface{}{
			"pid":  p.Pid,
			"name": name,
			"cpu":  cpu,
		})
	}

	return results, nil
}

// KillProcessTool terminates a process by PID
func KillProcessTool(pid int32) (string, error) {
	proc, err := os.FindProcess(int(pid))
	if err != nil {
		return "", err
	}
	if err := proc.Kill(); err != nil {
		return "", err
	}
	return "Process killed successfully", nil
}

// LaunchProcessTool starts a new process given a command and args
func LaunchProcessTool(command string, args []string) (string, error) {
	if command == "" {
		return "", errors.New("command cannot be empty")
	}

	cmd := exec.Command(command, args...)
	if err := cmd.Start(); err != nil {
		return "", err
	}

	return "Process launched successfully", nil
}

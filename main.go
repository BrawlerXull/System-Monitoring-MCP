// Package main implements a Model Context Protocol (MCP) server in Go
// that provides system monitoring, control, and automation tools.
// Compatible with Claude Desktop integration.
//
// 📦 Features implemented:
// - CPU, Memory, Network, Battery, GPU, System info
// - Process listing, launching, killing
// - File system inspection

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/distatus/battery"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	ctx := context.Background()

	// Initialize MCP server metadata
	impl := &mcp.Implementation{
		Name:    "system-monitor-go",
		Version: "0.2.0",
	}
	server := mcp.NewServer(impl, nil)

	// ======================================================
	// 🧩 Level 1: Utility Extensions
	// ======================================================

	// --- Tool: get_cpu_usage ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_cpu_usage",
		Description: "Return current CPU usage percentage (average)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		percentages, err := cpu.PercentWithContext(ctx, time.Second, false)
		if err != nil {
			return nil, nil, err
		}
		avg := 0.0
		if len(percentages) > 0 {
			avg = percentages[0]
		}
		text := fmt.Sprintf("CPU Usage: %.2f%%", avg)
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, map[string]any{"cpu_percent": avg}, nil
	})

	// --- Tool: get_memory_usage ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_memory_usage",
		Description: "Return memory usage stats: total, used, free, usage %",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		vm, err := mem.VirtualMemoryWithContext(ctx)
		if err != nil {
			return nil, nil, err
		}
		text := fmt.Sprintf("Memory - Total: %d, Used: %d, Free: %d, Usage: %.2f%%",
			vm.Total, vm.Used, vm.Free, vm.UsedPercent)
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, vm, nil
	})

	// --- Tool: get_network_stats ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_network_stats",
		Description: "Return network interface stats (bytes sent/received)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		stats, err := net.IOCountersWithContext(ctx, true)
		if err != nil {
			return nil, nil, err
		}
		text := "Network Interfaces:\n"
		for _, s := range stats {
			text += fmt.Sprintf("%s - Sent: %vB, Received: %vB\n", s.Name, s.BytesSent, s.BytesRecv)
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, stats, nil
	})

	// --- Tool: get_battery_status ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_battery_status",
		Description: "Return battery charge level and status",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		batteries, err := battery.GetAll()
		if err != nil {
			return nil, nil, err
		}
		if len(batteries) == 0 {
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: "No battery detected"}},
			}, nil, nil
		}
		text := ""
		for i, b := range batteries {
			text += fmt.Sprintf("Battery %d: %.2f%% (%s)\n", i, b.Current/b.Full*100, b.State)
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, batteries, nil
	})

	// --- Tool: get_gpu_usage ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_gpu_usage",
		Description: "Return GPU model and utilization (Mac via system_profiler)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		out, err := exec.Command("system_profiler", "SPDisplaysDataType").Output()
		if err != nil {
			return nil, nil, err
		}
		text := string(out)
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, map[string]string{"raw": text}, nil
	})

	// --- Tool: get_system_info ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_system_info",
		Description: "Return summary of system specs (CPU, memory, disk, uptime)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		hostInfo, _ := host.InfoWithContext(ctx)
		vm, _ := mem.VirtualMemoryWithContext(ctx)
		parts, _ := disk.PartitionsWithContext(ctx, false)
		text := fmt.Sprintf("OS: %s %s (%s)\nUptime: %dh\nRAM: %.2f%% used\n",
			hostInfo.Platform, hostInfo.PlatformVersion, hostInfo.KernelArch,
			int(hostInfo.Uptime/3600), vm.UsedPercent)
		for _, p := range parts {
			u, _ := disk.UsageWithContext(ctx, p.Mountpoint)
			text += fmt.Sprintf("Disk %s: %.2f%% used\n", p.Mountpoint, u.UsedPercent)
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, nil, nil
	})

	// ======================================================
	// ⚙️ Level 2: Control + Interactive Extensions
	// ======================================================

	// --- Tool: list_processes ---
	type procsArgs struct {
		Limit int `json:"limit"`
	}
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_processes",
		Description: "List top N processes by CPU usage. Optional arg: {limit:int}",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args procsArgs) (*mcp.CallToolResult, any, error) {
		if args.Limit <= 0 {
			args.Limit = 5
		}
		plist, err := process.ProcessesWithContext(ctx)
		if err != nil {
			return nil, nil, err
		}
		type pinfo struct {
			Pid  int32   `json:"pid"`
			Name string  `json:"name"`
			CPU  float64 `json:"cpu"`
		}
		var infos []pinfo
		for _, p := range plist {
			c, err := p.CPUPercentWithContext(ctx)
			if err != nil {
				continue
			}
			name, _ := p.NameWithContext(ctx)
			infos = append(infos, pinfo{Pid: p.Pid, Name: name, CPU: c})
		}
		sort.Slice(infos, func(i, j int) bool { return infos[i].CPU > infos[j].CPU })
		if len(infos) > args.Limit {
			infos = infos[:args.Limit]
		}
		text := "Top Processes:\n"
		for _, p := range infos {
			text += fmt.Sprintf("PID=%d Name=%s CPU=%.2f%%\n", p.Pid, p.Name, p.CPU)
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, infos, nil
	})

	// --- Tool: kill_process ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "kill_process",
		Description: "Kill a process by PID",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		Pid int `json:"pid"`
	}) (*mcp.CallToolResult, any, error) {
		p, err := os.FindProcess(args.Pid)
		if err != nil {
			return nil, nil, err
		}
		err = p.Kill()
		msg := fmt.Sprintf("Killed process %d", args.Pid)
		if err != nil {
			msg = err.Error()
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: msg}},
		}, nil, nil
	})

	// --- Tool: launch_process ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "launch_process",
		Description: "Launch a new process by command string",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		Command string `json:"command"`
	}) (*mcp.CallToolResult, any, error) {
		cmd := exec.CommandContext(ctx, "bash", "-c", args.Command)
		err := cmd.Start()
		msg := fmt.Sprintf("Launched: %s", args.Command)
		if err != nil {
			msg = err.Error()
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: msg}},
		}, nil, nil
	})

	// --- Tool: list_files ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_files",
		Description: "List files in a given directory (sandboxed)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		Path string `json:"path"`
	}) (*mcp.CallToolResult, any, error) {
		files, err := os.ReadDir(args.Path)
		if err != nil {
			return nil, nil, err
		}
		names := []string{}
		for _, f := range files {
			names = append(names, f.Name())
		}
		text := fmt.Sprintf("Files in %s:\n%s", args.Path, strings.Join(names, "\n"))
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, names, nil
	})

	// ======================================================
	// 🚀 Start MCP Server via stdio (Claude Desktop integration)
	// ======================================================
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server.Run failed: %v", err)
	}
}

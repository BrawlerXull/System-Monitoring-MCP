package tools

import (
	"context"
	"fmt"

	"github.com/brawlerxull/system-monitor-mcp-go/internal/monitor"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterProcessTools registers list_processes, kill_process, and launch_process tools.
func RegisterProcessTools(server *mcp.Server) {
	// --- list_processes ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_processes",
		Description: "List all running processes with PID, name, and CPU usage",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		results, err := monitor.ListProcessesTool()
		if err != nil {
			return nil, nil, err
		}

		text := "Running Processes:\n"
		for _, p := range results {
			text += fmt.Sprintf("PID=%v Name=%v CPU=%.2f%%\n", p["pid"], p["name"], p["cpu"])
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, results, nil
	})

	// --- kill_process ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "kill_process",
		Description: "Kill a process by PID",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		PID int32 `json:"pid"`
	}) (*mcp.CallToolResult, any, error) {
		msg, err := monitor.KillProcessTool(args.PID)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: msg}},
		}, map[string]string{"message": msg}, nil
	})

	// --- launch_process ---
	mcp.AddTool(server, &mcp.Tool{
		Name:        "launch_process",
		Description: "Launch a new process by command and args",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		Command string   `json:"command"`
		Args    []string `json:"args"`
	}) (*mcp.CallToolResult, any, error) {
		msg, err := monitor.LaunchProcessTool(args.Command, args.Args)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: msg}},
		}, map[string]string{"message": msg}, nil
	})
}

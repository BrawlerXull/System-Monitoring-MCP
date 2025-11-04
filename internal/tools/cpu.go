package tools

import (
	"context"
	"fmt"

	"github.com/brawlerxull/system-monitor-mcp-go/internal/monitor"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)


func RegisterCPUTool(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_cpu_usage",
		Description: "Return current CPU usage percentage (average)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		usage, err := monitor.GetCPUUsage(ctx)
		if err != nil {
			return nil, nil, err
		}
		text := fmt.Sprintf("CPU Usage: %.2f%%", usage)
		return TextResult(text, map[string]any{"cpu_percent": usage}), nil, nil
	})
}

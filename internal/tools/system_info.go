package tools

import (
	"context"

	"github.com/brawlerxull/system-monitor-mcp-go/internal/monitor"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterSystemInfoTool registers the get_system_info tool.
func RegisterSystemInfoTool(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_system_info",
		Description: "Return summary of system specs (CPU, memory, disk, uptime)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		info, text := monitor.GetSystemInfo(ctx)
		return TextResult(text, info), nil, nil
	})
}

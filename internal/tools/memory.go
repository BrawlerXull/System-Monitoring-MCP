package tools

import (
	"context"
	"fmt"

	"github.com/brawlerxull/system-monitor-mcp-go/internal/monitor"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterMemoryTool registers the get_memory_usage tool.
func RegisterMemoryTool(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_memory_usage",
		Description: "Return memory usage stats: total, used, free, usage %",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		stats, err := monitor.GetMemoryStats(ctx)
		if err != nil {
			return nil, nil, err
		}
		text := fmt.Sprintf("Memory - Total: %d, Used: %d, Free: %d, Usage: %.2f%%",
			stats.Total, stats.Used, stats.Free, stats.UsedPercent)
		return TextResult(text, stats), nil, nil
	})
}

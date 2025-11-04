package tools

import (
	"context"
	"fmt"

	"github.com/brawlerxull/system-monitor-mcp-go/internal/monitor"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterNetworkTool registers the get_network_stats tool.
func RegisterNetworkTool(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_network_stats",
		Description: "Return network interface stats (bytes sent/received)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		stats, err := monitor.GetNetworkStats(ctx)
		if err != nil {
			return nil, nil, err
		}
		text := "Network Interfaces:\n"
		for _, s := range stats {
			text += fmt.Sprintf("%s - Sent: %vB, Received: %vB\n", s.Name, s.BytesSent, s.BytesRecv)
		}
		return TextResult(text, stats), nil, nil
	})
}

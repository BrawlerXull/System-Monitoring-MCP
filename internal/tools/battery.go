package tools

import (
	"context"
	"fmt"

	"github.com/brawlerxull/system-monitor-mcp-go/internal/monitor"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterBatteryTool registers the get_battery_status tool.
func RegisterBatteryTool(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_battery_status",
		Description: "Return battery charge level and status",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		batts, err := monitor.GetBatteryStatus()
		if err != nil {
			return nil, nil, err
		}
		if len(batts) == 0 {
			return TextResult("No battery detected", nil), nil, nil
		}

		text := ""
		for i, b := range batts {
			current, _ := b["current"].(float64)
			full, _ := b["full"].(float64)
			state, _ := b["state"].(string)

			percent := 0.0
			if full > 0 {
				percent = current / full * 100
			}

			text += fmt.Sprintf("Battery %d: %.2f%% (%s)\n", i, percent, state)
		}

		return TextResult(text, batts), nil, nil
	})
}

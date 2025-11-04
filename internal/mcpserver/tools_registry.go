package mcpserver

import (
	"github.com/brawlerxull/system-monitor-mcp-go/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterTools registers all system monitoring and control tools
// into the MCP server instance.
func RegisterTools(server *mcp.Server) {
	tools.RegisterCPUTool(server)
	tools.RegisterMemoryTool(server)
	tools.RegisterNetworkTool(server)
	tools.RegisterGPUTool(server)
	tools.RegisterSystemInfoTool(server)
	tools.RegisterFileTools(server)
	tools.RegisterProcessTools(server)
	tools.RegisterBatteryTool(server)
}

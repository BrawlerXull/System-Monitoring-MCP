package mcpserver

import (
	"github.com/brawlerxull/system-monitor-mcp-go/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RegisterTools(server *mcp.Server) {
	tools.RegisterCPUTool(server)
}
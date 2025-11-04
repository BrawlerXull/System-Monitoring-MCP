package mcpserver

import "github.com/modelcontextprotocol/go-sdk/mcp"


func New() *mcp.Server {
	// Initialize MCP server metadata
	impl := &mcp.Implementation{
		Name:    "system-monitor-go",
		Version: "0.2.0",
	}
	server := mcp.NewServer(impl, nil)
	RegisterTools(server)
	return server
}
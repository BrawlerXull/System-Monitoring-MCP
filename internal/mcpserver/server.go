package mcpserver

import "github.com/modelcontextprotocol/go-sdk/mcp"

// New initializes and returns a new MCP server instance.
func New() *mcp.Server {
	impl := &mcp.Implementation{
		Name:    "system-monitor-go",
		Version: "0.2.0",
	}
	server := mcp.NewServer(impl, nil)
	RegisterTools(server)
	return server
}

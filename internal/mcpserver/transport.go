package mcpserver

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// StartViaStdio runs the MCP server via stdio (default for Claude Desktop).
func StartViaStdio(ctx context.Context, server *mcp.Server) {
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server.Run failed: %v", err)
	}
}

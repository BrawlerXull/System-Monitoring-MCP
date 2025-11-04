package main

import (
	"context"
	"log"

	"github.com/brawlerxull/system-monitor-mcp-go/internal/mcpserver"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Entry point of the system-monitor-go MCP server.
// Starts the MCP server over stdio for Claude Desktop integration.
func main() {
	ctx := context.Background()

	server := mcpserver.New()

	// Run server using standard I/O transport (Claude Desktop)
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server.Run failed: %v", err)
	}
}

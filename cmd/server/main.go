package server

import (
	"context"
	"log"

	"github.com/brawlerxull/system-monitor-mcp-go/internal/mcpserver"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)


func main(){
	ctx := context.Background()

	server := mcpserver.New()

	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server.Run failed: %v", err)
	}
}
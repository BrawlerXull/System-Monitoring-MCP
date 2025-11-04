package tools

import "github.com/modelcontextprotocol/go-sdk/mcp"

// TextResult is a helper to create a simple text-based MCP response.
func TextResult(text string, data any) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: text}},
	}
}

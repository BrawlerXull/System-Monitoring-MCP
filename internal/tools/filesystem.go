package tools

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterFileTools registers list_files tool.
func RegisterFileTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_files",
		Description: "List files in a given directory (sandboxed)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		Path string `json:"path"`
	}) (*mcp.CallToolResult, any, error) {
		files, err := os.ReadDir(args.Path)
		if err != nil {
			return nil, nil, err
		}
		names := []string{}
		for _, f := range files {
			names = append(names, f.Name())
		}
		text := fmt.Sprintf("Files in %s:\n%s", args.Path, strings.Join(names, "\n"))
		return TextResult(text, names), nil, nil
	})
}

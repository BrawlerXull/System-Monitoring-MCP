package tools

import (
	"context"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterGPUTool registers the get_gpu_usage tool (Mac only example).
func RegisterGPUTool(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_gpu_usage",
		Description: "Return GPU model and utilization (Mac via system_profiler)",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
		out, err := exec.Command("system_profiler", "SPDisplaysDataType").Output()
		if err != nil {
			return nil, nil, err
		}
		return TextResult(string(out), map[string]string{"raw": string(out)}), nil, nil
	})
}

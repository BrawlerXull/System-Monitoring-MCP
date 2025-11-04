# 🖥️ System Monitor MCP Server (Go)

A **Model Context Protocol (MCP)** server written in Go that provides system monitoring, control, and automation tools — designed to integrate seamlessly with **Claude Desktop** or any other MCP-compatible client.

---

## 🚀 Features

| Category | Tools | Description |
|-----------|--------|-------------|
| **System Monitoring** | `get_cpu_usage` | Returns average CPU usage (%) |
| | `get_memory_usage` | Shows total, used, free, and usage % |
| | `get_network_stats` | Displays bytes sent/received by each interface |
| | `get_battery_status` | Reports charge level and state of all batteries |
| | `get_gpu_usage` | Fetches GPU info via macOS `system_profiler` |
| | `get_system_info` | Gives OS, uptime, memory, and disk usage summary |
| **Process Control** | `list_processes` | Lists top N processes by CPU usage |
| | `kill_process` | Terminates a process by PID |
| | `launch_process` | Launches a command or application |
| **File System** | `list_files` | Lists files in a specified directory |

---

## 🧰 Requirements

- **Go 1.22+**
- macOS or Linux (for `system_profiler` or compatible commands)
- Claude Desktop (for MCP integration)
- Dependencies:
  ```bash
  go get github.com/modelcontextprotocol/go-sdk/mcp
  go get github.com/shirou/gopsutil/v3
  go get github.com/distatus/battery
  ```

---

## ⚙️ Installation

```bash
git clone https://github.com/yourusername/system-monitor-mcp-go.git
cd system-monitor-mcp-go
go mod tidy
go build -o system-monitor-mcp-go main.go
```

---

## 🧩 Integration with Claude Desktop

1. Open your **Claude Desktop config** (usually at `~/Library/Application Support/Claude/claude_desktop_config.json`).
2. Add this under `mcpServers`:

```json
{
  "mcpServers": {
    "system-monitor": {
      "command": "/Users/chinmaychaudhari/Documents/coding/golang-mcp/system-monitor-mcp-go",
      "args": [],
      "cwd": "/Users/chinmaychaudhari/Documents/coding/golang-mcp"
    }
  }
}
```

3. Restart Claude Desktop.

4. Ask Claude something like:
   > Ask system-monitor-go to show my current CPU usage.

If your MCP server is running correctly, Claude will invoke the `get_cpu_usage` tool and display the result.

---

## 🧪 Example Prompts

| Prompt | Description |
|--------|-------------|
| "Ask system-monitor-go to show my current CPU usage." | Calls `get_cpu_usage` |
| "Use system-monitor-go to check memory usage stats." | Calls `get_memory_usage` |
| "Ask system-monitor-go to list top 5 processes by CPU usage." | Calls `list_processes` |
| "Ask system-monitor-go to kill process with PID 1234." | Calls `kill_process` |
| "Use system-monitor-go to list files in ~/Documents." | Calls `list_files` |

---

## 🛠️ Development Notes

- Each MCP tool is registered with `mcp.AddTool()`.
- Data returned is both human-readable text and structured JSON.
- `server.Run(ctx, &mcp.StdioTransport{})` ensures Claude can communicate via stdio.

---

## 🧠 Troubleshooting

If you see this error in logs:
```
spawn ./system-monitor-mcp-go ENOENT
```
It means Claude can’t find the binary.  
✅ Fix it by ensuring the binary exists and has execute permissions:

```bash
chmod +x ./system-monitor-mcp-go
```

Or specify the full path in your config file.

---

## 🧑‍💻 Author

**Chinmay Chaudhari**  
Building applied AI systems and backend tools in Go.  

---

## 🪪 License

MIT License © 2025 Chinmay Chaudhari

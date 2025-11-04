package models

// Standard structure for MCP tool responses.
type MCPResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

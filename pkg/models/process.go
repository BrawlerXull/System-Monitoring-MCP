package models

type ProcessInfo struct {
	Pid  int32   `json:"pid"`
	Name string  `json:"name"`
	CPU  float64 `json:"cpu"`
}

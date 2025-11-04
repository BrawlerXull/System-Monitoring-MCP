package monitor

import (
	"context"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryStats(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemoryWithContext(ctx)
}

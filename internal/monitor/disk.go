package monitor

import (
	"context"

	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskUsage(ctx context.Context, path string) (*disk.UsageStat, error) {
	return disk.UsageWithContext(ctx, path)
}

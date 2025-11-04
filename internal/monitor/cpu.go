package monitor

import (
	"context"
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

func GetCPUUsage(ctx context.Context) (float64, error) {
	values, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil || len(values) == 0 {
		return 0, err
	}
	return values[0], nil
}

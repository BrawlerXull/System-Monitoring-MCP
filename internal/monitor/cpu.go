package monitor

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUUsage(ctx context.Context) (float64, error) {
	values, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil || len(values) == 0 {
		return 0, err
	}
	return values[0], nil
}

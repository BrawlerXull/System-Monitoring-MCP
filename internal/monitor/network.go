package monitor

import (
	"context"

	"github.com/shirou/gopsutil/v3/net"
)

func GetNetworkStats(ctx context.Context) ([]net.IOCountersStat, error) {
	return net.IOCountersWithContext(ctx, true)
}

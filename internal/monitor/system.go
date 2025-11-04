package monitor

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetSystemInfo(ctx context.Context) (map[string]any, string) {
	hostInfo, _ := host.InfoWithContext(ctx)
	vm, _ := mem.VirtualMemoryWithContext(ctx)
	parts, _ := disk.PartitionsWithContext(ctx, false)
	text := fmt.Sprintf("OS: %s %s (%s)\nUptime: %dh\nRAM: %.2f%% used\n",
		hostInfo.Platform, hostInfo.PlatformVersion, hostInfo.KernelArch,
		int(hostInfo.Uptime/3600), vm.UsedPercent)
	for _, p := range parts {
		u, _ := disk.UsageWithContext(ctx, p.Mountpoint)
		text += fmt.Sprintf("Disk %s: %.2f%% used\n", p.Mountpoint, u.UsedPercent)
	}
	info := map[string]any{
		"platform": hostInfo.Platform,
		"uptime":   hostInfo.Uptime,
		"ram_used": vm.UsedPercent,
	}
	return info, text
}

package worker

import (
	"log"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

type Stats struct {
	MemStats  *mem.VirtualMemoryStat
	DiskStats *disk.UsageStat
	LoadStats *load.AvgStat
	TaskCount int
}

func (s *Stats) MemToTalKb() uint64 {
	return s.MemStats.Total
}

func (s *Stats) MemAvailableKb() uint64 {
	return s.MemStats.Available
}

func (s *Stats) MemUsedKb() uint64 {
	return s.MemStats.Used
}

func (s *Stats) MemUsedPercent() uint64 {
	return s.MemStats.Available / s.MemStats.Total
}

func (s *Stats) DiskTotal() uint64 {
	return s.DiskStats.Total
}

func (s *Stats) DiskFree() uint64 {
	return s.DiskStats.Free
}

func (s *Stats) DiskUsed() uint64 {
	return s.DiskStats.Used
}

func (s *Stats) CpuUsage() float64 {
	usage, err := cpu.Percent(0, false)

	if err != nil {
		log.Printf("Error reading CPU percentage %v", err)
		return 0.00
	}

	return usage[0]
}

func GetStats() *Stats {
	return &Stats{
		MemStats:  GetMemoryInfo(),
		DiskStats: GetDiskInfo(),
		LoadStats: GetLoadAvg(),
	}
}

func GetMemoryInfo() *mem.VirtualMemoryStat {
	memstats, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error reading from /proc/meminfo")
		return &mem.VirtualMemoryStat{}
	}

	return memstats
}

func GetDiskInfo() *disk.UsageStat {
	diskstats, err := disk.Usage("/")
	if err != nil {
		log.Printf("Error reading from /")
		return &disk.UsageStat{}
	}

	return diskstats
}

func GetLoadAvg() *load.AvgStat {
	loadavg, err := load.Avg()
	if err != nil {
		log.Printf("Error reading from /proc/loadavg")
		return &load.AvgStat{}
	}

	return loadavg
}

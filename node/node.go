package node

type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int
	Role            string
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	TaskCount       int
}

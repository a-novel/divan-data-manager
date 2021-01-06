package divan_types

type SystemStats struct {
	CpuUtilizationRate float64 `json:"cpu_utilization_rate"`
	CpuStolenRate      float64 `json:"cpu_stolen_rate"`
	SwapTotal          uint64    `json:"swap_total"`
	SwapUsed           uint64    `json:"swap_used"`
	MemTotal           uint64    `json:"mem_total"`
	MemFree            uint64    `json:"mem_free"`
	MemLimit           uint64    `json:"mem_limit"`
	CpuCoresAvailable  uint64    `json:"cpu_cores_available"`
	AllocStall         uint64    `json:"allocstall"`
}

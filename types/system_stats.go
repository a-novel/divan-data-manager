package divan_types

type SystemStats struct {
	CpuUtilizationRate float64 `json:"cpu_utilization_rate"`
	CpuStolenRate      float64 `json:"cpu_stolen_rate"`
	SwapTotal          uint    `json:"swap_total"`
	SwapUsed           uint    `json:"swap_used"`
	MemTotal           uint    `json:"mem_total"`
	MemFree            uint    `json:"mem_free"`
	MemLimit           uint    `json:"mem_limit"`
	CpuCoresAvailable  uint    `json:"cpu_cores_available"`
	AllocStall         uint    `json:"allocstall"`
}

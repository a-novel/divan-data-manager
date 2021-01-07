package divan_types

type FastWarmupSettings struct {
	FastWarmupEnabled  bool   `json:"fastWarmupEnabled"`
	MinItemsThreshold  uint64 `json:"minItemsThreshold"`
	MinMemoryThreshold uint64 `json:"minMemoryThreshold"`
}

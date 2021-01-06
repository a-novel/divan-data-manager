package divan_types

type FastWarmupSettings struct {
	FastWarmupEnabled  bool `json:"fastWarmupEnabled"`
	MinItemsThreshold  uint `json:"minItemsThreshold"`
	MinMemoryThreshold uint `json:"minMemoryThreshold"`
}

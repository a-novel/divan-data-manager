package divan_types

type Controllers struct {
	CompactAll    string `json:"compactAll"`
	CompactDB     string `json:"compactDB"`
	PurgeDeletes  string `json:"purgeDeletes"`
	StartRecovery string `json:"startRecovery"`
}

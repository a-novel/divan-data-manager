package divan_types

type StorageTotals struct {
	Ram RamStorage `json:"ram"`
	HDD HDDStorage `json:"hdd"`
}

type RamStorage struct {
	Total             uint64 `json:"total"`
	QuotaTotal        uint64 `json:"quotaTotal"`
	QuotaUsed         uint64 `json:"quotaUsed"`
	Used              uint64 `json:"used"`
	UsedByData        uint64 `json:"usedByData"`
	QuotaUsedPerNode  uint64 `json:"quotaUsedPerNode"`
	QuotaTotalPerNode uint64 `json:"quotaTotalPerNode"`
}

type HDDStorage struct {
	Total      uint64 `json:"total"`
	QuotaTotal uint64 `json:"quotaTotal"`
	Used       uint64 `json:"used"`
	UsedByData uint64 `json:"usedByData"`
	Free       uint64 `json:"free"`
}

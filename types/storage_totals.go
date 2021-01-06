package divan_types

type StorageTotals struct {
	Ram RamStorage `json:"ram"`
	HDD HDDStorage `json:"hdd"`
}

type RamStorage struct {
	Total             uint `json:"total"`
	QuotaTotal        uint `json:"quotaTotal"`
	QuotaUsed         uint `json:"quotaUsed"`
	Used              uint `json:"used"`
	UsedByData        uint `json:"usedByData"`
	QuotaUsedPerNode  uint `json:"quotaUsedPerNode"`
	QuotaTotalPerNode uint `json:"quotaTotalPerNode"`
}

type HDDStorage struct {
	Total      uint `json:"total"`
	QuotaTotal uint `json:"quotaTotal"`
	Used       uint `json:"used"`
	UsedByData uint `json:"usedByData"`
	Free       uint `json:"free"`
}

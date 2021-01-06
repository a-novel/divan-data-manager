package divan_types

type BasicStats struct {
	VbActiveNumNonResident uint    `json:"vbActiveNumNonResident"`
	QuotaPercentUsed       float64 `json:"quotaPercentUsed"`
	OpsPerSec              int     `json:"opsPerSec"`
	DiskFetches            int     `json:"diskFetches"`
	ItemCount              int     `json:"itemCount"`
	DiskUsed               int     `json:"diskUsed"`
	DataUsed               int     `json:"dataUsed"`
	MemUsed                uint    `json:"memUsed"`
}

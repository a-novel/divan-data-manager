package divan_types

type InterestingStats struct {
	CmdGet                   uint64  `json:"cmd_get"`
	CouchDocsActualDiskSize  uint64  `json:"couch_docs_actual_disk_size"`
	CouchDocsDataSize        uint64  `json:"couch_docs_data_size"`
	CouchSpatialDataSize     uint64  `json:"couch_spatial_data_size"`
	CouchSpatialDiskSize     uint64  `json:"couch_spatial_disk_size"`
	CouchViewsActualDiskSize uint64  `json:"couch_views_actual_disk_size"`
	CouchViewsDataSize       uint64  `json:"couch_views_data_size"`
	CurrItems                uint64  `json:"curr_items"`
	CurrItemsTot             uint64  `json:"curr_items_tot"`
	EpBgFetched              uint64  `json:"ep_bg_fetched"`
	GetHits                  uint64  `json:"get_hits"`
	MemUsed                  uint64  `json:"mem_used"`
	Ops                      float64 `json:"ops"`
	VbActiveNumNonResident   uint64  `json:"vb_active_num_non_resident"`
	VbReplicaCurrItems       uint64  `json:"vb_replica_curr_items"`
}

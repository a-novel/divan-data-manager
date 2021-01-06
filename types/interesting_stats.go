package divan_types

type InterestingStats struct {
	CmdGet                   uint `json:"cmd_get"`
	CouchDocsActualDiskSize  uint `json:"couch_docs_actual_disk_size"`
	CouchDocsDataSize        uint `json:"couch_docs_data_size"`
	CouchSpatialDataSize     uint `json:"couch_spatial_data_size"`
	CouchSpatialDiskSize     uint `json:"couch_spatial_disk_size"`
	CouchViewsActualDiskSize uint `json:"couch_views_actual_disk_size"`
	CouchViewsDataSize       uint `json:"couch_views_data_size"`
	CurrItems                uint `json:"curr_items"`
	CurrItemsTot             uint `json:"curr_items_tot"`
	EpBgFetched              uint `json:"ep_bg_fetched"`
	GetHits                  uint `json:"get_hits"`
	MemUsed                  uint `json:"mem_used"`
	Ops                      uint `json:"ops"`
	VbActiveNumNonResident   uint `json:"vb_active_num_non_resident"`
	VbReplicaCurrItems       uint `json:"vb_replica_curr_items"`
}

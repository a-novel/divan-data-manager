package divan_data_manager

import "github.com/a-novel/divan-data-manager/types"

type ClusterData struct {
	Name                    string                         `json:"name"`
	Nodes                   []divan_types.Node             `json:"nodes"`
	Buckets                 divan_types.ClusterBuckets     `json:"buckets"`
	RemoteClusters          divan_types.RemoteClusters     `json:"remoteClusters"`
	Alerts                  []interface{}                  `json:"alerts"`
	AlertsSilenceURL        string                         `json:"alertsSilenceURL"`
	Controllers             divan_types.Controllers        `json:"controllers"`
	RebalanceStatus         string                         `json:"rebalanceStatus"`
	RebalanceProgressUri    string                         `json:"rebalanceProgressUri"`
	StopRebalanceUri        string                         `json:"stopRebalanceUri"`
	NodeStatusesUri         string                         `json:"nodeStatusesUri"`
	MaxBucketCount          uint                           `json:"maxBucketCount"`
	Tasks                   divan_types.Tasks              `json:"tasks"`
	Counters                divan_types.Counters           `json:"counters"`
	IndexStatusURI          string                         `json:"indexStatusURI"`
	CheckPermissionsURI     string                         `json:"checkPermissionsURI"`
	ServerGroupsUri         string                         `json:"serverGroupsUri"`
	ClusterName             string                         `json:"clusterName"`
	Balanced                bool                           `json:"balanced"`
	MemoryQuota             uint                           `json:"memoryQuota"`
	IndexMemoryQuota        uint                           `json:"indexMemoryQuota"`
	FtsMemoryQuota          uint                           `json:"ftsMemoryQuota"`
	CbasMemoryQuota         uint                           `json:"cbasMemoryQuota"`
	EventingMemoryQuota     uint                           `json:"eventingMemoryQuota"`
	AutoCompactionSettings  interface{}                    `json:"autoCompactionSettings"`
	FastWarmupSettings      divan_types.FastWarmupSettings `json:"fastWarmupSettings"`
	VisualSettingsUri       string                         `json:"visualSettingsUri"`
	StorageTotals           divan_types.StorageTotals      `json:"storageTotals"`
	AutoCompactionSettingsD *divan_types.AutoCompactionSettings
}

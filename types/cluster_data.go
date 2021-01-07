package divan_types

type ClusterData struct {
	Name                    string             `json:"name"`
	Nodes                   []Node             `json:"nodes"`
	Buckets                 ClusterBuckets     `json:"buckets"`
	RemoteClusters          RemoteClusters     `json:"remoteClusters"`
	Alerts                  []interface{}      `json:"alerts"`
	AlertsSilenceURL        string             `json:"alertsSilenceURL"`
	Controllers             Controllers        `json:"controllers"`
	RebalanceStatus         string             `json:"rebalanceStatus"`
	RebalanceProgressUri    string             `json:"rebalanceProgressUri"`
	StopRebalanceUri        string             `json:"stopRebalanceUri"`
	NodeStatusesUri         string             `json:"nodeStatusesUri"`
	MaxBucketCount          uint64             `json:"maxBucketCount"`
	Tasks                   Tasks              `json:"tasks"`
	Counters                Counters           `json:"counters"`
	IndexStatusURI          string             `json:"indexStatusURI"`
	CheckPermissionsURI     string             `json:"checkPermissionsURI"`
	ServerGroupsUri         string             `json:"serverGroupsUri"`
	ClusterName             string             `json:"clusterName"`
	Balanced                bool               `json:"balanced"`
	MemoryQuota             uint64             `json:"memoryQuota"`
	IndexMemoryQuota        uint64             `json:"indexMemoryQuota"`
	FtsMemoryQuota          uint64             `json:"ftsMemoryQuota"`
	CbasMemoryQuota         uint64             `json:"cbasMemoryQuota"`
	EventingMemoryQuota     uint64             `json:"eventingMemoryQuota"`
	AutoCompactionSettings  interface{}        `json:"autoCompactionSettings"`
	FastWarmupSettings      FastWarmupSettings `json:"fastWarmupSettings"`
	VisualSettingsUri       string             `json:"visualSettingsUri"`
	StorageTotals           StorageTotals      `json:"storageTotals"`
	AutoCompactionSettingsD *AutoCompactionSettings
}

package divan_types

type ClusterControllers struct {
	AddNode               DefaultController     `json:"addNode"`
	Rebalance             DefaultController     `json:"rebalance"`
	FailOver              DefaultController     `json:"failOver"`
	StartGracefulFailover DefaultController     `json:"startGracefulFailover"`
	ReAddNode             DefaultController     `json:"reAddNode"`
	ReFailOver            DefaultController     `json:"reFailOver"`
	EjectNode             DefaultController     `json:"ejectNode"`
	SetRecoveryType       DefaultController     `json:"setRecoveryType"`
	SetAutoCompaction     SetAutoCompaction     `json:"setAutoCompaction"`
	ClusterLogsCollection ClusterLogsCollection `json:"clusterLogsCollection"`
	SetFastWarmup         SetFastWarmup         `json:"setFastWarmup"`
	Replication           Replication           `json:"replication"`
}

type DefaultController struct {
	Uri string `json:"uri"`
}

type SetAutoCompaction struct {
	Uri         string `json:"uri"`
	ValidateURI string `json:"validateURI"`
}

type SetFastWarmup struct {
	Uri         string `json:"uri"`
	ValidateURI string `json:"validateURI"`
}

type ClusterLogsCollection struct {
	StartURI  string `json:"startURI"`
	CancelURI string `json:"cancelURI"`
}

type Replication struct {
	CreateURI   string `json:"createURI"`
	ValidateURI string `json:"validateURI"`
}

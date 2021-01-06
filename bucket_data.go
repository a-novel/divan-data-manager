package divan_data_manager

import "github.com/a-novel/divan-data-manager/types"

type BucketData struct {
	Name                    string                       `json:"name"`
	BucketType              string                       `json:"bucketType"`
	AuthType                string                       `json:"authType"`
	Uuid                    string                       `json:"uuid"`
	Uri                     string                       `json:"uri"`
	StreamingUri            string                       `json:"streamingUri"`
	LocalRandomKeyUri       string                       `json:"localRandomKeyUri"`
	Controllers             divan_types.Controllers      `json:"controllers"`
	BasicStats              divan_types.BasicStats       `json:"basicStats"`
	Nodes                   []divan_types.Node           `json:"nodes"`
	Stats                   divan_types.Stats            `json:"stats"`
	NodeLocator             string                       `json:"nodeLocator"`
	SaslPassword            string                       `json:"saslPassword"`
	PurgeInterval           float64                      `json:"purgeInterval"`
	VBucketServerMap        divan_types.VBucketServerMap `json:"vBucketServerMap"`
	ReplicaNumber           uint                         `json:"replicaNumber"`
	ThreadsNumber           uint                         `json:"threadsNumber"`
	Quota                   divan_types.Quota            `json:"quota"`
	EvictionPolicy          string                       `json:"evictionPolicy"`
	DurabilityMinLevel      string                       `json:"durabilityMinLevel"`
	ConflictResolutionType  string                       `json:"conflictResolutionType"`
	BucketCapabilitiesVer   string                       `json:"bucketCapabilitiesVer"`
	BucketCapabilities      []string                     `json:"bucketCapabilities"`
	Ddocs                   divan_types.Ddocs            `json:"ddocs"`
	ReplicaIndex            bool                         `json:"replicaIndex"`
	AutoCompactionSettings  interface{}                  `json:"autoCompactionSettings"`
	AutoCompactionSettingsD *divan_types.AutoCompactionSettings
}

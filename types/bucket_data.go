package divan_types

import "encoding/json"

type BucketData struct {
	Name                    string           `json:"name"`
	BucketType              string           `json:"bucketType"`
	AuthType                string           `json:"authType"`
	Uuid                    string           `json:"uuid"`
	Uri                     string           `json:"uri"`
	StreamingUri            string           `json:"streamingUri"`
	LocalRandomKeyUri       string           `json:"localRandomKeyUri"`
	Controllers             Controllers      `json:"controllers"`
	BasicStats              BasicStats       `json:"basicStats"`
	Nodes                   []Node           `json:"nodes"`
	Stats                   Stats            `json:"stats"`
	NodeLocator             string           `json:"nodeLocator"`
	SaslPassword            string           `json:"saslPassword"`
	PurgeInterval           float64          `json:"purgeInterval"`
	VBucketServerMap        VBucketServerMap `json:"vBucketServerMap"`
	ReplicaNumber           uint64           `json:"replicaNumber"`
	ThreadsNumber           uint64           `json:"threadsNumber"`
	Quota                   Quota            `json:"quota"`
	EvictionPolicy          string           `json:"evictionPolicy"`
	DurabilityMinLevel      string           `json:"durabilityMinLevel"`
	ConflictResolutionType  string           `json:"conflictResolutionType"`
	BucketCapabilitiesVer   string           `json:"bucketCapabilitiesVer"`
	BucketCapabilities      []string         `json:"bucketCapabilities"`
	Ddocs                   Ddocs            `json:"ddocs"`
	ReplicaIndex            bool             `json:"replicaIndex"`
	AutoCompactionSettings  interface{}      `json:"autoCompactionSettings"`
	AutoCompactionSettingsD *AutoCompactionSettings
}

func (bd *BucketData) GetAutocompaction() error {
	if bd.AutoCompactionSettings != nil {
		marshalled, err := json.Marshal(bd.AutoCompactionSettings)
		if err != nil {
			return err
		}

		var output AutoCompactionSettings
		if err := json.Unmarshal(marshalled, &output); err != nil {
			return err
		}

		output.Parse()
		bd.AutoCompactionSettingsD = &output
	}

	return nil
}

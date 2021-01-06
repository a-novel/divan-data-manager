package divan_types

type VBucketServerMap struct {
	HashAlgorithm string   `json:"hashAlgorithm"`
	NumReplicas   uint     `json:"numReplicas"`
	ServerList    []string `json:"serverList"`
	VBucketMap    [][]uint `json:"vBucketMap"`
}

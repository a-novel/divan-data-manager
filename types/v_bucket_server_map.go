package divan_types

type VBucketServerMap struct {
	HashAlgorithm string   `json:"hashAlgorithm"`
	NumReplicas   uint64     `json:"numReplicas"`
	ServerList    []string `json:"serverList"`
	VBucketMap    [][]uint64 `json:"vBucketMap"`
}

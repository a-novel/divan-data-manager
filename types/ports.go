package divan_types

type Ports struct {
	Proxy     uint64 `json:"proxy"`
	HttpsCAPI uint64 `json:"httpsCAPI"`
	HttpsMgmt uint64 `json:"httpsMgmt"`
	Direct    uint64 `json:"direct"`
	DistTCP   uint64 `json:"distTCP"`
	DistTLS   uint64 `json:"distTLS"`
}

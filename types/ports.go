package divan_types

type Ports struct {
	Proxy     uint `json:"proxy"`
	HttpsCAPI uint `json:"httpsCAPI"`
	HttpsMgmt uint `json:"httpsMgmt"`
	Direct    uint `json:"direct"`
	DistTCP   uint `json:"distTCP"`
	DistTLS   uint `json:"distTLS"`
}

package divan_types

type Node struct {
	CouchApiBase         string             `json:"couchApiBase"`
	SystemStats          SystemStats        `json:"systemStats"`
	InterestingStats     InterestingStats   `json:"interestingStats"`
	Uptime               string             `json:"uptime"`
	MemoryTotal          uint64               `json:"memoryTotal"`
	MemoryFree           uint64               `json:"memoryFree"`
	McdMemoryReserved    uint64               `json:"mcdMemoryReserved"`
	McdMemoryAllocated   uint64               `json:"mcdMemoryAllocated"`
	Replication          uint64               `json:"replication"`
	ClusterMembership    string             `json:"clusterMembership"`
	RecoveryType         string             `json:"recoveryType"`
	Status               string             `json:"status"`
	OtpNode              string             `json:"otpNode"`
	ThisNode             bool               `json:"thisNode"`
	Hostname             string             `json:"hostname"`
	NodeUUID             string             `json:"nodeUUID"`
	ClusterCompatibility uint64             `json:"clusterCompatibility"`
	Version              string             `json:"version"`
	Os                   string             `json:"os"`
	CpuCount             uint64               `json:"cpuCount"`
	Ports                Ports              `json:"ports"`
	Services             []string           `json:"services"`
	NodeEncryption       bool               `json:"nodeEncryption"`
	ConfiguredHostname   string             `json:"configuredHostname"`
	AddressFamily        string             `json:"addressFamily"`
	ExternalListeners    []ExternalListener `json:"externalListeners"`
}

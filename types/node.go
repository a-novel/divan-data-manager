package divan_types

type Node struct {
	CouchApiBase         string             `json:"couchApiBase"`
	SystemStats          SystemStats        `json:"systemStats"`
	InterestingStats     InterestingStats   `json:"interestingStats"`
	Uptime               string             `json:"uptime"`
	MemoryTotal          uint               `json:"memoryTotal"`
	MemoryFree           uint               `json:"memoryFree"`
	McdMemoryReserved    uint               `json:"mcdMemoryReserved"`
	McdMemoryAllocated   uint               `json:"mcdMemoryAllocated"`
	Replication          uint               `json:"replication"`
	ClusterMembership    string             `json:"clusterMembership"`
	RecoveryType         string             `json:"recoveryType"`
	Status               string             `json:"status"`
	OtpNode              string             `json:"otpNode"`
	ThisNode             bool               `json:"thisNode"`
	Hostname             string             `json:"hostname"`
	NodeUUID             string             `json:"nodeUUID"`
	ClusterCompatibility string             `json:"clusterCompatibility"`
	Version              string             `json:"version"`
	Os                   string             `json:"os"`
	CpuCount             uint               `json:"cpuCount"`
	Ports                Ports              `json:"ports"`
	Services             []string           `json:"services"`
	NodeEncryption       bool               `json:"nodeEncryption"`
	ConfiguredHostname   string             `json:"configuredHostname"`
	AddressFamily        string             `json:"addressFamily"`
	ExternalListeners    []ExternalListener `json:"externalListeners"`
}

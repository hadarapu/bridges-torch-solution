package common

type InputObject struct {
	BridgeId      int
	PersonIdsList []int
	Cfg *ConfigInfo
}

type OutputObject struct {
	BridgeId int
	QuickestTime float64
}


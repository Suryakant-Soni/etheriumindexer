package beacon

type GenesisResponse struct {
	Data GenesisData `json:"data"`
}

type GenesisData struct {
	GenesisTime string `json:"genesis_time"`
}

type BeaconConfigResponse struct {
	Data ConfigData `json:"data"`
}

type ConfigData struct {
	SecondsPerSlot string `json:"SECONDS_PER_SLOT"`
	SlotsPerEpoch  string `json:"SLOTS_PER_EPOCH"`
}

type CommitteeResponse struct {
	Data []CommitteeItem `json:"data"`
}

type CommitteeItem struct {
	Index      string   `json:"indesx"`
	Slot       string   `json:"slot"`
	Validators []string `json:"validators"`
}

type AttestationResponse struct {
	Data []AttestationItem `json:"data"`
}

type AttestationItem struct {
	AggregationBits string `json:"aggregation_bits"`
}

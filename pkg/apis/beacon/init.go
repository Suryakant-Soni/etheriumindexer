package beacon

import (
	"encoding/json"
	httpClient "etheriumindexer/pkg/client/http"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
)

type Endpoints struct {
	GetGenesisTime       string `json:"get_genesis_time"`
	EtheriumConfig       string `json:"etherium_config"`
	GetActiveValidators  string `json:"get_active_validators"`
	GetAttestationStatus string `json:"get_attestation_status"`
}

type Beacon struct {
	BaseUrl   string
	Endpoints Endpoints
	Config    BeaconConfigResponse
}

// var beacon Beacon
func NewBeaconClient() *Beacon {
	_, f, _, _ := runtime.Caller(0)
	fname := filepath.Join(filepath.Dir(f), "..", "..", "..", "config", "end_points.json")
	// endpointsFile, err := os.OpenFile(fname, os.O_RDONLY, 0400)
	endpointsFile, err := os.Open(fname)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer endpointsFile.Close()
	decoder := json.NewDecoder(endpointsFile)
	b := &Beacon{
		Endpoints: Endpoints{},
	}
	err = decoder.Decode(&b.Endpoints)

	scheme := "https"

	u := &url.URL{
		Scheme: scheme,
		Host:   os.Getenv("QUICKNODE_HOST"),
		Path:   os.Getenv("QUICKNODE_KEY"),
	}
	b.BaseUrl = u.String()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	b.injectBeaconConfig()
	return b
}

func (b *Beacon) injectBeaconConfig() (r *BeaconConfigResponse) {
	fullPath := b.BaseUrl + b.Endpoints.EtheriumConfig
	r, err := httpClient.Get[BeaconConfigResponse](fullPath, &b.Config)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return r
}

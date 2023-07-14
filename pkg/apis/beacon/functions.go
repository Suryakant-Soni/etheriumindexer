package beacon

import (
	httpClient "etheriumindexer/pkg/client/http"
	"fmt"
	"log"
)

func (b *Beacon) GetCurrentEpoch() (r *GenesisResponse) {
	fullPath := b.BaseUrl + b.Endpoints.GetGenesisTime
	r, err := httpClient.Get[GenesisResponse](fullPath, &GenesisResponse{})
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return r
}

func (b *Beacon) GetActiveValidators(slot any) (r any) {
	fullPath := b.BaseUrl + fmt.Sprintf(b.Endpoints.GetActiveValidators, slot)
	r, err := httpClient.Get[CommitteeResponse](fullPath, &CommitteeResponse{})
	if err != nil {
		log.Println("Error: ", err)
	}
	// log.Println("active validators", r)
	return r
}

func (b *Beacon) GetAttestationStatus(commResp any) (r any) {
	slot := ""
	if commRespStruct, ok := commResp.(*CommitteeResponse); ok {
		if commRespStruct != nil && commRespStruct.Data != nil && len(commRespStruct.Data) > 0 && commRespStruct.Data[0].Slot != "" {
			slot = commRespStruct.Data[0].Slot
		}
	} else {
		// data is not of type MyStruct
		fmt.Println("Unknown type")
		return nil
	}

	fullPath := b.BaseUrl + fmt.Sprintf(b.Endpoints.GetAttestationStatus, slot)
	r, err := httpClient.Get[AttestationResponse](fullPath, &AttestationResponse{})
	if err != nil {
		log.Println("Error: ", err)
	}

	// if attestationStruct, ok := r.(AttestationResponse); ok {
	// 	if attestationStruct.Data != nil && len(attestationStruct.Data) > 0 {
	// 		log.Println("AttestationResponse", attestationStruct.Data)
	// 	}
	// }

	return r
}

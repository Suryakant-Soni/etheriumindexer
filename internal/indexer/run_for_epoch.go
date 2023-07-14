package indexer

import (
	"etheriumindexer/internal/utils"
	"log"
	"sync"
)

func RunForEpoch(epoch int, wg *sync.WaitGroup) {
	slotsCh := supplySlots(epoch)
	defer func(epoch int) {
		log.Println("done for epoch", epoch)
		wg.Done()
	}(epoch)
	ActiveValidatorsCh := Pipeline(slotsCh, func(a any) any {
		return b.GetActiveValidators(a)
	})
	AttestationStatusCh := Pipeline(ActiveValidatorsCh, func(a any) any {
		return b.GetAttestationStatus(a)
	})

	if AttestationStatusCh != nil {
		log.Print("")
	}
	// filenames := pipeline(webpImages, saveToDisk)

	for in := range AttestationStatusCh {
		log.Printf("\naggregation bits %v ", in)
	}
}

func supplySlots(epoch int) <-chan any {
	slotsCh := make(chan any, 32)
	slotsPerEpoch := utils.MustAtoi(b.Config.Data.SlotsPerEpoch)
	for i := epoch * slotsPerEpoch; i < (epoch+1)*slotsPerEpoch; i++ {
		slotsCh <- i
	}
	return slotsCh
}

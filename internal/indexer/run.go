package indexer

import (
	"etheriumindexer/internal/utils"
	"etheriumindexer/pkg/apis/beacon"
	"log"
	"sync"
)

type Indexer struct {
	Epoch int
}

var indexer Indexer

var b *beacon.Beacon

func Run() {
	var wg sync.WaitGroup
	b = beacon.NewBeaconClient()
	genesisData := b.GetCurrentEpoch()
	log.Printf("genesis data %v", genesisData.Data)
	calculateEpoch(utils.MustAtoi(genesisData.Data.GenesisTime), b)

	for i := indexer.Epoch - 4; i <= indexer.Epoch; i++ {
		wg.Add(1)
		log.Println("epoch - ", i)
		go RunForEpoch(i, &wg)
	}
	wg.Wait()
}

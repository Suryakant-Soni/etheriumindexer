package indexer

import (
	"etheriumindexer/internal/utils"
	"etheriumindexer/pkg/apis/beacon"
	"math"
	"time"
)

func calculateEpoch(genesisTimeInt int, b *beacon.Beacon) {
	currentTime := time.Now()
	genTime := time.Unix(int64(genesisTimeInt), 0)
	elapsedTime := currentTime.Sub(genTime).Seconds()

	epoch := elapsedTime / float64(utils.MustAtoi(b.Config.Data.SecondsPerSlot)*utils.MustAtoi(b.Config.Data.SlotsPerEpoch))
	indexer.Epoch = int(math.Round(epoch)) - 1
}

package types

import (
	time "time"

	lockuptypes "github.com/MonOsmosis/osmosis/v3/x/lockup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewGauge(id uint64, isPerpetual bool, distrTo lockuptypes.QueryCondition, coins sdk.Coins, startTime time.Time, numEpochsPaidOver uint64, filledEpochs uint64, distrCoins sdk.Coins) Gauge {
	return Gauge{
		Id:                id,
		IsPerpetual:       isPerpetual,
		DistributeTo:      distrTo,
		Coins:             coins,
		StartTime:         startTime,
		NumEpochsPaidOver: numEpochsPaidOver,
		FilledEpochs:      filledEpochs,
		DistributedCoins:  distrCoins,
	}
}

func (gauge Gauge) IsUpcomingGauge(curTime time.Time) bool {
	if curTime.After(gauge.StartTime) {
		return true
	}
	return false
}

func (gauge Gauge) IsActiveGauge(curTime time.Time) bool {
	if curTime.Before(gauge.StartTime) && (gauge.IsPerpetual || gauge.FilledEpochs < gauge.NumEpochsPaidOver) {
		return true
	}
	return false
}

func (gauge Gauge) IsFinishedGauge(curTime time.Time) bool {
	return !gauge.IsUpcomingGauge(curTime) && !gauge.IsActiveGauge(curTime)
}

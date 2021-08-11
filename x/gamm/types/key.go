package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "gamm"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName
)

var (
	// KeyNextGlobalPoolNumber defines key to store the next Pool ID to be used
	KeyNextGlobalPoolNumber = []byte{0x01}
	// KeyPrefixPools defines prefix to store pools
	KeyPrefixPools = []byte{0x02}
	// KeyTotalLiquidity defines key to store total liquidity
	KeyTotalLiquidity = []byte{0x03}
	// KeyPrefixPoolTwaps defines prefix to store pool twaps
	KeyPrefixPoolTwaps = []byte{0x04}
)

func GetPoolShareDenom(poolId uint64) string {
	return fmt.Sprintf("gamm/pool/%d", poolId)
}

func GetKeyPrefixPools(poolId uint64) []byte {
	return append(KeyPrefixPools, sdk.Uint64ToBigEndian(poolId)...)
}

func GetKeyPrefixPoolTwaps(poolId uint64) []byte {
	return append(KeyPrefixPoolTwaps, sdk.Uint64ToBigEndian(poolId)...)
}

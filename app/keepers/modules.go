package keepers

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	transfer "github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v3/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v3/modules/core/02-client/client"
	"github.com/osmosis-labs/bech32-ibc/x/bech32ibc"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"

	_ "github.com/MonOsmosis/osmosis/v10/client/docs/statik"
	"github.com/MonOsmosis/osmosis/v10/x/epochs"
	"github.com/MonOsmosis/osmosis/v10/x/gamm"
	"github.com/MonOsmosis/osmosis/v10/x/incentives"
	"github.com/MonOsmosis/osmosis/v10/x/lockup"
	"github.com/MonOsmosis/osmosis/v10/x/mint"
	poolincentives "github.com/MonOsmosis/osmosis/v10/x/pool-incentives"
	poolincentivesclient "github.com/MonOsmosis/osmosis/v10/x/pool-incentives/client"
	superfluid "github.com/MonOsmosis/osmosis/v10/x/superfluid"
	superfluidclient "github.com/MonOsmosis/osmosis/v10/x/superfluid/client"
	"github.com/MonOsmosis/osmosis/v10/x/tokenfactory"
	"github.com/MonOsmosis/osmosis/v10/x/txfees"
)

// AppModuleBasics returns ModuleBasics for the module BasicManager.
var AppModuleBasics = []module.AppModuleBasic{
	auth.AppModuleBasic{},
	genutil.AppModuleBasic{},
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	distr.AppModuleBasic{},
	gov.NewAppModuleBasic(
		append(
			wasmclient.ProposalHandlers,
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
			upgradeclient.CancelProposalHandler,
			poolincentivesclient.UpdatePoolIncentivesHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
			superfluidclient.SetSuperfluidAssetsProposalHandler,
			superfluidclient.RemoveSuperfluidAssetsProposalHandler,
		)...,
	),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	ibc.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	vesting.AppModuleBasic{},
	gamm.AppModuleBasic{},
	txfees.AppModuleBasic{},
	incentives.AppModuleBasic{},
	lockup.AppModuleBasic{},
	poolincentives.AppModuleBasic{},
	epochs.AppModuleBasic{},
	superfluid.AppModuleBasic{},
	tokenfactory.AppModuleBasic{},
	bech32ibc.AppModuleBasic{},
	wasm.AppModuleBasic{},
}

package oracle

import (
	"context"
	"encoding/json"

	"cosmossdk.io/core/appmodule"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/skip-mev/slinky/x/oracle/client/cli"
	"github.com/skip-mev/slinky/x/oracle/keeper"
	"github.com/skip-mev/slinky/x/oracle/types"
	"github.com/spf13/cobra"
)

// ConsensusVersion is the x/oracle module's current version, as modules integrate and updates are made, this value determines what
// version of the module is being run by the chain.
const ConsensusVersion = 1

var (
	_ module.AppModule      = AppModule{}
	_ appmodule.AppModule   = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the base interface that the x/oracle module exposes to the application.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the name of this module
func (amb AppModuleBasic) Name() string { return types.ModuleName }

// RegisterLegacyAminoCodec registers the necessary types from the x/oracle module for amino serialization.
func (amb AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the necessary implementations / interfaces in the x/oracle module w/ the interface-registry ir.
func (amb AppModuleBasic) RegisterInterfaces(ir codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(ir)
}

// RegisterGRPCGatewayRoutes registers the necessary REST routes for the GRPC-gateway to the x/oracle module QueryService on mux. This method
// panics on failure
func (amb AppModuleBasic) RegisterGRPCGatewayRoutes(cliCtx client.Context, mux *runtime.ServeMux) {
	// register the gate-way routes w/ the provided mux
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(cliCtx)); err != nil {
		panic(err)
	}
}

// GetTxCmd is a no-op, as no txs are registered for submission (apart from messages that can only be executed by governance)
func (amb AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns the x/oracle module base query cli-command
func (amb AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// AppModule represents an application module for the x/oracle module
type AppModule struct {
	AppModuleBasic

	k keeper.Keeper
}

// NewAppModule returns an application module for the x/oracle module
func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{
			cdc: cdc,
		},
		k: k,
	}
}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// RegisterServices registers the module's services with the app's module configurator
func (am AppModule) RegisterServices(cfc module.Configurator) {
	// register MsgServer
	types.RegisterMsgServer(cfc.MsgServer(), keeper.NewMsgServer(am.k))
	// register Query Service
	types.RegisterQueryServer(cfc.QueryServer(), keeper.NewQueryServer(am.k))
}

// DefaultGenesis returns default genesis state as raw bytes for the oracle
// module.
func (AppModuleBasic) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	// by default no CurrencyPairs will be added to state initially
	return json.RawMessage{}
}

// ValidateGenesis performs genesis state validation for the oracle module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var gs types.GenesisState
	// unmarshal genesis-state
	if err := cdc.UnmarshalJSON(bz, &gs); err != nil {
		return err
	}

	// validate
	return gs.Validate()
}

// No RESTful routes exist for the oracle module (outside of those served via the grpc-gateway).
func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

// RegisterInvariants registers the invariants of the oracle module. If an invariant
// deviates from its predicted value, the InvariantRegistry triggers appropriate
// logic (most often the chain will be halted). No invariants exist for the oracle module.
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the genesis initialization for the x/oracle module. It determines the
// genesis state to initialize from via a json-encoded genesis-state. This method returns no validator set updates.
// This method panics on any errors
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) []abci.ValidatorUpdate {
	// unmarshal genesis-state (panic on errors)
	var gs types.GenesisState
	cdc.MustUnmarshalJSON(bz, &gs)

	// initialize genesis
	am.k.InitGenesis(ctx, gs)

	// return no validator-set updates
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the oracle module's exported genesis state as raw
// JSON bytes. This method panics on any error.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.k.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}
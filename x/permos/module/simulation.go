package permos

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"permos/testutil/sample"
	permossimulation "permos/x/permos/simulation"
	"permos/x/permos/types"
)

// avoid unused import issue
var (
	_ = permossimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePermos = "op_weight_msg_permos"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePermos int = 100

	opWeightMsgUpdatePermos = "op_weight_msg_permos"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePermos int = 100

	opWeightMsgDeletePermos = "op_weight_msg_permos"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePermos int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	permosGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PermosList: []types.Permos{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PermosCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&permosGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePermos int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePermos, &weightMsgCreatePermos, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePermos = defaultWeightMsgCreatePermos
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePermos,
		permossimulation.SimulateMsgCreatePermos(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePermos int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePermos, &weightMsgUpdatePermos, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePermos = defaultWeightMsgUpdatePermos
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePermos,
		permossimulation.SimulateMsgUpdatePermos(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePermos int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePermos, &weightMsgDeletePermos, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePermos = defaultWeightMsgDeletePermos
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePermos,
		permossimulation.SimulateMsgDeletePermos(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePermos,
			defaultWeightMsgCreatePermos,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				permossimulation.SimulateMsgCreatePermos(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePermos,
			defaultWeightMsgUpdatePermos,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				permossimulation.SimulateMsgUpdatePermos(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePermos,
			defaultWeightMsgDeletePermos,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				permossimulation.SimulateMsgDeletePermos(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

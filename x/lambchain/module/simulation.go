package lambchain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"lambchain/testutil/sample"
	lambchainsimulation "lambchain/x/lambchain/simulation"
	"lambchain/x/lambchain/types"
)

// avoid unused import issue
var (
	_ = lambchainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateSuma = "op_weight_msg_suma"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSuma int = 100

	opWeightMsgUpdateSuma = "op_weight_msg_suma"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSuma int = 100

	opWeightMsgDeleteSuma = "op_weight_msg_suma"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSuma int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	lambchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		SumaList: []types.Suma{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SumaCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&lambchainGenesis)
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

	var weightMsgCreateSuma int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSuma, &weightMsgCreateSuma, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSuma = defaultWeightMsgCreateSuma
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSuma,
		lambchainsimulation.SimulateMsgCreateSuma(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSuma int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSuma, &weightMsgUpdateSuma, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSuma = defaultWeightMsgUpdateSuma
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSuma,
		lambchainsimulation.SimulateMsgUpdateSuma(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSuma int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSuma, &weightMsgDeleteSuma, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSuma = defaultWeightMsgDeleteSuma
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSuma,
		lambchainsimulation.SimulateMsgDeleteSuma(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSuma,
			defaultWeightMsgCreateSuma,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				lambchainsimulation.SimulateMsgCreateSuma(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateSuma,
			defaultWeightMsgUpdateSuma,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				lambchainsimulation.SimulateMsgUpdateSuma(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteSuma,
			defaultWeightMsgDeleteSuma,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				lambchainsimulation.SimulateMsgDeleteSuma(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

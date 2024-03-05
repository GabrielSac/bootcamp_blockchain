package lambchain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "lambchain/api/lambchain/lambchain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "SumaAll",
					Use:       "list-suma",
					Short:     "List all suma",
				},
				{
					RpcMethod:      "Suma",
					Use:            "show-suma [id]",
					Short:          "Shows a suma by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateSuma",
					Use:            "create-suma [op1] [op2] [res]",
					Short:          "Create suma",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "op1"}, {ProtoField: "op2"}, {ProtoField: "res"}},
				},
				{
					RpcMethod:      "UpdateSuma",
					Use:            "update-suma [id] [op1] [op2] [res]",
					Short:          "Update suma",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "op1"}, {ProtoField: "op2"}, {ProtoField: "res"}},
				},
				{
					RpcMethod:      "DeleteSuma",
					Use:            "delete-suma [id]",
					Short:          "Delete suma",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

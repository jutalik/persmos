package permos

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "permos/api/permos/permos"
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
					RpcMethod: "PermosAll",
					Use:       "list-permos",
					Short:     "List all permos",
				},
				{
					RpcMethod:      "Permos",
					Use:            "show-permos [id]",
					Short:          "Shows a permos by id",
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
					RpcMethod:      "CreatePermos",
					Use:            "create-permos [mypermos]",
					Short:          "Create permos",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "mypermos"}},
				},
				{
					RpcMethod:      "UpdatePermos",
					Use:            "update-permos [id] [mypermos]",
					Short:          "Update permos",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "mypermos"}},
				},
				{
					RpcMethod:      "DeletePermos",
					Use:            "delete-permos [id]",
					Short:          "Delete permos",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

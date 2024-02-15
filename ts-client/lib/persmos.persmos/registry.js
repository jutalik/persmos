import { MsgUpdateParams } from "./types/persmos/persmos/tx";
import { QueryParamsResponse } from "./types/persmos/persmos/query";
import { GenesisState } from "./types/persmos/persmos/genesis";
import { Params } from "./types/persmos/persmos/params";
import { QueryParamsRequest } from "./types/persmos/persmos/query";
import { MsgUpdateParamsResponse } from "./types/persmos/persmos/tx";
const msgTypes = [
    ["/persmos.persmos.MsgUpdateParams", MsgUpdateParams],
    ["/persmos.persmos.QueryParamsResponse", QueryParamsResponse],
    ["/persmos.persmos.GenesisState", GenesisState],
    ["/persmos.persmos.Params", Params],
    ["/persmos.persmos.QueryParamsRequest", QueryParamsRequest],
    ["/persmos.persmos.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
];
export { msgTypes };

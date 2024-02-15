import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsRequest } from "./types/permos/permos/query";
import { QueryParamsResponse } from "./types/permos/permos/query";
import { MsgUpdateParamsResponse } from "./types/permos/permos/tx";
import { MsgUpdatePermosResponse } from "./types/permos/permos/tx";
import { QueryGetPermosResponse } from "./types/permos/permos/query";
import { MsgCreatePermosResponse } from "./types/permos/permos/tx";
import { MsgDeletePermosResponse } from "./types/permos/permos/tx";
import { Params } from "./types/permos/permos/params";
import { MsgUpdateParams } from "./types/permos/permos/tx";
import { Permos } from "./types/permos/permos/permos";
import { QueryAllPermosResponse } from "./types/permos/permos/query";
import { MsgUpdatePermos } from "./types/permos/permos/tx";
import { QueryGetPermosRequest } from "./types/permos/permos/query";
import { QueryAllPermosRequest } from "./types/permos/permos/query";
import { MsgCreatePermos } from "./types/permos/permos/tx";
import { MsgDeletePermos } from "./types/permos/permos/tx";
import { GenesisState } from "./types/permos/permos/genesis";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/permos.permos.QueryParamsRequest", QueryParamsRequest],
    ["/permos.permos.QueryParamsResponse", QueryParamsResponse],
    ["/permos.permos.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/permos.permos.MsgUpdatePermosResponse", MsgUpdatePermosResponse],
    ["/permos.permos.QueryGetPermosResponse", QueryGetPermosResponse],
    ["/permos.permos.MsgCreatePermosResponse", MsgCreatePermosResponse],
    ["/permos.permos.MsgDeletePermosResponse", MsgDeletePermosResponse],
    ["/permos.permos.Params", Params],
    ["/permos.permos.MsgUpdateParams", MsgUpdateParams],
    ["/permos.permos.Permos", Permos],
    ["/permos.permos.QueryAllPermosResponse", QueryAllPermosResponse],
    ["/permos.permos.MsgUpdatePermos", MsgUpdatePermos],
    ["/permos.permos.QueryGetPermosRequest", QueryGetPermosRequest],
    ["/permos.permos.QueryAllPermosRequest", QueryAllPermosRequest],
    ["/permos.permos.MsgCreatePermos", MsgCreatePermos],
    ["/permos.permos.MsgDeletePermos", MsgDeletePermos],
    ["/permos.permos.GenesisState", GenesisState],
    
];

export { msgTypes }
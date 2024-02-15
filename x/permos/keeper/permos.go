package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"permos/x/permos/types"
)

// GetPermosCount get the total number of permos
func (k Keeper) GetPermosCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PermosCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPermosCount set the total number of permos
func (k Keeper) SetPermosCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PermosCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPermos appends a permos in the store with a new id and update the count
func (k Keeper) AppendPermos(
	ctx context.Context,
	permos types.Permos,
) uint64 {
	// Create the permos
	count := k.GetPermosCount(ctx)

	// Set the ID of the appended value
	permos.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PermosKey))
	appendedValue := k.cdc.MustMarshal(&permos)
	store.Set(GetPermosIDBytes(permos.Id), appendedValue)

	// Update permos count
	k.SetPermosCount(ctx, count+1)

	return count
}

// SetPermos set a specific permos in the store
func (k Keeper) SetPermos(ctx context.Context, permos types.Permos) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PermosKey))
	b := k.cdc.MustMarshal(&permos)
	store.Set(GetPermosIDBytes(permos.Id), b)
}

// GetPermos returns a permos from its id
func (k Keeper) GetPermos(ctx context.Context, id uint64) (val types.Permos, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PermosKey))
	b := store.Get(GetPermosIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePermos removes a permos from the store
func (k Keeper) RemovePermos(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PermosKey))
	store.Delete(GetPermosIDBytes(id))
}

// GetAllPermos returns all permos
func (k Keeper) GetAllPermos(ctx context.Context) (list []types.Permos) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PermosKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Permos
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPermosIDBytes returns the byte representation of the ID
func GetPermosIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.PermosKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

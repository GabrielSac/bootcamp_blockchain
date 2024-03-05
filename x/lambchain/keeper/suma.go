package keeper

import (
	"context"
	"encoding/binary"

	"lambchain/x/lambchain/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetSumaCount get the total number of suma
func (k Keeper) GetSumaCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.SumaCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSumaCount set the total number of suma
func (k Keeper) SetSumaCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.SumaCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSuma appends a suma in the store with a new id and update the count
func (k Keeper) AppendSuma(
	ctx context.Context,
	suma types.Suma,
) uint64 {
	// Create the suma
	count := k.GetSumaCount(ctx)

	// Set the ID of the appended value
	suma.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SumaKey))
	appendedValue := k.cdc.MustMarshal(&suma)
	store.Set(GetSumaIDBytes(suma.Id), appendedValue)

	// Update suma count
	k.SetSumaCount(ctx, count+1)

	return count
}

// SetSuma set a specific suma in the store
func (k Keeper) SetSuma(ctx context.Context, suma types.Suma) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SumaKey))
	b := k.cdc.MustMarshal(&suma)
	store.Set(GetSumaIDBytes(suma.Id), b)
}

// GetSuma returns a suma from its id
func (k Keeper) GetSuma(ctx context.Context, id uint64) (val types.Suma, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SumaKey))
	b := store.Get(GetSumaIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSuma removes a suma from the store
func (k Keeper) RemoveSuma(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SumaKey))
	store.Delete(GetSumaIDBytes(id))
}

// GetAllSuma returns all suma
func (k Keeper) GetAllSuma(ctx context.Context) (list []types.Suma) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SumaKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Suma
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSumaIDBytes returns the byte representation of the ID
func GetSumaIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.SumaKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

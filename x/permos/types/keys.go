package types

const (
	// ModuleName defines the module name
	ModuleName = "permos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_permos"
)

var (
	ParamsKey = []byte("p_permos")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PermosKey      = "Permos/value/"
	PermosCountKey = "Permos/count/"
)

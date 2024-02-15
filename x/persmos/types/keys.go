package types

const (
	// ModuleName defines the module name
	ModuleName = "persmos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_persmos"
)

var (
	ParamsKey = []byte("p_persmos")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

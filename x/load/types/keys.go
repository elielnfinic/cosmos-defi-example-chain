package types

const (
	// ModuleName defines the module name
	ModuleName = "load"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_load"
)

var (
	ParamsKey = []byte("p_load")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

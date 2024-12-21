package types

const (
	// ModuleName defines the module name
	ModuleName = "omnichainwallet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_omnichainwallet"
)

var (
	ParamsKey = []byte("p_omnichainwallet")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

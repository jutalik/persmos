package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PermosList: []Permos{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in permos
	permosIdMap := make(map[uint64]bool)
	permosCount := gs.GetPermosCount()
	for _, elem := range gs.PermosList {
		if _, ok := permosIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for permos")
		}
		if elem.Id >= permosCount {
			return fmt.Errorf("permos id should be lower or equal than the last id")
		}
		permosIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

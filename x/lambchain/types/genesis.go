package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SumaList: []Suma{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in suma
	sumaIdMap := make(map[uint64]bool)
	sumaCount := gs.GetSumaCount()
	for _, elem := range gs.SumaList {
		if _, ok := sumaIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for suma")
		}
		if elem.Id >= sumaCount {
			return fmt.Errorf("suma id should be lower or equal than the last id")
		}
		sumaIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

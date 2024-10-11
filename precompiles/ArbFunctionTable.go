// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package precompiles

import (
	"errors"
	"math/big"
)

// ArbFunctionTable  precompile provided aggregator's the ability to manage function tables.
// Aggregation works differently in Nitro, so these methods have been stubbed and their effects disabled.
// They are kept for backwards compatibility.
type ArbFunctionTable struct {
	Address addr // 0x68
}

// TODO: add system test
// Upload does nothing
func (con ArbFunctionTable) Upload(c ctx, evm mech, buf []byte) error {
	return nil
}

// TODO: add system test
// Size returns the empty table's size, which is 0
func (con ArbFunctionTable) Size(c ctx, evm mech, addr addr) (huge, error) {
	return big.NewInt(0), nil
}

// TODO: add system test
// Get reverts since the table is empty
func (con ArbFunctionTable) Get(c ctx, evm mech, addr addr, index huge) (huge, bool, huge, error) {
	return nil, false, nil, errors.New("table is empty")
}

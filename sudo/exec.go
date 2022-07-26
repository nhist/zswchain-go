package sudo

import (
	eos "github.com/eoscanada/eos-go"
)

// NewExec creates an `exec` action, found in the `zswhq.wrap`
// contract.
//
// Given an `eos.Transaction`, call `eos.MarshalBinary` on it first,
// pass the resulting bytes as `eos.HexBytes` here.
func NewExec(executer eos.AccountName, transaction eos.Transaction) *eos.Action {
	a := &eos.Action{
		Account: eos.AccountName("zswhq.wrap"),
		Name:    eos.ActionName("exec"),
		Authorization: []eos.PermissionLevel{
			{Actor: executer, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(Exec{
			Executer:    executer,
			Transaction: transaction,
		}),
	}
	return a
}

// Exec represents the `zswhq.system::exec` action.
type Exec struct {
	Executer    eos.AccountName `json:"executer"`
	Transaction eos.Transaction `json:"trx"`
}

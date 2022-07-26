package system

import (
	eos "github.com/eoscanada/eos-go"
)

// NewRefund returns a `refund` action that lives on the
// `zswhq.system` contract.
func NewRefund(owner eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("zswhq"),
		Name:    ActN("refund"),
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `zswhq.system::refund` action
type Refund struct {
	Owner eos.AccountName `json:"owner"`
}

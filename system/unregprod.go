package system

import (
	eos "github.com/eoscanada/eos-go"
)

// NewUnregProducer returns a `unregprod` action that lives on the
// `zswhq.system` contract.
func NewUnregProducer(producer eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("zswhq"),
		Name:    ActN("unregprod"),
		Authorization: []eos.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(UnregProducer{
			Producer: producer,
		}),
	}
}

// UnregProducer represents the `zswhq.system::unregprod` action
type UnregProducer struct {
	Producer eos.AccountName `json:"producer"`
}

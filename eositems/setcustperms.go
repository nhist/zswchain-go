package eositems

import (
	eos "github.com/eoscanada/eos-go"
)

// NewSetCustodianPermissions is an action representing a setting a custodian's permissions to be broadcast
// through the chain network.
func NewSetCustodianPermissions(sender, custodian eos.AccountName, permissions eos.Uint128) *eos.Action {
	a := &eos.Action{
		Account: ZswItemsAN,
		Name:    ActN("setcustperms"),
		Authorization: []eos.PermissionLevel{
			{Actor: sender, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(SetCustodianPermissions{
			Sender:      sender,
			Custodian:   custodian,
			Permissions: permissions,
		}),
	}
	return a
}

// UnVote represents the `eos.items::unvote` action.
type SetCustodianPermissions struct {
	Sender      eos.AccountName `json:"sender"`
	Custodian   eos.AccountName `json:"custodian"`
	Permissions eos.Uint128     `json:"permissions"`
}

package eositems

import (
	eos "github.com/eoscanada/eos-go"
)

// NewSetUserPermissions is an action representing a setting a user's permissions to be broadcast
// through the chain network.
func NewSetUserPermissions(sender, user eos.AccountName, permissions eos.Uint128) *eos.Action {
	a := &eos.Action{
		Account: ZswItemsAN,
		Name:    ActN("setuserperms"),
		Authorization: []eos.PermissionLevel{
			{Actor: sender, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(SetUserPermissions{
			Sender:      sender,
			User:        user,
			Permissions: permissions,
		}),
	}
	return a
}

// UnVote represents the `eos.items::unvote` action.
type SetUserPermissions struct {
	Sender      eos.AccountName `json:"sender"`
	User        eos.AccountName `json:"user"`
	Permissions eos.Uint128     `json:"permissions"`
}

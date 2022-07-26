package eosperms

import (
	eos "github.com/eoscanada/eos-go"
)

func NewRemoveZswPerms(sender, user, scope eos.AccountName, permissions eos.Uint128) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.perms"),
		Name:    ActN("rmperms"),
		Authorization: []eos.PermissionLevel{
			{Actor: sender, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(RemoveZswPerms{
			Sender: sender,
			Scope: scope,
			User: user,
			Permissions: permissions,
		}),
	}
}


type RemoveZswPerms struct {
  Sender eos.AccountName `json:"sender"`
  Scope eos.AccountName `json:"scope"`
  User eos.AccountName `json:"user"`
  Permissions eos.Uint128 `json:"perm_bits"`
}
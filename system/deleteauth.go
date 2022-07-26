package system

import "github.com/eoscanada/eos-go"

// NewDeleteAuth creates an action from the `zswhq.system` contract
// called `deleteauth`.
//
// You cannot delete the `owner` or `active` permissions.  Also, if a
// permission is still linked through a previous `updatelink` action,
// you will need to `unlinkauth` first.
func NewDeleteAuth(account eos.AccountName, permission eos.PermissionName) *eos.Action {
	a := &eos.Action{
		Account: AN("zswhq"),
		Name:    ActN("deleteauth"),
		Authorization: []eos.PermissionLevel{
			{Actor: account, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(DeleteAuth{
			Account:    account,
			Permission: permission,
		}),
	}

	return a
}

// DeleteAuth represents the native `deleteauth` action, reachable
// through the `zswhq.system` contract.
type DeleteAuth struct {
	Account    eos.AccountName    `json:"account"`
	Permission eos.PermissionName `json:"permission"`
}

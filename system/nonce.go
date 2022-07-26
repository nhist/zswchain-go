package system

import "github.com/eoscanada/eos-go"

// NewNonce returns a `nonce` action that lives on the
// `zswhq.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `eos-bios` boot process by the
// `zswhq.system` contract.
func NewNonce(nonce string) *eos.Action {
	a := &eos.Action{
		Account:       AN("zswhq"),
		Name:          ActN("nonce"),
		Authorization: []eos.PermissionLevel{
			//{Actor: AN("zswhq"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}

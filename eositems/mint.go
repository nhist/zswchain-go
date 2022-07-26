package eositems

import (
	eos "github.com/eoscanada/eos-go"
)

// NewItemMint is an action representing a simple item minting to be broadcast
// through the chain network.

func NewItemMint(minter, to, toCustodian eos.AccountName, freezeTime uint32, itemIds []uint64, amounts []uint64, memo string) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mint"),
		Authorization: []eos.PermissionLevel{
			{Actor: minter, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(ItemMint{
			Minter:      minter,
			To:          to,
			ToCustodian: toCustodian,
			ItemIds:     itemIds,
			Amounts:     amounts,
			Memo:        memo,
			FreezeTime:  freezeTime,
		}),
	}
}

// ItemTransfer represents the `transfer` struct on `eos.items` contract.
type ItemMint struct {
	Minter      eos.AccountName `json:"minter"`
	To          eos.AccountName `json:"to"`
	ToCustodian eos.AccountName `json:"to_custodian"`
	ItemIds     []uint64        `json:"item_ids"`
	Amounts     []uint64        `json:"amounts"`
	Memo        string          `json:"memo"`
	FreezeTime  uint32          `json:"freeze_time"`
}

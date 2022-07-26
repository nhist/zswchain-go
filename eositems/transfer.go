package eositems

import (
	eos "github.com/eoscanada/eos-go"
)

// NewItemTransfer is an action representing a simple item transfer to be broadcast
// through the chain network.

func NewItemTransfer(authorizer, from, to, fromCustodian, toCustodian eos.AccountName, freezeTime uint32, useLiquidBalance bool, maxUnfreezeIterations uint32, itemIds []uint64, amounts []uint64, memo string) *eos.Action {
	var authorizers []eos.PermissionLevel
	if useLiquidBalance == true && fromCustodian != "nullnullnull" {
		authorizers = []eos.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: fromCustodian, Permission: PN("active")},
		}
		
	}else{
		authorizers = []eos.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
		}
	}
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("transfer"),
		Authorization: authorizers,
		ActionData: eos.NewActionData(ItemTransfer{
			Authorizer:            authorizer,
			From:                  from,
			To:                    to,
			FromCustodian:         fromCustodian,
			ToCustodian:           toCustodian,
			FreezeTime:            freezeTime,
			UseLiquidBalance:      useLiquidBalance,
			MaxUnfreezeIterations: maxUnfreezeIterations,
			ItemIds:               itemIds,
			Amounts:               amounts,
			Memo:                  memo,
		}),
	}
}

// ItemTransfer represents the `transfer` struct on `eos.items` contract.
type ItemTransfer struct {
	Authorizer            eos.AccountName `json:"authorizer"`
	From                  eos.AccountName `json:"from"`
	To                    eos.AccountName `json:"to"`
	FromCustodian         eos.AccountName `json:"from_custodian"`
	ToCustodian           eos.AccountName `json:"to_custodian"`
	FreezeTime            uint32          `json:"freeze_time"`
	UseLiquidBalance      bool            `json:"use_liquid_backup"`
	MaxUnfreezeIterations uint32          `json:"max_unfreeze_iterations"`
	ItemIds               []uint64        `json:"item_ids"`
	Amounts               []uint64        `json:"amounts"`
	Memo                  string          `json:"memo"`
}

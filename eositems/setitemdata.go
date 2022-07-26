package eositems

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/eosattr"
)

func NewSetItemData(authorizedEditor eos.AccountName, itemId uint64, newMutableData eosattr.AttributeMap) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("setitemdata"),
		Authorization: []eos.PermissionLevel{
			{Actor: authorizedEditor, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SetItemData{
			AuthorizedEditor: authorizedEditor,
			ItemId:           itemId,
			NewMutableData:   newMutableData,
		}),
	}
}

type SetItemData struct {
	AuthorizedEditor eos.AccountName      `json:"authorized_editor"`
	ItemId           uint64               `json:"item_id"`
	NewMutableData   eosattr.AttributeMap `json:"new_mutable_data"`
}

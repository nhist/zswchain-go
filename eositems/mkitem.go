package eositems

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/eosattr"
)

func NewMakeItem(authorizer eos.AccountName, authorizedMinter eos.AccountName, itemId uint64, eosId eos.Uint128, itemConfig uint32, itemTemplateId uint64, maxSupply uint64, schemaName eos.Name, immutableMetadata eosattr.AttributeMap, mutableMetadata eosattr.AttributeMap) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkitem"),
		Authorization: []eos.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: authorizedMinter, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(MakeItem{
			Authorizer:        authorizer,
			AuthorizedMinter:  authorizedMinter,
			ItemId:            itemId,
			ZswId:             eosId,
			ItemConfig:        itemConfig,
			ItemTemplateId:    itemTemplateId,
			MaxSupply:         maxSupply,
			SchemaName:        schemaName,
			ImmutableMetadata: immutableMetadata,
			MutableMetadata:   mutableMetadata,
		}),
	}
}

type MakeItem struct {
	Authorizer        eos.AccountName      `json:"authorizer"`
	AuthorizedMinter  eos.AccountName      `json:"authorized_minter"`
	ItemId            uint64               `json:"item_id"`
	ZswId             eos.Uint128          `json:"z5w_id"`
	ItemConfig        uint32               `json:"item_config"`
	ItemTemplateId    uint64               `json:"item_template_id"`
	MaxSupply         uint64               `json:"max_supply"`
	SchemaName        eos.Name             `json:"schema_name"`
	ImmutableMetadata eosattr.AttributeMap `json:"immutable_metadata"`
	MutableMetadata   eosattr.AttributeMap `json:"mutable_metadata"`
}

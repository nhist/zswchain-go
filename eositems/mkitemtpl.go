package eositems

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/eosattr"
)

func NewMakeItemTemplate(authorizer eos.AccountName, creator eos.AccountName, eosId eos.Uint128, itemTemplateId uint64, collectionId uint64, itemType uint32, schemaName eos.Name, immutableMetadata eosattr.AttributeMap, itemExternalMetadataUrlTemplate string) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkitemtpl"),
		Authorization: []eos.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(MakeItemTemplate{
			Authorizer:                      authorizer,
			Creator:                         creator,
			ZswId:                           eosId,
			ItemTemplateId:                  itemTemplateId,
			CollectionId:                    collectionId,
			ItemType:                        itemType,
			SchemaName:                      schemaName,
			ImmutableMetadata:               immutableMetadata,
			ItemExternalMetadataUrlTemplate: itemExternalMetadataUrlTemplate,
		}),
	}
}

type MakeItemTemplate struct {
	Authorizer                      eos.AccountName      `json:"authorizer"`
	Creator                         eos.AccountName      `json:"creator"`
	ZswId                           eos.Uint128          `json:"z5w_id"`
	ItemTemplateId                  uint64               `json:"item_template_id"`
	CollectionId                    uint64               `json:"collection_id"`
	ItemType                        uint32               `json:"item_type"`
	SchemaName                      eos.Name             `json:"schema_name"`
	ImmutableMetadata               eosattr.AttributeMap `json:"immutable_metadata"`
	ItemExternalMetadataUrlTemplate string               `json:"item_external_metadata_url_template"`
}

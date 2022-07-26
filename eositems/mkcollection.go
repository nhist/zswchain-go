package eositems

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/eosattr"
)

func NewMakeCollection(authorizer eos.AccountName, eosId eos.Uint128, collectionId uint64, collectionType uint32, creator eos.AccountName, issuingPlatform eos.AccountName, itemConfig uint32, secondaryMarketFee uint16, primaryMarketFee uint16, royaltyFeeCollector eos.AccountName, maxSupply uint64, maxItems uint64, maxSupplyPerItem uint64, schemaName eos.Name, authorizedMinters []eos.AccountName, notifyAccounts []eos.AccountName, authorizedMutableDataEditors []eos.AccountName, metadata eosattr.AttributeMap, externalMetadataUrl string) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkcollection"),
		Authorization: []eos.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
			{Actor: authorizer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(MakeCollection{
			Authorizer:                   authorizer,
			ZswId:                        eosId,
			CollectionId:                 collectionId,
			CollectionType:               collectionType,
			Creator:                      creator,
			IssuingPlatform:              issuingPlatform,
			ItemConfig:                   itemConfig,
			SecondaryMarketFee:           secondaryMarketFee,
			PrimaryMarketFee:             primaryMarketFee,
			RoyaltyFeeCollector:          royaltyFeeCollector,
			MaxSupply:                    maxSupply,
			MaxItems:                     maxItems,
			MaxSupplyPerItem:             maxSupplyPerItem,
			SchemaName:                   schemaName,
			AuthorizedMinters:            authorizedMinters,
			NotifyAccounts:               notifyAccounts,
			AuthorizedMutableDataEditors: authorizedMutableDataEditors,
			Metadata:                     metadata,
			ExternalMetadataUrl:          externalMetadataUrl,
		}),
	}
}

type MakeCollection struct {
	Authorizer                   eos.AccountName      `json:"authorizer"`
	ZswId                        eos.Uint128          `json:"z5w_id"`
	CollectionId                 uint64               `json:"collection_id"`
	CollectionType               uint32               `json:"collection_type"`
	Creator                      eos.AccountName      `json:"creator"`
	IssuingPlatform              eos.AccountName      `json:"issuing_platform"`
	ItemConfig                   uint32               `json:"item_config"`
	SecondaryMarketFee           uint16               `json:"secondary_market_fee"`
	PrimaryMarketFee             uint16               `json:"primary_market_fee"`
	RoyaltyFeeCollector          eos.AccountName      `json:"royalty_fee_collector"`
	MaxSupply                    uint64               `json:"max_supply"`
	MaxItems                     uint64               `json:"max_items"`
	MaxSupplyPerItem             uint64               `json:"max_supply_per_item"`
	SchemaName                   eos.Name             `json:"schema_name"`
	AuthorizedMinters            []eos.AccountName    `json:"authorized_minters"`
	NotifyAccounts               []eos.AccountName    `json:"notify_accounts"`
	AuthorizedMutableDataEditors []eos.AccountName    `json:"authorized_mutable_data_editors"`
	Metadata                     eosattr.AttributeMap `json:"metadata"`
	ExternalMetadataUrl          string               `json:"external_metadata_url"`
}

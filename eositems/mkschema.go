package eositems

import (
	eos "github.com/eoscanada/eos-go"
)

func NewMakeSchema(authorizer eos.AccountName, creator eos.AccountName, schemaName eos.Name, schemaFormat []eos.FieldDef) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkschema"),
		Authorization: []eos.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(MakeSchema{
			Authorizer:   authorizer,
			Creator:      creator,
			SchemaName:   schemaName,
			SchemaFormat: schemaFormat,
		}),
	}
}

type MakeSchema struct {
	Authorizer   eos.AccountName `json:"authorizer"`
	Creator      eos.AccountName `json:"creator"`
	SchemaName   eos.Name        `json:"schema_name"`
	SchemaFormat []eos.FieldDef  `json:"schema_format"`
}

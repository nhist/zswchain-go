package eositems

import (
	eos "github.com/eoscanada/eos-go"
)

func NewMakeIssuer(authorizer eos.AccountName, issuerName eos.AccountName, eosId eos.Uint128, altId eos.Uint128, permissions eos.Uint128, status uint32) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkissuer"),
		Authorization: []eos.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(MakeIssuer{
			Authorizer: authorizer,
			IssuerName: issuerName,
			ZswId: eosId,
			AltId: altId,
			Permissions: permissions,
			Status: status,
		}),
	}
}


type MakeIssuer struct {
  Authorizer eos.AccountName `json:"authorizer"`
  IssuerName eos.AccountName `json:"issuer_name"`
  ZswId eos.Uint128 `json:"z5w_id"`
  AltId eos.Uint128 `json:"alt_id"`
  Permissions eos.Uint128 `json:"permissions"`
  Status uint32 `json:"status"`
}
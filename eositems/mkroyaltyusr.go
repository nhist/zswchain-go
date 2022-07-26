package eositems

import (
	eos "github.com/eoscanada/eos-go"
)

func NewMakeRoyaltyUser(authorizer eos.AccountName, newRoyaltyUser eos.AccountName, eosId eos.Uint128, altId eos.Uint128, status uint32) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkroyaltyusr"),
		Authorization: []eos.PermissionLevel{
			{Actor: authorizer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(MakeRoyaltyUser{
			Authorizer: authorizer,
			NewRoyaltyUser: newRoyaltyUser,
			ZswId: eosId,
			AltId: altId,
			Status: status,
		}),
	}
}


type MakeRoyaltyUser struct {
  Authorizer eos.AccountName `json:"authorizer"`
  NewRoyaltyUser eos.AccountName `json:"newroyaltyusr"`
  ZswId eos.Uint128 `json:"z5w_id"`
  AltId eos.Uint128 `json:"alt_id"`
  Status uint32 `json:"status"`
}
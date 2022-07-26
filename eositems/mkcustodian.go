package eositems

import (
	eos "github.com/eoscanada/eos-go"
)

func NewMakeCustodian(creator eos.AccountName, custodianName eos.AccountName, eosId eos.Uint128, altId eos.Uint128, permissions eos.Uint128, status uint32, incomingFreezePeriod uint32, notifyAccounts []eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("zsw.items"),
		Name:    ActN("mkcustodian"),
		Authorization: []eos.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(MakeCustodian{
			Creator: creator,
			CustodianName: custodianName,
			ZswId: eosId,
			AltId: altId,
			Permissions: permissions,
			Status: status,
			IncomingFreezePeriod: incomingFreezePeriod,
			NotifyAccounts: notifyAccounts,
		}),
	}
}


type MakeCustodian struct {
  Creator eos.AccountName `json:"creator"`
  CustodianName eos.AccountName `json:"custodian_name"`
  ZswId eos.Uint128 `json:"z5w_id"`
  AltId eos.Uint128 `json:"alt_id"`
  Permissions eos.Uint128 `json:"permissions"`
  Status uint32 `json:"status"`
  IncomingFreezePeriod uint32 `json:"incoming_freeze_period"`
  NotifyAccounts []eos.AccountName `json:"notify_accounts"`
}
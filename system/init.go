package system

import (
	"github.com/eoscanada/eos-go"
)

func init() {
	eos.RegisterAction(AN("zswhq"), ActN("setcode"), SetCode{})
	eos.RegisterAction(AN("zswhq"), ActN("setabi"), SetABI{})
	eos.RegisterAction(AN("zswhq"), ActN("newaccount"), NewAccount{})
	eos.RegisterAction(AN("zswhq"), ActN("delegatebw"), DelegateBW{})
	eos.RegisterAction(AN("zswhq"), ActN("undelegatebw"), UndelegateBW{})
	eos.RegisterAction(AN("zswhq"), ActN("refund"), Refund{})
	eos.RegisterAction(AN("zswhq"), ActN("regproducer"), RegProducer{})
	eos.RegisterAction(AN("zswhq"), ActN("unregprod"), UnregProducer{})
	eos.RegisterAction(AN("zswhq"), ActN("regproxy"), RegProxy{})
	eos.RegisterAction(AN("zswhq"), ActN("voteproducer"), VoteProducer{})
	eos.RegisterAction(AN("zswhq"), ActN("claimrewards"), ClaimRewards{})
	eos.RegisterAction(AN("zswhq"), ActN("buyram"), BuyRAM{})
	eos.RegisterAction(AN("zswhq"), ActN("buyrambytes"), BuyRAMBytes{})
	eos.RegisterAction(AN("zswhq"), ActN("linkauth"), LinkAuth{})
	eos.RegisterAction(AN("zswhq"), ActN("unlinkauth"), UnlinkAuth{})
	eos.RegisterAction(AN("zswhq"), ActN("deleteauth"), DeleteAuth{})
	eos.RegisterAction(AN("zswhq"), ActN("rmvproducer"), RemoveProducer{})
	eos.RegisterAction(AN("zswhq"), ActN("setprods"), SetProds{})
	eos.RegisterAction(AN("zswhq"), ActN("setpriv"), SetPriv{})
	eos.RegisterAction(AN("zswhq"), ActN("canceldelay"), CancelDelay{})
	eos.RegisterAction(AN("zswhq"), ActN("bidname"), Bidname{})
	// eos.RegisterAction(AN("zswhq"), ActN("nonce"), &Nonce{})
	eos.RegisterAction(AN("zswhq"), ActN("sellram"), SellRAM{})
	eos.RegisterAction(AN("zswhq"), ActN("updateauth"), UpdateAuth{})
	eos.RegisterAction(AN("zswhq"), ActN("setramrate"), SetRAMRate{})
	eos.RegisterAction(AN("zswhq"), ActN("setalimits"), Setalimits{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN

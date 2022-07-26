package token

import "github.com/eoscanada/eos-go"

func init() {
	eos.RegisterAction(AN("zswhq.token"), ActN("transfer"), Transfer{})
	eos.RegisterAction(AN("zswhq.token"), ActN("issue"), Issue{})
	eos.RegisterAction(AN("zswhq.token"), ActN("create"), Create{})
}

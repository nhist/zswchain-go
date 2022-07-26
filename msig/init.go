package msig

import (
	"github.com/eoscanada/eos-go"
)

func init() {
	eos.RegisterAction(AN("zswhq.msig"), ActN("propose"), &Propose{})
	eos.RegisterAction(AN("zswhq.msig"), ActN("approve"), &Approve{})
	eos.RegisterAction(AN("zswhq.msig"), ActN("unapprove"), &Unapprove{})
	eos.RegisterAction(AN("zswhq.msig"), ActN("cancel"), &Cancel{})
	eos.RegisterAction(AN("zswhq.msig"), ActN("exec"), &Exec{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN

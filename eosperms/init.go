package eosperms

import eos "github.com/eoscanada/eos-go"

func init() {
	eos.RegisterAction(ZswPermsAN, ActN("addperms"), AddZswPerms{})
	eos.RegisterAction(ZswPermsAN, ActN("rmperms"), RemoveZswPerms{})
	eos.RegisterAction(ZswPermsAN, ActN("setperms"), SetZswPerms{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN

var ZswPermsAN = AN("zsw.perms")

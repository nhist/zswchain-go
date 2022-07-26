package eositems

import eos "github.com/eoscanada/eos-go"

func init() {
	eos.RegisterAction(ZswItemsAN, ActN("mint"), ItemMint{})
	eos.RegisterAction(ZswItemsAN, ActN("setcustperms"), SetCustodianPermissions{})
	eos.RegisterAction(ZswItemsAN, ActN("setuserperms"), SetUserPermissions{})
	eos.RegisterAction(ZswItemsAN, ActN("mkcollection"), MakeCollection{})
	eos.RegisterAction(ZswItemsAN, ActN("mkcustodian"), MakeCustodian{})
	eos.RegisterAction(ZswItemsAN, ActN("mkissuer"), MakeIssuer{})
	eos.RegisterAction(ZswItemsAN, ActN("mkitem"), MakeItem{})
	eos.RegisterAction(ZswItemsAN, ActN("mkschema"), MakeSchema{})
	eos.RegisterAction(ZswItemsAN, ActN("mkitemtpl"), MakeItemTemplate{})
	eos.RegisterAction(ZswItemsAN, ActN("mkroyaltyusr"), MakeRoyaltyUser{})
	eos.RegisterAction(ZswItemsAN, ActN("setcustperms"), SetCustodianPermissions{})
	eos.RegisterAction(ZswItemsAN, ActN("setuserperms"), SetUserPermissions{})
	eos.RegisterAction(ZswItemsAN, ActN("transfer"), ItemTransfer{})
	eos.RegisterAction(ZswItemsAN, ActN("setitemdata"), SetItemData{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN

var ZswItemsAN = AN("zsw.items")

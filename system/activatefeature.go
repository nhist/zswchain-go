package system

import (
	"github.com/eoscanada/eos-go"
)

func NewActivateFeature(featureDigest eos.Checksum256) *eos.Action {
	return &eos.Action{
		Account: AN("zswhq"),
		Name:    ActN("activate"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("zswhq"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Activate{
			FeatureDigest: featureDigest,
		}),
	}
}

// Activate represents a `activate` action on the `zswhq` contract.
type Activate struct {
	FeatureDigest eos.Checksum256 `json:"feature_digest"`
}

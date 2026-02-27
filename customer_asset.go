package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type CustomerAsset struct {
	*resources.CustomerAsset
	Asset *Asset
}

func NewCustomerAsset() *CustomerAsset {
	return &CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{},
		Asset:         NewAsset(),
	}
}

type CustomerAssets []*CustomerAsset

func NewCustomerAssets() CustomerAssets {
	return make(CustomerAssets, 0)
}

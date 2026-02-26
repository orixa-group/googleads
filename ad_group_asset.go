package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type AdGroupAsset struct {
	*resources.AdGroupAsset
	Asset *Asset
}

func NewAdGroupAsset() *AdGroupAsset {
	return &AdGroupAsset{
		AdGroupAsset: &resources.AdGroupAsset{},
		Asset:        NewAsset(),
	}
}

type AdGroupAssets []*AdGroupAsset

func NewAdGroupAssets() AdGroupAssets {
	return make(AdGroupAssets, 0)
}

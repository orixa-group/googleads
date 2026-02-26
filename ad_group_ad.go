package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type AdGroupAd struct {
	*resources.AdGroupAd
	Ad *Ad
}

func NewAdGroupAd() *AdGroupAd {
	return &AdGroupAd{
		AdGroupAd: &resources.AdGroupAd{},
		Ad:        NewAd(),
	}
}

type AdGroupAds []*AdGroupAd

func NewAdGroupAds() AdGroupAds {
	return make(AdGroupAds, 0)
}

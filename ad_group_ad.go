package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AdGroupAd struct {
	*resources.AdGroupAd
}

func NewAdGroupAd() *AdGroupAd {
	return &AdGroupAd{&resources.AdGroupAd{}}
}

func (aga *AdGroupAd) createOperation(adGroup *AdGroup) *services.MutateOperation {
	aga.AdGroup = String(adGroup.GetResourceName())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_AdGroupAdOperation{
			AdGroupAdOperation: &services.AdGroupAdOperation{
				Operation: &services.AdGroupAdOperation_Create{
					Create: aga.AdGroupAd,
				},
			},
		},
	}
}

type AdGroupAds []*AdGroupAd

func NewAdGroupAds() AdGroupAds {
	return make(AdGroupAds, 0)
}

func (agas AdGroupAds) createOperations(adGroup *AdGroup) []*services.MutateOperation {
	return Map(agas, func(item *AdGroupAd) *services.MutateOperation {
		return item.createOperation(adGroup)
	})
}

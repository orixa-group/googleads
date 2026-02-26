package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

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

func (aga *AdGroupAsset) createOperations(adGroup *AdGroup, tempId tempIdGenerator) []*services.MutateOperation {
	aop := aga.Asset.createOperation(adGroup.Campaign.Customer, tempId)

	aga.AdGroupAsset.AdGroup = adGroup.GetResourceName()
	aga.AdGroupAsset.Asset = aga.Asset.GetResourceName()

	return []*services.MutateOperation{
		aop,
		{
			Operation: &services.MutateOperation_AdGroupAssetOperation{
				AdGroupAssetOperation: &services.AdGroupAssetOperation{
					Operation: &services.AdGroupAssetOperation_Create{
						Create: aga.AdGroupAsset,
					},
				},
			},
		},
	}
}

type AdGroupAssets []*AdGroupAsset

func NewAdGroupAssets() AdGroupAssets {
	return make(AdGroupAssets, 0)
}

func (agas AdGroupAssets) createOperations(adGroup *AdGroup, tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(agas, func(item *AdGroupAsset) []*services.MutateOperation {
		return item.createOperations(adGroup, tempId)
	}))
}

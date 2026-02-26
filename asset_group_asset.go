package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AssetGroupAsset struct {
	*resources.AssetGroupAsset
	Asset *Asset
}

func NewAssetGroupAsset() *AssetGroupAsset {
	return &AssetGroupAsset{
		AssetGroupAsset: &resources.AssetGroupAsset{},
		Asset:           NewAsset(),
	}
}

func (aga *AssetGroupAsset) createOperations(assetGroup *AssetGroup, tempId tempIdGenerator) []*services.MutateOperation {
	aop := aga.Asset.createOperation(assetGroup.Campaign.Customer, tempId)

	aga.AssetGroupAsset.AssetGroup = assetGroup.GetResourceName()
	aga.AssetGroupAsset.Asset = aga.Asset.GetResourceName()

	return []*services.MutateOperation{
		aop,
		{
			Operation: &services.MutateOperation_AssetGroupAssetOperation{
				AssetGroupAssetOperation: &services.AssetGroupAssetOperation{
					Operation: &services.AssetGroupAssetOperation_Create{
						Create: aga.AssetGroupAsset,
					},
				},
			},
		},
	}
}

type AssetGroupAssets []*AssetGroupAsset

func NewAssetGroupAssets() AssetGroupAssets {
	return make(AssetGroupAssets, 0)
}

func (aga AssetGroupAssets) createOperations(assetGroup *AssetGroup, tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(aga, func(item *AssetGroupAsset) []*services.MutateOperation {
		return item.createOperations(assetGroup, tempId)
	}))
}

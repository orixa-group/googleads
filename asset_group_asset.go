package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/resources"
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

type AssetGroupAssets []*AssetGroupAsset

func NewAssetGroupAssets() AssetGroupAssets {
	return make(AssetGroupAssets, 0)
}

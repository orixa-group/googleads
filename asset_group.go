package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type AssetGroup struct {
	*resources.AssetGroup
	Campaign *Campaign
	Assets   AssetGroupAssets
}

func NewAssetGroup() *AssetGroup {
	return &AssetGroup{
		AssetGroup: &resources.AssetGroup{},
		Campaign:   NewCampaign(),
		Assets:     NewAssetGroupAssets(),
	}
}

type AssetGroups []*AssetGroup

func NewAssetGroups() AssetGroups {
	return make(AssetGroups, 0)
}

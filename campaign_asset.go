package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type CampaignAsset struct {
	*resources.CampaignAsset
	Asset *Asset
}

func NewCampaignAsset() *CampaignAsset {
	return &CampaignAsset{
		CampaignAsset: &resources.CampaignAsset{},
		Asset:         NewAsset(),
	}
}

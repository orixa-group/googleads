package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type AdGroup struct {
	*resources.AdGroup
	Campaign *Campaign
	Criteria AdGroupCriteria
	Assets   AdGroupAssets
	Ads      AdGroupAds
}

func NewAdGroup() *AdGroup {
	return &AdGroup{
		AdGroup:  &resources.AdGroup{},
		Campaign: NewCampaign(),
		Criteria: NewAdGroupCriteria(),
		Assets:   NewAdGroupAssets(),
		Ads:      NewAdGroupAds(),
	}
}

type AdGroups []*AdGroup

func NewAdGroups() AdGroups {
	return make(AdGroups, 0)
}

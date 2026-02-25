package googleads

import "fmt"

type CampaignAssetFilter Filter

func CampaignAssetByResourceName(resourceName string) CampaignAssetFilter {
	return func() string {
		return fmt.Sprintf("campaign_asset.resource_name = '%s'", resourceName)
	}
}

func CampaignAssetByAsset(resourceName string) CampaignAssetFilter {
	return func() string {
		return fmt.Sprintf("campaign_asset.asset = '%s'", resourceName)
	}
}

func CampaignAssetByCampaign(resourceName string) CampaignAssetFilter {
	return func() string {
		return fmt.Sprintf("campaign_asset.campaign = '%s'", resourceName)
	}
}

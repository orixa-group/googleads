package googleads

import "fmt"

type AdGroupAssetFilter Filter

func AdGroupAssetByAdGroup(resourceName string) AdGroupAssetFilter {
	return func() string {
		return fmt.Sprintf("ad_group_asset.ad_group = '%s'", resourceName)
	}
}

func AdGroupAssetByAsset(resourceName string) AdGroupAssetFilter {
	return func() string {
		return fmt.Sprintf("ad_group_asset.asset = '%s'", resourceName)
	}
}

func AdGroupAssetByResourceName(resourceName string) AdGroupAssetFilter {
	return func() string {
		return fmt.Sprintf("ad_group_asset.resource_name = '%s'", resourceName)
	}
}

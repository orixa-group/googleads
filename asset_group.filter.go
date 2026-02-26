package googleads

import "fmt"

type AssetGroupFilter Filter

func AssetGroupByCampaign(resourceName string) AssetGroupFilter {
	return func() string {
		return fmt.Sprintf("asset_group.campaign = '%s'", resourceName)
	}
}

func AssetGroupById(id string) AssetGroupFilter {
	return func() string {
		return fmt.Sprintf("asset_group.id = '%s'", id)
	}
}

func AssetGroupByResourceName(resourceName string) AssetGroupFilter {
	return func() string {
		return fmt.Sprintf("asset_group.resource_name = '%s'", resourceName)
	}
}

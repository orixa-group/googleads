package googleads

import "fmt"

type AssetGroupAssetFilter Filter

func AssetGroupAssetByAssetGroup(resourceName string) AssetGroupAssetFilter {
	return func() string {
		return fmt.Sprintf("asset_group_asset.asset_group = '%s'", resourceName)
	}
}

func AssetGroupAssetByAsset(resourceName string) AssetGroupAssetFilter {
	return func() string {
		return fmt.Sprintf("asset_group_asset.asset = '%s'", resourceName)
	}
}

func AssetGroupAssetByResourceName(resourceName string) AssetGroupAssetFilter {
	return func() string {
		return fmt.Sprintf("asset_group_asset.resource_name = '%s'", resourceName)
	}
}

package googleads

import "fmt"

type AssetFilter Filter

func AssetById(id string) AssetFilter {
	return func() string {
		return fmt.Sprintf("asset.id = '%s'", id)
	}
}

func AssetByResourceName(resourceName string) AssetFilter {
	return func() string {
		return fmt.Sprintf("asset.resource_name = '%s'", resourceName)
	}
}

package googleads

import "fmt"

type CustomerAssetFilter func() string

func CustomerAssetByResourceName(resourceName string) CustomerAssetFilter {
	return func() string {
		return fmt.Sprintf("customer_asset.resource_name = '%s'", resourceName)
	}
}

func CustomerAssetByAsset(resourceName string) CustomerAssetFilter {
	return func() string {
		return fmt.Sprintf("customer_asset.asset = '%s'", resourceName)
	}
}

func CustomerAssetByCustomer(resourceName string) CustomerAssetFilter {
	return func() string {
		return fmt.Sprintf("customer_asset.customer = '%s'", resourceName)
	}
}

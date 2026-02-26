package googleads

import "fmt"

type AdGroupAdFilter func() string

func AdGroupAdByAdGroup(resourceName string) AdGroupAdFilter {
	return func() string {
		return fmt.Sprintf("ad_group_ad.ad_group = '%s'", resourceName)
	}
}

func AdGroupAdByResourceName(resourceName string) AdGroupAdFilter {
	return func() string {
		return fmt.Sprintf("ad_group_ad.resource_name = '%s'", resourceName)
	}
}

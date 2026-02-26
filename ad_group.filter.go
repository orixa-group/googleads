package googleads

import "fmt"

type AdGroupFilter Filter

func AdGroupsByCampaign(resourceName string) AdGroupFilter {
	return func() string {
		return fmt.Sprintf("campaign.resource_name = '%s'", resourceName)
	}
}

func AdGroupById(id string) AdGroupFilter {
	return func() string {
		return fmt.Sprintf("ad_group.id = '%s'", id)
	}
}

func AdGroupByResourceName(resourceName string) AdGroupFilter {
	return func() string {
		return fmt.Sprintf("ad_group.resource_name = '%s'", resourceName)
	}
}

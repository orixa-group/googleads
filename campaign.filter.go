package googleads

import "fmt"

type CampaignFilter Filter

func CampaignById(id string) CampaignFilter {
	return func() string {
		return fmt.Sprintf("campaign.id = '%s'", id)
	}
}

func CampaignByResourceName(resourceName string) CampaignFilter {
	return func() string {
		return fmt.Sprintf("campaign.resource_name = '%s'", resourceName)
	}
}

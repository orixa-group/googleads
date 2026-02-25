package googleads

import "fmt"

type CampaignCriterionFilter Filter

func CampaignCriterionByResourceName(resourceName string) CampaignCriterionFilter {
	return func() string {
		return fmt.Sprintf("campaign_criterion.resource_name = '%s'", resourceName)
	}
}

func CampaignCriterionByCampaign(resourceName string) CampaignCriterionFilter {
	return func() string {
		return fmt.Sprintf("campaign_criterion.campaign = '%s'", resourceName)
	}
}

package googleads

import "fmt"

type CampaignBudgetFilter Filter

func CampaignBudgetById(id string) CampaignBudgetFilter {
	return func() string {
		return fmt.Sprintf("campaign_budget.id = '%s'", id)
	}
}

func CampaignBudgetByResourceName(resourceName string) CampaignBudgetFilter {
	return func() string {
		return fmt.Sprintf("campaign_budget.resource_name = '%s'", resourceName)
	}
}

func CampaignBudgetByCustomer(id string) CampaignBudgetFilter {
	return func() string {
		return fmt.Sprintf("customer.id = '%s'", id)
	}
}

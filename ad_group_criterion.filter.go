package googleads

import "fmt"

type AdGroupCriterionFilter func() string

func AdGroupCriterionByResourceName(resourceName string) AdGroupCriterionFilter {
	return func() string {
		return fmt.Sprintf("ad_group_criterion.resource_name = '%s'", resourceName)
	}
}

func AdGroupCriterionByAdGroup(resourceName string) AdGroupCriterionFilter {
	return func() string {
		return fmt.Sprintf("ad_group_criterion.ad_group = '%s'", resourceName)
	}
}

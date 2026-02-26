package googleads

import "fmt"

type AccountBudgetFilter func() string

func AccountBudgetById(id string) AccountBudgetFilter {
	return func() string {
		return fmt.Sprintf("account_budget_proposal.id = '%s'", id)
	}
}

func AccountBudgetByResourceName(resourceName string) AccountBudgetFilter {
	return func() string {
		return fmt.Sprintf("account_budget_proposal.resource_name = '%s'", resourceName)
	}
}

func AccountBudgetByCustomer(id string) AccountBudgetFilter {
	return func() string {
		return fmt.Sprintf("customer.id = '%s'", id)
	}
}

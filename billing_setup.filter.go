package googleads

import "fmt"

type BillingSetupFilter func() string

func BillingSetupById(id string) BillingSetupFilter {
	return func() string {
		return fmt.Sprintf("billing_setup.id = '%s'", id)
	}
}

func BillingSetupByResourceName(resourceName string) BillingSetupFilter {
	return func() string {
		return fmt.Sprintf("billing_setup.resource_name = '%s'", resourceName)
	}
}

func BillingSetupByCustomer(id string) BillingSetupFilter {
	return func() string {
		return fmt.Sprintf("customer.id = '%s'", id)
	}
}

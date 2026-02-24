package googleads

import "fmt"

type CustomerFilter func() string

func CustomerById(id string) CustomerFilter {
	return func() string {
		return fmt.Sprintf("customer.id = '%s'", id)
	}
}

func CustomerByResourceName(resourceName string) CustomerFilter {
	return func() string {
		return fmt.Sprintf("customer.resource_name = '%s'", resourceName)
	}
}

package googleads

import "fmt"

type AdFilter func() string

func AdById(id string) AdFilter {
	return func() string {
		return fmt.Sprintf("ad.id = '%s'", id)
	}
}

func AdByResourceName(resourceName string) AdFilter {
	return func() string {
		return fmt.Sprintf("ad.resource_name = '%s'", resourceName)
	}
}

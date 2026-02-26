package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type BillingSetup struct {
	*resources.BillingSetup
}

func NewBillingSetup() *BillingSetup {
	return &BillingSetup{&resources.BillingSetup{}}
}

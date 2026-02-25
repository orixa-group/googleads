package googleads

import (
	"fmt"

	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type Asset struct {
	*resources.Asset
}

func NewAsset() *Asset {
	return &Asset{&resources.Asset{}}
}

func (a *Asset) createOperation(customer *Customer, tempId tempIdGenerator) *services.MutateOperation {
	a.ResourceName = fmt.Sprintf("customers/%s/assets/%s", customer.GetId(), tempId())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_AssetOperation{
			AssetOperation: &services.AssetOperation{
				Operation: &services.AssetOperation_Create{
					Create: a.Asset,
				},
			},
		},
	}
}

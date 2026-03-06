package googleads

import (
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type Asset struct {
	*resources.Asset
}

func (a Asset) GetId() string {
	return strconv.Itoa(int(a.Asset.GetId()))
}

func (a *Asset) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	a.Asset.Id = Int64(i)
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

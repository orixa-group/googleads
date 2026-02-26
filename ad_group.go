package googleads

import (
	"context"
	"fmt"

	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AdGroup struct {
	*resources.AdGroup
	Campaign *Campaign
	Criteria AdGroupCriteria
	Assets   AdGroupAssets
	Ads      AdGroupAds
}

func NewAdGroup() *AdGroup {
	return &AdGroup{
		AdGroup:  &resources.AdGroup{},
		Campaign: NewCampaign(),
		Criteria: NewAdGroupCriteria(),
		Assets:   NewAdGroupAssets(),
		Ads:      NewAdGroupAds(),
	}
}

func (ag *AdGroup) createOperation(tempId tempIdGenerator) *services.MutateOperation {
	ag.ResourceName = fmt.Sprintf("customers/%s/adGroups/%s", ag.Campaign.Customer.GetId(), tempId())
	ag.AdGroup.Campaign = String(ag.Campaign.GetResourceName())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_AdGroupOperation{
			AdGroupOperation: &services.AdGroupOperation{
				Operation: &services.AdGroupOperation_Create{
					Create: ag.AdGroup,
				},
			},
		},
	}
}

func (ag *AdGroup) Create(ctx context.Context) error {
	tempId := newTempIdGenerator()

	ops := make([]*services.MutateOperation, 0)

	ops = append(ops, ag.createOperation(tempId))
	ops = append(ops, ag.Criteria.createOperations(ag)...)
	ops = append(ops, ag.Assets.createOperations(ag, tempId)...)
	ops = append(ops, ag.Ads.createOperations(ag)...)

	_, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       ag.Campaign.Customer.GetId(),
		MutateOperations: ops,
	})
	return err
}

type AdGroups []*AdGroup

func NewAdGroups() AdGroups {
	return make(AdGroups, 0)
}

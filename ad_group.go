package googleads

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/enums"
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

func (ag AdGroup) GetId() string {
	return strconv.Itoa(int(ag.AdGroup.GetId()))
}

func (ag *AdGroup) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	ag.AdGroup.Id = Int64(i)
}

func (ag *AdGroup) SetName(name string) {
	ag.AdGroup.Name = String(name)
}

func (ag AdGroup) GetEnabled() bool {
	return ag.AdGroup.GetStatus() == enums.AdGroupStatusEnum_ENABLED
}

func (ag *AdGroup) SetEnabled(enabled bool) {
	if enabled {
		ag.AdGroup.Status = enums.AdGroupStatusEnum_ENABLED
	} else {
		ag.AdGroup.Status = enums.AdGroupStatusEnum_PAUSED
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

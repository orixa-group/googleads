package googleads

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type AdGroup struct {
	*resources.AdGroup
	Campaign *Campaign
	Criteria AdGroupCriteria
	Assets   AdGroupAssets
	Ads      AdGroupAds
	resource
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
	ag.addUpdatedField("name")
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
	ag.addUpdatedField("status")
}

func (ag *AdGroup) ListCriteria(ctx context.Context) (AdGroupCriteria, error) {
	return ListAdGroupCriteria(ctx, ag.Campaign.Customer.GetId(), AdGroupCriterionByAdGroup(ag.GetResourceName()))
}

func (ag *AdGroup) ListAssets(ctx context.Context) (AdGroupAssets, error) {
	return ListAdGroupAssets(ctx, ag.Campaign.Customer.GetId(), AdGroupAssetByAdGroup(ag.GetResourceName()))
}

func (ag *AdGroup) ListAds(ctx context.Context) (AdGroupAds, error) {
	return ListAdGroupAds(ctx, ag.Campaign.Customer.GetId(), AdGroupAdByAdGroup(ag.GetResourceName()))
}

func (ag *AdGroup) Save(ctx context.Context) error {
	if ag.isNew(ag.GetId()) {
		return ag.Create(ctx)
	}
	return ag.Update(ctx)
}

func (ag *AdGroup) Update(ctx context.Context) error {
	paths := ag.getUpdatedFields()
	if len(paths) == 0 {
		return nil
	}
	_, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       ag.Campaign.Customer.GetId(),
		MutateOperations: []*services.MutateOperation{ag.updateOperation(paths)},
	})
	return err
}

func (ag *AdGroup) updateOperation(paths []string) *services.MutateOperation {
	return &services.MutateOperation{
		Operation: &services.MutateOperation_AdGroupOperation{
			AdGroupOperation: &services.AdGroupOperation{
				Operation: &services.AdGroupOperation_Update{
					Update: ag.AdGroup,
				},
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: paths,
				},
			},
		},
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

func (ag *AdGroup) createOperations(tempId tempIdGenerator) []*services.MutateOperation {
	ops := make([]*services.MutateOperation, 0)

	ops = append(ops, ag.createOperation(tempId))
	ops = append(ops, ag.Criteria.createOperations(ag)...)
	ops = append(ops, ag.Assets.createOperations(ag, tempId)...)
	ops = append(ops, ag.Ads.createOperations(ag)...)

	return ops
}

func (ag *AdGroup) Create(ctx context.Context) error {
	resp, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       ag.Campaign.Customer.GetId(),
		MutateOperations: ag.createOperations(newTempIdGenerator()),
	})
	if err != nil {
		return err
	}

	res := resp.GetMutateOperationResponses()[0].GetAdGroupResult()
	ag.ResourceName = res.GetResourceName()
	ag.SetId(strings.Split(ag.ResourceName, "/")[3])

	return nil
}

type AdGroups []*AdGroup

func (ags AdGroups) createOperations(tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(ags, func(item *AdGroup) []*services.MutateOperation {
		return item.createOperations(tempId)
	}))
}

package googleads

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AssetGroup struct {
	*resources.AssetGroup
	Campaign *Campaign
	Assets   AssetGroupAssets
}

func (ag AssetGroup) GetId() string {
	return strconv.Itoa(int(ag.AssetGroup.GetId()))
}

func (ag *AssetGroup) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	ag.AssetGroup.Id = i
}

func (ag *AssetGroup) SetName(name string) {
	ag.AssetGroup.Name = name
}

func (ag *AssetGroup) SetFinalUrls(urls []string) {
	ag.AssetGroup.FinalUrls = urls
}

func (ag AssetGroup) GetEnabled() bool {
	return ag.AssetGroup.GetStatus() == enums.AssetGroupStatusEnum_ENABLED
}

func (ag *AssetGroup) SetEnabled(enabled bool) {
	if enabled {
		ag.AssetGroup.Status = enums.AssetGroupStatusEnum_ENABLED
	} else {
		ag.AssetGroup.Status = enums.AssetGroupStatusEnum_PAUSED
	}
}

func (ag *AssetGroup) createOperation(tempId tempIdGenerator) *services.MutateOperation {
	ag.ResourceName = fmt.Sprintf("customers/%s/assetGroups/%s", ag.Campaign.Customer.GetId(), tempId())
	ag.AssetGroup.Campaign = ag.Campaign.GetResourceName()

	return &services.MutateOperation{
		Operation: &services.MutateOperation_AssetGroupOperation{
			AssetGroupOperation: &services.AssetGroupOperation{
				Operation: &services.AssetGroupOperation_Create{
					Create: ag.AssetGroup,
				},
			},
		},
	}
}

func (ag *AssetGroup) Create(ctx context.Context) error {
	tempId := newTempIdGenerator()

	ops := make([]*services.MutateOperation, 0)

	ops = append(ops, ag.createOperation(tempId))
	ops = append(ops, ag.Assets.createOperations(ag, tempId)...)

	_, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       ag.Campaign.Customer.GetId(),
		MutateOperations: ops,
	})
	return err
}

type AssetGroups []*AssetGroup

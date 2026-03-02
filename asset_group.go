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

type AssetGroup struct {
	*resources.AssetGroup
	Campaign *Campaign
	Assets   AssetGroupAssets
	resource
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
	ag.addUpdatedField("name")
}

func (ag *AssetGroup) SetFinalUrls(urls []string) {
	ag.AssetGroup.FinalUrls = urls
	ag.addUpdatedField("final_urls")
}

func (ag *AssetGroup) Save(ctx context.Context) error {
	if ag.isNew(ag.GetId()) {
		return ag.Create(ctx)
	}
	return ag.Update(ctx)
}

func (ag *AssetGroup) Update(ctx context.Context) error {
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

func (ag *AssetGroup) updateOperation(paths []string) *services.MutateOperation {
	return &services.MutateOperation{
		Operation: &services.MutateOperation_AssetGroupOperation{
			AssetGroupOperation: &services.AssetGroupOperation{
				Operation: &services.AssetGroupOperation_Update{
					Update: ag.AssetGroup,
				},
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: paths,
				},
			},
		},
	}
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
	ag.addUpdatedField("status")
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

func (ag *AssetGroup) createOperations(tempId tempIdGenerator) []*services.MutateOperation {
	ops := make([]*services.MutateOperation, 0)

	ops = append(ops, ag.createOperation(tempId))
	ops = append(ops, ag.Assets.createOperations(ag, tempId)...)

	return ops
}

func (ag *AssetGroup) Create(ctx context.Context) error {
	resp, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       ag.Campaign.Customer.GetId(),
		MutateOperations: ag.createOperations(newTempIdGenerator()),
	})
	if err != nil {
		return err
	}

	res := resp.GetMutateOperationResponses()[0].GetAssetGroupResult()
	ag.ResourceName = res.GetResourceName()
	ag.SetId(strings.Split(ag.ResourceName, "/")[3])

	return nil
}

type AssetGroups []*AssetGroup

func (ags AssetGroups) createOperations(tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(ags, func(item *AssetGroup) []*services.MutateOperation {
		return item.createOperations(tempId)
	}))
}

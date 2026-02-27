package googleads

import (
	"context"
	"fmt"

	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AssetGroup struct {
	*resources.AssetGroup
	Campaign *Campaign
	Assets   AssetGroupAssets
}

func NewAssetGroup() *AssetGroup {
	return &AssetGroup{
		AssetGroup: &resources.AssetGroup{},
		Campaign:   NewCampaign(),
		Assets:     NewAssetGroupAssets(),
	}
}

func (ag AssetGroup) GetName() string {
	return ag.AssetGroup.Name
}

func (ag *AssetGroup) SetName(name string) {
	ag.AssetGroup.Name = name
}

func (ag AssetGroup) GetFinalUrls() []string {
	return ag.AssetGroup.FinalUrls
}

func (ag *AssetGroup) SetFinalUrls(urls []string) {
	ag.AssetGroup.FinalUrls = urls
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

func NewAssetGroups() AssetGroups {
	return make(AssetGroups, 0)
}

package googleads

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type Campaign struct {
	*resources.Campaign
	Budget   *CampaignBudget
	Customer *Customer
	Criteria CampaignCriteria
	Assets   CampaignAssets
}

func NewCampaign() *Campaign {
	return &Campaign{
		Campaign: &resources.Campaign{},
		Budget:   NewCampaignBudget(),
		Customer: NewCustomer(),
		Criteria: NewCampaignCriteria(),
		Assets:   NewCampaignAssets(),
	}
}

func (c Campaign) GetId() string {
	return strconv.Itoa(int(c.Campaign.GetId()))
}

func (c *Campaign) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	c.Campaign.Id = Int64(i)
}

func (c Campaign) GetName() string {
	return c.Campaign.GetName()
}

func (c *Campaign) SetName(name string) {
	c.Campaign.Name = String(name)
	c.Budget.Name = String(name)
}

func (c Campaign) GetEnabled() bool {
	return c.Campaign.GetStatus() == enums.CampaignStatusEnum_ENABLED
}

func (c *Campaign) SetEnabled(enabled bool) {
	if enabled {
		c.Campaign.Status = enums.CampaignStatusEnum_ENABLED
	} else {
		c.Campaign.Status = enums.CampaignStatusEnum_PAUSED
	}
}

func (c Campaign) GetBudget() int {
	return c.Budget.GetAmountCents()
}

func (c *Campaign) SetBudget(budget int) {
	c.Budget.SetAmountCents(budget)
}

func (c *Campaign) ListCriteria(ctx context.Context) (CampaignCriteria, error) {
	return ListCampaignCriteria(ctx, c.Customer.GetId(), CampaignCriterionByCampaign(c.GetResourceName()))
}

func (c *Campaign) ListAssets(ctx context.Context) (CampaignAssets, error) {
	return ListCampaignAssets(ctx, c.Customer.GetId(), CampaignAssetByCampaign(c.GetResourceName()))
}

func (c *Campaign) ListAssetGroups(ctx context.Context) (AssetGroups, error) {
	return ListAssetGroups(ctx, c.Customer.GetId(), AssetGroupByCampaign(c.GetResourceName()))
}

func (c *Campaign) AddAssetGroup() *AssetGroup {
	return &AssetGroup{
		AssetGroup: &resources.AssetGroup{},
		Campaign:   c,
		Assets:     NewAssetGroupAssets(),
	}
}

func (c *Campaign) AddAdGroup() *AdGroup {
	return &AdGroup{
		AdGroup:  &resources.AdGroup{},
		Campaign: c,
		Criteria: NewAdGroupCriteria(),
		Assets:   NewAdGroupAssets(),
		Ads:      NewAdGroupAds(),
	}
}

func (c *Campaign) Create(ctx context.Context, customer *Customer) error {
	tempId := newTempIdGenerator()

	ops := make([]*services.MutateOperation, 0)

	ops = append(ops, c.Budget.createOperation(customer, tempId))
	ops = append(ops, c.createOperation(customer, c.Budget, tempId))
	ops = append(ops, c.Criteria.createOperations(c)...)
	ops = append(ops, c.Assets.createOperations(customer, c, tempId)...)

	_, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       customer.GetId(),
		MutateOperations: ops,
	})
	return err
}

func (c *Campaign) createOperation(customer *Customer, budget *CampaignBudget, tempId tempIdGenerator) *services.MutateOperation {
	c.ResourceName = fmt.Sprintf("customers/%s/campaigns/%s", customer.GetId(), tempId())
	c.CampaignBudget = String(budget.GetResourceName())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_CampaignOperation{
			CampaignOperation: &services.CampaignOperation{
				Operation: &services.CampaignOperation_Create{
					Create: c.Campaign,
				},
			},
		},
	}
}

type Campaigns []*Campaign

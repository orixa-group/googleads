package googleads

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type Campaign struct {
	*resources.Campaign
	Budget      *CampaignBudget
	Customer    *Customer
	Criteria    CampaignCriteria
	Assets      CampaignAssets
	AdGroups    AdGroups
	AssetGroups AssetGroups
	resource
}

func NewEmptyCampaign() *Campaign {
	return &Campaign{
		Campaign:    &resources.Campaign{},
		Budget:      &CampaignBudget{CampaignBudget: &resources.CampaignBudget{}},
		Customer:    &Customer{},
		Criteria:    make(CampaignCriteria, 0),
		Assets:      make(CampaignAssets, 0),
		AdGroups:    make(AdGroups, 0),
		AssetGroups: make(AssetGroups, 0),
	}
}

func NewSearchCampaign(name string, enabled bool, budget int) *Campaign {
	c := NewEmptyCampaign()
	c.Budget.DeliveryMethod = enums.BudgetDeliveryMethodEnum_STANDARD
	c.Budget.ExplicitlyShared = Bool(false)

	c.SetName(name)
	c.SetEnabled(enabled)
	c.SetBudget(budget)
	c.SetChannelType(ChannelTypeSearch)
	c.ContainsEuPoliticalAdvertising = enums.EuPoliticalAdvertisingStatusEnum_DOES_NOT_CONTAIN_EU_POLITICAL_ADVERTISING

	return c
}

func NewPerformanceMaxCampaign(name string, enabled bool, budget int) *Campaign {
	c := NewEmptyCampaign()
	c.Budget.DeliveryMethod = enums.BudgetDeliveryMethodEnum_STANDARD
	c.Budget.ExplicitlyShared = Bool(false)

	c.SetName(name)
	c.SetEnabled(enabled)
	c.SetBudget(budget)
	c.SetChannelType(ChannelTypePerformanceMax)
	c.ContainsEuPoliticalAdvertising = enums.EuPoliticalAdvertisingStatusEnum_DOES_NOT_CONTAIN_EU_POLITICAL_ADVERTISING

	return c
}

func (c Campaign) GetId() string {
	return strconv.Itoa(int(c.Campaign.GetId()))
}

func (c *Campaign) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	c.Campaign.Id = Int64(i)
}

func (c *Campaign) SetName(name string) {
	c.Campaign.Name = String(name)
	c.Budget.SetName(name)
	c.addUpdatedField("name")
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
	c.addUpdatedField("status")
}

func (c Campaign) GetBudget() int {
	return c.Budget.GetAmountCents()
}

func (c *Campaign) SetBudget(budget int) {
	c.Budget.SetAmountCents(budget)
}

func (c Campaign) GetChannelType() ChannelType {
	return enumToChannelType[c.Campaign.GetAdvertisingChannelType()]
}

func (c *Campaign) SetChannelType(channel ChannelType) {
	c.Campaign.AdvertisingChannelType = channelTypeToEnum[channel]

	if channel.is(ChannelTypeSearch) {
		c.Campaign.NetworkSettings = &resources.Campaign_NetworkSettings{
			TargetGoogleSearch:         Bool(true),
			TargetSearchNetwork:        Bool(true),
			TargetContentNetwork:       Bool(false),
			TargetPartnerSearchNetwork: Bool(false),
			TargetYoutube:              Bool(false),
			TargetGoogleTvNetwork:      Bool(false),
		}
	}
}

func (c *Campaign) SetObjective(objective Objective) {
	objectiveToEnum[objective](c)
	c.addUpdatedField("campaign_bidding_strategy")
}

func (c Campaign) GetObjective() Objective {
	return enumToObjective[c.Campaign.GetBiddingStrategyType()]
}

func (c Campaign) GetStartDate() string {
	return c.Campaign.GetStartDateTime()
}

func (c *Campaign) SetStartDate(date string) {
	c.Campaign.StartDateTime = String(date)
	c.addUpdatedField("start_date_time")
}

func (c Campaign) GetEndDate() string {
	return c.Campaign.GetEndDateTime()
}

func (c *Campaign) SetEndDate(date string) {
	c.Campaign.EndDateTime = String(date)
	c.addUpdatedField("end_date_time")
}

func (c *Campaign) IsPMax() bool {
	return c.GetChannelType().is(ChannelTypePerformanceMax)
}

func (c *Campaign) IsSearch() bool {
	return c.GetChannelType().is(ChannelTypeSearch)
}

func (c *Campaign) Save(ctx context.Context) error {
	if c.isNew(c.GetId()) {
		return c.Create(ctx, c.Customer)
	}
	return c.Update(ctx)
}

func (c *Campaign) Update(ctx context.Context) error {
	ops := make([]*services.MutateOperation, 0)

	paths := c.getUpdatedFields()
	if len(paths) > 0 {
		ops = append(ops, c.updateOperation(paths))
	}

	budgetPaths := c.Budget.getUpdatedFields()
	if len(budgetPaths) > 0 {
		ops = append(ops, c.Budget.updateOperation(budgetPaths))
	}

	if len(ops) == 0 {
		return nil
	}

	_, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       c.Customer.GetId(),
		MutateOperations: ops,
	})
	return err
}

func (c *Campaign) updateOperation(paths []string) *services.MutateOperation {
	return &services.MutateOperation{
		Operation: &services.MutateOperation_CampaignOperation{
			CampaignOperation: &services.CampaignOperation{
				Operation: &services.CampaignOperation_Update{
					Update: c.Campaign,
				},
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: paths,
				},
			},
		},
	}
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

func (c *Campaign) ListAdGroups(ctx context.Context) (AdGroups, error) {
	return ListAdGroups(ctx, c.Customer.GetId(), AdGroupsByCampaign(c.GetResourceName()))
}

func (c *Campaign) NewAssetGroup(name string) *AssetGroup {
	ag := &AssetGroup{
		AssetGroup: &resources.AssetGroup{
			Name: name,
		},
		Campaign: c,
		Assets:   make(AssetGroupAssets, 0),
	}
	c.AssetGroups = append(c.AssetGroups, ag)
	return ag
}

func (c *Campaign) NewAdGroup(name string) *AdGroup {
	ag := &AdGroup{
		AdGroup: &resources.AdGroup{
			Name: String(name),
		},
		Campaign: c,
		Criteria: make(AdGroupCriteria, 0),
		Assets:   make(AdGroupAssets, 0),
		Ads:      make(AdGroupAds, 0),
	}
	c.AdGroups = append(c.AdGroups, ag)
	return ag
}

func (c *Campaign) Create(ctx context.Context, customer *Customer) error {
	tempId := newTempIdGenerator()

	ops := make([]*services.MutateOperation, 0)

	ops = append(ops, c.Budget.createOperation(customer, tempId))
	ops = append(ops, c.createOperation(customer, c.Budget, tempId))
	ops = append(ops, c.Criteria.createOperations(c)...)
	ops = append(ops, c.Assets.createOperations(customer, c, tempId)...)
	ops = append(ops, c.AdGroups.createOperations(tempId)...)
	ops = append(ops, c.AssetGroups.createOperations(tempId)...)

	resp, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       customer.GetId(),
		MutateOperations: ops,
	})
	if err != nil {
		return err
	}

	c.ResourceName = resp.GetMutateOperationResponses()[1].GetCampaignResult().GetResourceName()
	c.Id = Int64(resp.GetMutateOperationResponses()[1].GetCampaignResult().GetCampaign().GetId())

	return nil
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

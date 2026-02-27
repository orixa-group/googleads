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
		Customer: &Customer{},
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

func (c Campaign) GetAdvertisingChannelType() enums.AdvertisingChannelTypeEnum_AdvertisingChannelType {
	return c.Campaign.GetAdvertisingChannelType()
}

func (c *Campaign) SetAdvertisingChannel(channel AdvertisingChannel) {
	channel(c.Campaign)
}

func (c *Campaign) SetNetworkSettings(settings ...NetworkSetting) {
	if c.NetworkSettings == nil {
		c.NetworkSettings = &resources.Campaign_NetworkSettings{}
	}
	for _, s := range settings {
		s(c.NetworkSettings)
	}
}

func (c *Campaign) SetContainsEuPoliticalAdvertising(containsEuPoliticalAdvertising bool) {
	if containsEuPoliticalAdvertising {
		c.Campaign.ContainsEuPoliticalAdvertising = enums.EuPoliticalAdvertisingStatusEnum_CONTAINS_EU_POLITICAL_ADVERTISING
	} else {
		c.Campaign.ContainsEuPoliticalAdvertising = enums.EuPoliticalAdvertisingStatusEnum_DOES_NOT_CONTAIN_EU_POLITICAL_ADVERTISING
	}
}

func (c *Campaign) SetBiddingStrategy(strategy BiddingStrategy) {
	strategy(c.Campaign)
}

func (c *Campaign) SetStartDate(date string) {
	c.Campaign.StartDateTime = String(date)
}

func (c *Campaign) SetEndDate(date string) {
	c.Campaign.EndDateTime = String(date)
}

func (c Campaign) GetStartDate() string {
	return c.Campaign.GetStartDateTime()
}

func (c Campaign) GetEndDate() string {
	return c.Campaign.GetEndDateTime()
}

func (c Campaign) IsBrandGuidelinesEnabled() bool {
	return c.Campaign.GetBrandGuidelinesEnabled()
}

func (c Campaign) GetContainsEuPoliticalAdvertising() bool {
	return c.Campaign.GetContainsEuPoliticalAdvertising() == enums.EuPoliticalAdvertisingStatusEnum_CONTAINS_EU_POLITICAL_ADVERTISING
}

func (c Campaign) GetTrackingUrl() string {
	return c.Campaign.GetTrackingUrlTemplate()
}

func (c *Campaign) SetTrackingUrl(url string) {
	c.Campaign.TrackingUrlTemplate = String(url)
}

func (c Campaign) GetFinalUrlSuffix() string {
	return c.Campaign.GetFinalUrlSuffix()
}

func (c *Campaign) SetFinalUrlSuffix(suffix string) {
	c.Campaign.FinalUrlSuffix = String(suffix)
}

func (c Campaign) GetBiddingStrategyType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	return c.Campaign.GetBiddingStrategyType()
}

func (c Campaign) GetAdServingOptimizationStatus() enums.AdServingOptimizationStatusEnum_AdServingOptimizationStatus {
	return c.Campaign.GetAdServingOptimizationStatus()
}

func (c *Campaign) SetKeywordMatchType(matchType CampaignKeywordMatchType) {
	matchType(c.Campaign)
}

func (c *Campaign) SetGeoTargetTypeSetting(settings ...GeoTargetTypeSetting) {
	if c.GeoTargetTypeSetting == nil {
		c.GeoTargetTypeSetting = &resources.Campaign_GeoTargetTypeSetting{}
	}
	for _, s := range settings {
		s(c.GeoTargetTypeSetting)
	}
}

func (c *Campaign) SetBrandGuidelinesEnabled(enabled bool) {
	c.BrandGuidelinesEnabled = Bool(enabled)
}

func (c *Campaign) SetAdServingOptimizationStatus(status AdServingOptimizationStatus) {
	status(c.Campaign)
}

func (c *Campaign) AddLabel(labelResourceName string) {
	c.Campaign.Labels = append(c.Campaign.Labels, labelResourceName)
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

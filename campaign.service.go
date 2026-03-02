package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCampaign(ctx context.Context, customerId string, filters ...CampaignFilter) (*Campaign, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignQueryBuilder().Where(filters...).Build(), createCampaignInstance)
}

func ListCampaigns(ctx context.Context, customerId string, filters ...CampaignFilter) (Campaigns, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignQueryBuilder().Where(filters...).Build(), createCampaignInstance)
}

func CreateCampaign(ctx context.Context, customer *Customer, campaign *Campaign) error {
	return campaign.Create(ctx, customer)
}

func createCampaignInstance(row *services.GoogleAdsRow) *Campaign {
	return &Campaign{
		Campaign: row.GetCampaign(),
		Budget:   &CampaignBudget{CampaignBudget: row.GetCampaignBudget()},
		Customer: &Customer{
			Customer: row.GetCustomer(),
			Assets:   make(CustomerAssets, 0),
		},
		Criteria: make(CampaignCriteria, 0),
		Assets:   make(CampaignAssets, 0),
	}
}

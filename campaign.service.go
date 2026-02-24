package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCampaign(ctx context.Context, customerId string, filters ...CampaignFilter) (*Campaign, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignQueryBuilder().Where(filters...).Build(), createCampaignInstance)
}

func ListCampaigns(ctx context.Context, customerId string, filters ...CampaignFilter) ([]*Campaign, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignQueryBuilder().Where(filters...).Build(), createCampaignInstance)
}

func CreateCampaign(ctx context.Context, customerId string, req *Campaign) (*Campaign, error) {
	return Create(ctx, services.NewCampaignServiceClient(instance.conn).MutateCampaigns, &services.MutateCampaignsRequest{
		CustomerId: customerId,
		Operations: []*services.CampaignOperation{
			{
				Operation: &services.CampaignOperation_Create{Create: req.Campaign},
			},
		},
	}, func(customerId string, res *services.MutateCampaignsResponse) string {
		return res.GetResults()[0].GetResourceName()
	}, services.NewGoogleAdsServiceClient(instance.conn), NewCampaignQueryBuilder(), CampaignByResourceName, createCampaignInstance)
}

func createCampaignInstance(row *services.GoogleAdsRow) *Campaign {
	return &Campaign{
		Campaign: row.GetCampaign(),
		Budget:   &CampaignBudget{row.GetCampaignBudget()},
	}
}

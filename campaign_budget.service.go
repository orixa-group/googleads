package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCampaignBudget(ctx context.Context, customerId string, filters ...CampaignBudgetFilter) (*CampaignBudget, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignBudgetQueryBuilder().Where(filters...).Build(), createCampaignBudgetInstance)
}

func ListCampaignBudgets(ctx context.Context, customerId string, filters ...CampaignBudgetFilter) ([]*CampaignBudget, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignBudgetQueryBuilder().Where(filters...).Build(), createCampaignBudgetInstance)
}

func CreateCampaignBudget(ctx context.Context, customerId string, req *CampaignBudget) (*CampaignBudget, error) {
	return Create(ctx, services.NewCampaignBudgetServiceClient(instance.conn).MutateCampaignBudgets, &services.MutateCampaignBudgetsRequest{
		CustomerId: customerId,
		Operations: []*services.CampaignBudgetOperation{
			{
				Operation: &services.CampaignBudgetOperation_Create{Create: req.CampaignBudget},
			},
		},
	}, func(customerId string, res *services.MutateCampaignBudgetsResponse) string {
		return res.GetResults()[0].GetResourceName()
	}, services.NewGoogleAdsServiceClient(instance.conn), NewCampaignBudgetQueryBuilder(), CampaignBudgetByResourceName, createCampaignBudgetInstance)
}

func createCampaignBudgetInstance(row *services.GoogleAdsRow) *CampaignBudget {
	return &CampaignBudget{row.GetCampaignBudget()}
}

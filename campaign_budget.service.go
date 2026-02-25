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

func createCampaignBudgetInstance(row *services.GoogleAdsRow) *CampaignBudget {
	return &CampaignBudget{row.GetCampaignBudget()}
}

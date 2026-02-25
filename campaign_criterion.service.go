package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCampaignCriterion(ctx context.Context, customerId string, filters ...CampaignCriterionFilter) (*CampaignCriterion, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignCriterionQueryBuilder().Where(filters...).Build(), createCampaignCriterionInstance)
}

func ListCampaignCriteria(ctx context.Context, customerId string, filters ...CampaignCriterionFilter) ([]*CampaignCriterion, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignCriterionQueryBuilder().Where(filters...).Build(), createCampaignCriterionInstance)
}

func createCampaignCriterionInstance(row *services.GoogleAdsRow) *CampaignCriterion {
	return &CampaignCriterion{row.GetCampaignCriterion()}
}

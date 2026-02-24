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

func CreateCampaignCriterion(ctx context.Context, customerId string, req *services.CampaignCriterionOperation_Create) (*CampaignCriterion, error) {
	return Create(ctx, services.NewCampaignCriterionServiceClient(instance.conn).MutateCampaignCriteria, &services.MutateCampaignCriteriaRequest{
		CustomerId: customerId,
		Operations: []*services.CampaignCriterionOperation{
			{
				Operation: req,
			},
		},
	}, func(customerId string, res *services.MutateCampaignCriteriaResponse) string {
		return res.GetResults()[0].GetResourceName()
	}, services.NewGoogleAdsServiceClient(instance.conn), NewCampaignCriterionQueryBuilder(), CampaignCriterionByResourceName, createCampaignCriterionInstance)
}

func createCampaignCriterionInstance(row *services.GoogleAdsRow) *CampaignCriterion {
	return &CampaignCriterion{row.GetCampaignCriterion()}
}

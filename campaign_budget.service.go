package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

type CampaignBudgetService struct {
	conn grpc.ClientConnInterface
}

func NewCampaignBudgetService(conn grpc.ClientConnInterface) *CampaignBudgetService {
	return &CampaignBudgetService{conn}
}

func (s *CampaignBudgetService) Fetch(ctx context.Context, customerId string, filters ...CampaignBudgetFilter) (*CampaignBudget, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(s.conn), customerId, NewCampaignBudgetQueryBuilder().Where(filters...).Build(), s.createInstance)
}

func (s *CampaignBudgetService) List(ctx context.Context, customerId string, filters ...CampaignBudgetFilter) ([]*CampaignBudget, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(s.conn), customerId, NewCampaignBudgetQueryBuilder().Where(filters...).Build(), s.createInstance)
}

func (s *CampaignBudgetService) Create(ctx context.Context, customerId string, req *CampaignBudget) (*CampaignBudget, error) {
	return Create(ctx, services.NewCampaignBudgetServiceClient(s.conn).MutateCampaignBudgets, &services.MutateCampaignBudgetsRequest{
		CustomerId: customerId,
		Operations: []*services.CampaignBudgetOperation{
			{
				Operation: &services.CampaignBudgetOperation_Create{Create: req.CampaignBudget},
			},
		},
	}, func(customerId string, res *services.MutateCampaignBudgetsResponse) string {
		return res.GetResults()[0].GetResourceName()
	}, services.NewGoogleAdsServiceClient(s.conn), NewCampaignBudgetQueryBuilder(), CampaignBudgetByResourceName, s.createInstance)
}

func (s *CampaignBudgetService) createInstance(row *services.GoogleAdsRow) *CampaignBudget {
	return &CampaignBudget{row.GetCampaignBudget()}
}

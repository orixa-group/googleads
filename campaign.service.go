package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

type CampaignService struct {
	conn grpc.ClientConnInterface
}

func NewCampaignService(conn grpc.ClientConnInterface) *CampaignService {
	return &CampaignService{conn}
}

func (s *CampaignService) Fetch(ctx context.Context, customerId string, filters ...CampaignFilter) (*Campaign, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(s.conn), customerId, NewCampaignQueryBuilder().Where(filters...).Build(), s.createInstance)
}

func (s *CampaignService) List(ctx context.Context, customerId string, filters ...CampaignFilter) ([]*Campaign, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(s.conn), customerId, NewCampaignQueryBuilder().Where(filters...).Build(), s.createInstance)
}

func (s *CampaignService) Create(ctx context.Context, customerId string, req *resources.Campaign) (*Campaign, error) {
	return Create(ctx, services.NewCampaignServiceClient(s.conn).MutateCampaigns, &services.MutateCampaignsRequest{
		CustomerId: customerId,
		Operations: []*services.CampaignOperation{
			{
				Operation: &services.CampaignOperation_Create{Create: req},
			},
		},
	}, func(customerId string, res *services.MutateCampaignsResponse) string {
		return res.GetResults()[0].GetResourceName()
	}, services.NewGoogleAdsServiceClient(s.conn), NewCampaignQueryBuilder(), CampaignByResourceName, s.createInstance)
}

func (s *CampaignService) createInstance(row *services.GoogleAdsRow) *Campaign {
	return &Campaign{
		Campaign: row.GetCampaign(),
		Budget:   &CampaignBudget{row.GetCampaignBudget()},
	}
}

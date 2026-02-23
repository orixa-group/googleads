package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

type GoogleAdsService struct {
	conn grpc.ClientConnInterface
}

func NewGoogleAdsService(conn grpc.ClientConnInterface) *GoogleAdsService {
	return &GoogleAdsService{conn}
}

func (s *GoogleAdsService) Mutate(ctx context.Context, customerId string, operations []*services.MutateOperation) (*services.MutateGoogleAdsResponse, error) {
	req := &services.MutateGoogleAdsRequest{
		CustomerId:       customerId,
		MutateOperations: operations,
	}

	return services.NewGoogleAdsServiceClient(s.conn).Mutate(ctx, req)
}

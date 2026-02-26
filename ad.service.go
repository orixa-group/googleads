package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAd(ctx context.Context, customerId string, filters ...AdFilter) (*Ad, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdQueryBuilder().Where(filters...).Build(), createAdInstance)
}

func ListAds(ctx context.Context, customerId string, filters ...AdFilter) ([]*Ad, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdQueryBuilder().Where(filters...).Build(), createAdInstance)
}

func createAdInstance(row *services.GoogleAdsRow) *Ad {
	return &Ad{row.GetAd()}
}

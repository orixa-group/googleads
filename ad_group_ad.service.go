package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAdGroupAd(ctx context.Context, customerId string, filters ...AdGroupAdFilter) (*AdGroupAd, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdGroupAdQueryBuilder().Where(filters...).Build(), createAdGroupAdInstance)
}

func ListAdGroupAds(ctx context.Context, customerId string, filters ...AdGroupAdFilter) ([]*AdGroupAd, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdGroupAdQueryBuilder().Where(filters...).Build(), createAdGroupAdInstance)
}

func createAdGroupAdInstance(row *services.GoogleAdsRow) *AdGroupAd {
	return &AdGroupAd{row.GetAdGroupAd()}
}

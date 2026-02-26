package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAdGroupAsset(ctx context.Context, customerId string, filters ...AdGroupAssetFilter) (*AdGroupAsset, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdGroupAssetQueryBuilder().Where(filters...).Build(), createAdGroupAssetInstance)
}

func ListAdGroupAssets(ctx context.Context, customerId string, filters ...AdGroupAssetFilter) ([]*AdGroupAsset, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdGroupAssetQueryBuilder().Where(filters...).Build(), createAdGroupAssetInstance)
}

func createAdGroupAssetInstance(row *services.GoogleAdsRow) *AdGroupAsset {
	return &AdGroupAsset{
		AdGroupAsset: row.GetAdGroupAsset(),
		Asset:        &Asset{row.GetAsset()},
	}
}

package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAsset(ctx context.Context, customerId string, filters ...AssetFilter) (*Asset, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAssetQueryBuilder().Where(filters...).Build(), createAssetInstance)
}

func ListAssets(ctx context.Context, customerId string, filters ...AssetFilter) ([]*Asset, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAssetQueryBuilder().Where(filters...).Build(), createAssetInstance)
}

func createAssetInstance(row *services.GoogleAdsRow) *Asset {
	return &Asset{row.GetAsset()}
}

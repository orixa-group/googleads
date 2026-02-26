package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAssetGroupAsset(ctx context.Context, customerId string, filters ...AssetGroupAssetFilter) (*AssetGroupAsset, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAssetGroupAssetQueryBuilder().Where(filters...).Build(), createAssetGroupAssetInstance)
}

func ListAssetGroupAssets(ctx context.Context, customerId string, filters ...AssetGroupAssetFilter) ([]*AssetGroupAsset, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAssetGroupAssetQueryBuilder().Where(filters...).Build(), createAssetGroupAssetInstance)
}

func createAssetGroupAssetInstance(row *services.GoogleAdsRow) *AssetGroupAsset {
	return &AssetGroupAsset{
		AssetGroupAsset: row.GetAssetGroupAsset(),
		Asset:           &Asset{row.GetAsset()},
	}
}

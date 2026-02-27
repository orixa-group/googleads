package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCustomerAsset(ctx context.Context, customerId string, filters ...CustomerAssetFilter) (*CustomerAsset, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCustomerAssetQueryBuilder().Where(filters...).Build(), createCustomerAssetInstance)
}

func ListCustomerAssets(ctx context.Context, customerId string, filters ...CustomerAssetFilter) ([]*CustomerAsset, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCustomerAssetQueryBuilder().Where(filters...).Build(), createCustomerAssetInstance)
}

func createCustomerAssetInstance(row *services.GoogleAdsRow) *CustomerAsset {
	return &CustomerAsset{
		CustomerAsset: row.GetCustomerAsset(),
		Asset:         &Asset{row.GetAsset()},
	}
}

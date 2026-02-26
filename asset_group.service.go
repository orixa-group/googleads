package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAssetGroup(ctx context.Context, customerId string, filters ...AssetGroupFilter) (*AssetGroup, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAssetGroupQueryBuilder().Where(filters...).Build(), createAssetGroupInstance)
}

func ListAssetGroups(ctx context.Context, customerId string, filters ...AssetGroupFilter) ([]*AssetGroup, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAssetGroupQueryBuilder().Where(filters...).Build(), createAssetGroupInstance)
}

func createAssetGroupInstance(row *services.GoogleAdsRow) *AssetGroup {
	return &AssetGroup{
		AssetGroup: row.GetAssetGroup(),
		Campaign: &Campaign{
			Campaign: row.GetCampaign(),
			Customer: &Customer{
				Customer: row.GetCustomer(),
			},
		},
	}
}

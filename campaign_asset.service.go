package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCampaignAsset(ctx context.Context, customerId string, filters ...CampaignAssetFilter) (*CampaignAsset, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignAssetQueryBuilder().Where(filters...).Build(), createCampaignAssetInstance)
}

func ListCampaignAssets(ctx context.Context, customerId string, filters ...CampaignAssetFilter) ([]*CampaignAsset, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCampaignAssetQueryBuilder().Where(filters...).Build(), createCampaignAssetInstance)
}

func createCampaignAssetInstance(row *services.GoogleAdsRow) *CampaignAsset {
	return &CampaignAsset{
		CampaignAsset: row.GetCampaignAsset(),
		Asset:         &Asset{row.GetAsset()},
	}
}

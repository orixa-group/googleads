package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAdGroup(ctx context.Context, customerId string, filters ...AdGroupFilter) (*AdGroup, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdGroupQueryBuilder().Where(filters...).Build(), createAdGroupInstance)
}

func ListAdGroups(ctx context.Context, customerId string, filters ...AdGroupFilter) ([]*AdGroup, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdGroupQueryBuilder().Where(filters...).Build(), createAdGroupInstance)
}

func createAdGroupInstance(row *services.GoogleAdsRow) *AdGroup {
	return &AdGroup{
		AdGroup: row.GetAdGroup(),
		Campaign: &Campaign{
			Campaign: row.GetCampaign(),
			Customer: &Customer{
				Customer: row.GetCustomer(),
			},
		},
	}
}

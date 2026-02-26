package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAdGroupCriterion(ctx context.Context, customerId string, filters ...AdGroupCriterionFilter) (*AdGroupCriterion, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdGroupCriterionQueryBuilder().Where(filters...).Build(), createAdGroupCriterionInstance)
}

func ListAdGroupCriteria(ctx context.Context, customerId string, filters ...AdGroupCriterionFilter) ([]*AdGroupCriterion, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewAdGroupCriterionQueryBuilder().Where(filters...).Build(), createAdGroupCriterionInstance)
}

func createAdGroupCriterionInstance(row *services.GoogleAdsRow) *AdGroupCriterion {
	return &AdGroupCriterion{row.GetAdGroupCriterion()}
}

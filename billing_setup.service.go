package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchBillingSetup(ctx context.Context, customer *Customer, filters ...BillingSetupFilter) (*BillingSetup, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customer.GetId(), NewBillingSetupQueryBuilder().Where(filters...).Build(), createBillingSetupInstance)
}

func ListBillingSetups(ctx context.Context, customer *Customer, filters ...BillingSetupFilter) ([]*BillingSetup, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customer.GetId(), NewBillingSetupQueryBuilder().Where(filters...).Build(), createBillingSetupInstance)
}

func CreateBillingSetup(ctx context.Context, customer *Customer, paymentsAccountId string) (*BillingSetup, error) {
	return customer.CreateBillingSetup(ctx, paymentsAccountId)
}

func createBillingSetupInstance(row *services.GoogleAdsRow) *BillingSetup {
	return &BillingSetup{row.GetBillingSetup()}
}

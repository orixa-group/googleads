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

func CreateBillingSetup(ctx context.Context, customer *Customer, bs *BillingSetup) (*BillingSetup, error) {
	return Create(ctx, services.NewBillingSetupServiceClient(instance.conn).MutateBillingSetup, &services.MutateBillingSetupRequest{
		CustomerId: customer.GetId(),
		Operation: &services.BillingSetupOperation{
			Operation: &services.BillingSetupOperation_Create{
				Create: bs.BillingSetup,
			},
		},
	}, func(customerId string, res *services.MutateBillingSetupResponse) string {
		return res.GetResult().GetResourceName()
	}, services.NewGoogleAdsServiceClient(instance.conn), NewBillingSetupQueryBuilder(), BillingSetupByResourceName, createBillingSetupInstance)
}

func createBillingSetupInstance(row *services.GoogleAdsRow) *BillingSetup {
	return &BillingSetup{row.GetBillingSetup()}
}

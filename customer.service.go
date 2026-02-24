package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCustomer(ctx context.Context, customerId string, filters ...CustomerFilter) (*Customer, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCustomerQueryBuilder().Where(filters...).Build(), createCustomerInstance)
}

func ListCustomers(ctx context.Context, customerId string, filters ...CustomerFilter) ([]*Customer, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customerId, NewCustomerQueryBuilder().Where(filters...).Build(), createCustomerInstance)
}

func createCustomerInstance(row *services.GoogleAdsRow) *Customer {
	return &Customer{row.GetCustomer()}
}

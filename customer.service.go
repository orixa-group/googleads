package googleads

import (
	"context"
	"fmt"
	"strings"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCustomer(ctx context.Context, id string) (*Customer, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), id, NewCustomerQueryBuilder().Where(func() string {
		return fmt.Sprintf("customer.id = '%s'", id)
	}).Build(), func(row *services.GoogleAdsRow) *Customer {
		return &Customer{row.GetCustomer()}
	})
}

func CreateCustomer(ctx context.Context, customer, parent *Customer) (*Customer, error) {
	resp, err := services.NewCustomerServiceClient(instance.conn).CreateCustomerClient(ctx, &services.CreateCustomerClientRequest{
		CustomerId:     parent.GetId(),
		CustomerClient: customer.Customer,
	})
	if err != nil {
		return nil, err
	}

	return FetchCustomer(ctx, strings.Split(resp.GetResourceName(), "/")[1])
}

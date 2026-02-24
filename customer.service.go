package googleads

import (
	"context"
	"fmt"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchCustomer(ctx context.Context, id string) (*Customer, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), id, NewCustomerQueryBuilder().Where(func() string {
		return fmt.Sprintf("customer.id = '%s'", id)
	}).Build(), func(row *services.GoogleAdsRow) *Customer {
		return &Customer{row.GetCustomer()}
	})
}

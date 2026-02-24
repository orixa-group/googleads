package googleads

import (
	"context"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/resources"
)

type Customer struct {
	*resources.Customer
}

func NewCustomer() *Customer {
	return &Customer{&resources.Customer{}}
}

func (c Customer) GetId() string {
	return strconv.Itoa(int(c.Customer.GetId()))
}

func (c Customer) ListCampaigns(ctx context.Context) (Campaigns, error) {
	return ListCampaigns(ctx, c.GetId())
}

func (c Customer) FetchCampaign(ctx context.Context, id string) (*Campaign, error) {
	return FetchCampaign(ctx, c.GetId(), CampaignById(id))
}

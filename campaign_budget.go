package googleads

import (
	"context"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/resources"
)

type CampaignBudget struct {
	*resources.CampaignBudget
}

func NewCampaignBudget() *CampaignBudget {
	return &CampaignBudget{
		CampaignBudget: &resources.CampaignBudget{},
	}
}

func (c CampaignBudget) GetId() string {
	return strconv.Itoa(int(c.CampaignBudget.GetId()))
}

func (c CampaignBudget) GetAmount() int {
	return int(c.CampaignBudget.GetAmountMicros() / 10_000)
}

func (c *CampaignBudget) SetAmount(amount int) {
	c.CampaignBudget.AmountMicros = Int64(int64(amount * 10_000))
}

func (c *CampaignBudget) Create(ctx context.Context, client *Client, customerId string) error {
	new, err := client.CampaignBudget().Create(ctx, customerId, c.CampaignBudget)
	if err != nil {
		return err
	}

	*c = *new
	return nil
}

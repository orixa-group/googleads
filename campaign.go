package googleads

import (
	"context"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type Campaign struct {
	*resources.Campaign
	Budget *CampaignBudget
}

func NewCampaign() *Campaign {
	return &Campaign{
		Campaign: &resources.Campaign{},
		Budget:   NewCampaignBudget(),
	}
}

func (c Campaign) GetId() string {
	return strconv.Itoa(int(c.Campaign.GetId()))
}

func (c *Campaign) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	c.Campaign.Id = Int64(i)
}

func (c Campaign) GetName() string {
	return c.Campaign.GetName()
}

func (c *Campaign) SetName(name string) {
	c.Campaign.Name = String(name)
}

func (c Campaign) GetEnabled() bool {
	return c.Campaign.GetStatus() == enums.CampaignStatusEnum_ENABLED
}

func (c *Campaign) SetEnabled(enabled bool) {
	if enabled {
		c.Campaign.Status = enums.CampaignStatusEnum_ENABLED
	} else {
		c.Campaign.Status = enums.CampaignStatusEnum_PAUSED
	}
}

func (c Campaign) GetBudget() int {
	return c.Budget.GetAmount()
}

func (c *Campaign) SetBudget(budget int) {
	c.Budget.SetAmount(budget)
}

func (c *Campaign) Create(ctx context.Context, client *Client, customerId string) error {
	new, err := client.Campaign().Create(ctx, customerId, c.Campaign)
	if err != nil {
		return err
	}

	*c = *new
	return nil
}

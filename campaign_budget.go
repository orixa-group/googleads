package googleads

import (
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

func (c CampaignBudget) GetAmountCents() int {
	return int(c.CampaignBudget.GetAmountMicros() / 10_000)
}

func (c *CampaignBudget) SetAmountCents(amount int) {
	c.CampaignBudget.AmountMicros = Int64(int64(amount * 10_000))
}

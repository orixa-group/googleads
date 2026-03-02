package googleads

import (
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type CampaignBudget struct {
	*resources.CampaignBudget
}

func (c CampaignBudget) GetId() string {
	return strconv.Itoa(int(c.CampaignBudget.GetId()))
}

func (c *CampaignBudget) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	c.CampaignBudget.Id = Int64(i)
}

func (c *CampaignBudget) SetName(name string) {
	c.CampaignBudget.Name = String(name)
}

func (c CampaignBudget) GetAmountCents() int {
	return int(c.CampaignBudget.GetAmountMicros() / 10_000)
}

func (c *CampaignBudget) SetAmountCents(amount int) {
	c.CampaignBudget.AmountMicros = Int64(int64(amount * 10_000))
}

func (c *CampaignBudget) createOperation(customer *Customer, tempId tempIdGenerator) *services.MutateOperation {
	c.ResourceName = fmt.Sprintf("customers/%s/campaignBudgets/%s", customer.GetId(), tempId())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_CampaignBudgetOperation{
			CampaignBudgetOperation: &services.CampaignBudgetOperation{
				Operation: &services.CampaignBudgetOperation_Create{
					Create: c.CampaignBudget,
				},
			},
		},
	}
}

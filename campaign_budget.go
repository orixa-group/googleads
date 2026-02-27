package googleads

import (
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type CampaignBudget struct {
	*resources.CampaignBudget
}

func NewCampaignBudget() *CampaignBudget {
	return &CampaignBudget{
		CampaignBudget: &resources.CampaignBudget{
			DeliveryMethod:   enums.BudgetDeliveryMethodEnum_STANDARD,
			ExplicitlyShared: Bool(false),
		},
	}
}

func (c CampaignBudget) GetId() string {
	return strconv.Itoa(int(c.CampaignBudget.GetId()))
}

func (c *CampaignBudget) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	c.CampaignBudget.Id = Int64(i)
}

func (c CampaignBudget) GetName() string {
	return c.CampaignBudget.GetName()
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

func (c CampaignBudget) GetEnabled() bool {
	return c.CampaignBudget.GetStatus() == enums.BudgetStatusEnum_ENABLED
}

func (c *CampaignBudget) SetEnabled(enabled bool) {
	if enabled {
		c.CampaignBudget.Status = enums.BudgetStatusEnum_ENABLED
	} else {
		c.CampaignBudget.Status = enums.BudgetStatusEnum_REMOVED
	}
}

func (c CampaignBudget) IsExplicitlyShared() bool {
	return c.CampaignBudget.GetExplicitlyShared()
}

func (c *CampaignBudget) SetExplicitlyShared(shared bool) {
	c.CampaignBudget.ExplicitlyShared = Bool(shared)
}

func (c *CampaignBudget) SetDeliveryMethod(method CampaignBudgetDeliveryMethod) {
	method(c.CampaignBudget)
}

func (c CampaignBudget) GetDeliveryMethod() enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod {
	return c.CampaignBudget.GetDeliveryMethod()
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

type CampaignBudgetDeliveryMethod func(budget *resources.CampaignBudget)

func BudgetDeliveryStandard(budget *resources.CampaignBudget) {
	budget.DeliveryMethod = enums.BudgetDeliveryMethodEnum_STANDARD
}

func BudgetDeliveryAccelerated(budget *resources.CampaignBudget) {
	budget.DeliveryMethod = enums.BudgetDeliveryMethodEnum_ACCELERATED
}

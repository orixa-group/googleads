package googleads

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type Customer struct {
	*resources.Customer
}

func NewCustomer() *Customer {
	return &Customer{&resources.Customer{
		CurrencyCode: String("EUR"),
		TimeZone:     String("Europe/Paris"),
	}}
}

func (c Customer) GetId() string {
	return strconv.Itoa(int(c.Customer.GetId()))
}

func (c *Customer) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	c.Customer.Id = Int64(i)
}

func (c Customer) GetName() string {
	return c.Customer.GetDescriptiveName()
}

func (c *Customer) SetName(name string) {
	c.Customer.DescriptiveName = String(name)
}

func (c Customer) ListCampaigns(ctx context.Context) (Campaigns, error) {
	return ListCampaigns(ctx, c.GetId())
}

func (c Customer) FetchCampaign(ctx context.Context, id string) (*Campaign, error) {
	return FetchCampaign(ctx, c.GetId(), CampaignById(id))
}

func (c *Customer) Create(ctx context.Context, parent *Customer) error {
	customer, err := CreateCustomer(ctx, c, parent)
	if err != nil {
		return err
	}

	c.Customer = customer.Customer
	return nil
}

func (c *Customer) CreateBillingSetup(ctx context.Context, paymentsAccountId string) (*BillingSetup, error) {
	return CreateBillingSetup(ctx, c, &BillingSetup{&resources.BillingSetup{
		PaymentsAccount: String(fmt.Sprintf("customers/%s/paymentsAccounts/%s", c.GetId(), paymentsAccountId)),
		StartTime: &resources.BillingSetup_StartTimeType{
			StartTimeType: enums.TimeTypeEnum_NOW,
		},
	}})
}

func (c *Customer) CreateAccountBudget(ctx context.Context, bs *BillingSetup) (*AccountBudget, error) {
	return CreateAccountBudget(ctx, c, &AccountBudget{&resources.AccountBudgetProposal{
		BillingSetup: String(bs.GetResourceName()),
		ProposalType: enums.AccountBudgetProposalTypeEnum_CREATE,
		ProposedName: String(c.GetDescriptiveName()),
		ProposedStartTime: &resources.AccountBudgetProposal_ProposedStartTimeType{
			ProposedStartTimeType: enums.TimeTypeEnum_NOW,
		},
		ProposedEndTime: &resources.AccountBudgetProposal_ProposedEndTimeType{
			ProposedEndTimeType: enums.TimeTypeEnum_FOREVER,
		},
		ProposedSpendingLimit: &resources.AccountBudgetProposal_ProposedSpendingLimitType{
			ProposedSpendingLimitType: enums.SpendingLimitTypeEnum_INFINITE,
		},
	}})
}

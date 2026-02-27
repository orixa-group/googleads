package googleads

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
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

func (c *Customer) SetResourceName(resourceName string) {
	c.Customer.ResourceName = resourceName
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
	resp, err := services.NewCustomerServiceClient(instance.conn).CreateCustomerClient(ctx, &services.CreateCustomerClientRequest{
		CustomerId:     parent.GetId(),
		CustomerClient: c.Customer,
	})
	if err != nil {
		return err
	}

	c.SetId(strings.Split(resp.GetResourceName(), "/")[1])
	c.SetResourceName(resp.GetResourceName())

	return nil
}

func (c *Customer) CreateBillingSetup(ctx context.Context, paymentsAccountId string) (*BillingSetup, error) {
	resp, err := services.NewBillingSetupServiceClient(instance.conn).MutateBillingSetup(ctx, &services.MutateBillingSetupRequest{
		CustomerId: c.GetId(),
		Operation: &services.BillingSetupOperation{
			Operation: &services.BillingSetupOperation_Create{
				Create: &resources.BillingSetup{
					PaymentsAccount: String(fmt.Sprintf("customers/%s/paymentsAccounts/%s", c.GetId(), paymentsAccountId)),
					StartTime: &resources.BillingSetup_StartTimeType{
						StartTimeType: enums.TimeTypeEnum_NOW,
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &BillingSetup{&resources.BillingSetup{
		ResourceName: resp.GetResult().GetResourceName(),
	}}, nil
}

func (c *Customer) CreateAccountBudget(ctx context.Context, bs *BillingSetup) (*AccountBudget, error) {
	resp, err := services.NewAccountBudgetProposalServiceClient(instance.conn).MutateAccountBudgetProposal(ctx, &services.MutateAccountBudgetProposalRequest{
		CustomerId: c.GetId(),
		Operation: &services.AccountBudgetProposalOperation{
			Operation: &services.AccountBudgetProposalOperation_Create{
				Create: &resources.AccountBudgetProposal{
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
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &AccountBudget{&resources.AccountBudgetProposal{
		ResourceName: resp.GetResult().GetResourceName(),
	}}, nil
}

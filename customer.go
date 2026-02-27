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
	Assets CustomerAssets
}

func NewCustomer() *Customer {
	return &Customer{
		Customer: &resources.Customer{
			CurrencyCode: String("EUR"),
			TimeZone:     String("Europe/Paris"),
		},
		Assets: NewCustomerAssets(),
	}
}

func (c Customer) GetId() string {
	return strconv.Itoa(int(c.Customer.GetId()))
}

func (c *Customer) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	c.Customer.Id = Int64(i)
}

func (c Customer) GetCurrency() string {
	return c.Customer.GetCurrencyCode()
}

func (c *Customer) SetCurrency(code string) {
	c.Customer.CurrencyCode = String(code)
}

func (c Customer) GetTimeZone() string {
	return c.Customer.GetTimeZone()
}

func (c *Customer) SetTimeZone(tz string) {
	c.Customer.TimeZone = String(tz)
}

func (c Customer) IsAutoTaggingEnabled() bool {
	return c.Customer.GetAutoTaggingEnabled()
}

func (c *Customer) SetAutoTagging(enabled bool) {
	c.Customer.AutoTaggingEnabled = Bool(enabled)
}

func (c Customer) IsTestAccount() bool {
	return c.Customer.GetTestAccount()
}

func (c *Customer) SetTestAccount(test bool) {
	c.Customer.TestAccount = Bool(test)
}

func (c Customer) GetTrackingUrl() string {
	return c.Customer.GetTrackingUrlTemplate()
}

func (c *Customer) SetTrackingUrl(url string) {
	c.Customer.TrackingUrlTemplate = String(url)
}

func (c Customer) GetFinalUrlSuffix() string {
	return c.Customer.GetFinalUrlSuffix()
}

func (c *Customer) SetFinalUrlSuffix(suffix string) {
	c.Customer.FinalUrlSuffix = String(suffix)
}

func (c Customer) IsManager() bool {
	return c.Customer.GetManager()
}

func (c *Customer) SetIsManager(manager bool) {
	c.Customer.Manager = Bool(manager)
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

func (c Customer) ListAssets(ctx context.Context) (CustomerAssets, error) {
	return ListCustomerAssets(ctx, c.GetId(), CustomerAssetByCustomer(c.GetResourceName()))
}

func (c *Customer) CreateAssets(ctx context.Context) error {
	tempId := newTempIdGenerator()

	ops := c.Assets.createOperations(c, tempId)
	if len(ops) == 0 {
		return nil
	}

	_, err := services.NewGoogleAdsServiceClient(instance.conn).Mutate(ctx, &services.MutateGoogleAdsRequest{
		CustomerId:       c.GetId(),
		MutateOperations: ops,
	})
	return err
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
	c.ResourceName = resp.GetResourceName()

	return nil
}

func (c *Customer) CreateBillingSetup(ctx context.Context, paymentsAccountId string) (*BillingSetup, error) {
	bs := &BillingSetup{&resources.BillingSetup{
		PaymentsAccount: String(fmt.Sprintf("customers/%s/paymentsAccounts/%s", c.GetId(), paymentsAccountId)),
		StartTime: &resources.BillingSetup_StartTimeType{
			StartTimeType: enums.TimeTypeEnum_NOW,
		},
	}}

	resp, err := services.NewBillingSetupServiceClient(instance.conn).MutateBillingSetup(ctx, &services.MutateBillingSetupRequest{
		CustomerId: c.GetId(),
		Operation: &services.BillingSetupOperation{
			Operation: &services.BillingSetupOperation_Create{
				Create: bs.BillingSetup,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	bs.SetId(strings.Split(resp.GetResult().GetResourceName(), "/")[3])
	bs.ResourceName = resp.GetResult().GetResourceName()

	return bs, nil
}

func (c *Customer) CreateAccountBudget(ctx context.Context, bs *BillingSetup) (*AccountBudget, error) {
	ab := &AccountBudget{&resources.AccountBudgetProposal{
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
	}}

	resp, err := services.NewAccountBudgetProposalServiceClient(instance.conn).MutateAccountBudgetProposal(ctx, &services.MutateAccountBudgetProposalRequest{
		CustomerId: c.GetId(),
		Operation: &services.AccountBudgetProposalOperation{
			Operation: &services.AccountBudgetProposalOperation_Create{
				Create: ab.AccountBudgetProposal,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	ab.SetId(strings.Split(resp.GetResult().GetResourceName(), "/")[3])
	ab.ResourceName = resp.GetResult().GetResourceName()

	return ab, nil
}

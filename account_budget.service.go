package googleads

import (
	"context"

	"github.com/shenzhencenter/google-ads-pb/services"
)

func FetchAccountBudget(ctx context.Context, customer *Customer, filters ...AccountBudgetFilter) (*AccountBudget, error) {
	return Fetch(ctx, services.NewGoogleAdsServiceClient(instance.conn), customer.GetId(), NewAccountBudgetQueryBuilder().Where(filters...).Build(), createAccountBudgetInstance)
}

func ListAccountBudgets(ctx context.Context, customer *Customer, filters ...AccountBudgetFilter) ([]*AccountBudget, error) {
	return List(ctx, services.NewGoogleAdsServiceClient(instance.conn), customer.GetId(), NewAccountBudgetQueryBuilder().Where(filters...).Build(), createAccountBudgetInstance)
}

func CreateAccountBudget(ctx context.Context, customer *Customer, ab *AccountBudget) (*AccountBudget, error) {
	return Create(ctx, services.NewAccountBudgetProposalServiceClient(instance.conn).MutateAccountBudgetProposal, &services.MutateAccountBudgetProposalRequest{
		CustomerId: customer.GetId(),
		Operation: &services.AccountBudgetProposalOperation{
			Operation: &services.AccountBudgetProposalOperation_Create{
				Create: ab.AccountBudgetProposal,
			},
		},
	}, func(customerId string, res *services.MutateAccountBudgetProposalResponse) string {
		return res.GetResult().GetResourceName()
	}, services.NewGoogleAdsServiceClient(instance.conn), NewAccountBudgetQueryBuilder(), AccountBudgetByResourceName, createAccountBudgetInstance)
}

func createAccountBudgetInstance(row *services.GoogleAdsRow) *AccountBudget {
	return &AccountBudget{row.GetAccountBudgetProposal()}
}

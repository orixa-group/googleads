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

func CreateAccountBudget(ctx context.Context, customer *Customer, bs *BillingSetup) (*AccountBudget, error) {
	return customer.CreateAccountBudget(ctx, bs)
}

func createAccountBudgetInstance(row *services.GoogleAdsRow) *AccountBudget {
	return &AccountBudget{row.GetAccountBudgetProposal()}
}

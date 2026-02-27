package googleads

import (
	"strconv"

	"github.com/shenzhencenter/google-ads-pb/resources"
)

type AccountBudget struct {
	*resources.AccountBudgetProposal
}

func NewAccountBudget() *AccountBudget {
	return &AccountBudget{&resources.AccountBudgetProposal{}}
}

func (ab *AccountBudget) GetId() string {
	return strconv.Itoa(int(ab.AccountBudgetProposal.GetId()))
}

func (ab *AccountBudget) SetId(id string) {
	i, _ := strconv.ParseInt(id, 10, 64)
	ab.AccountBudgetProposal.Id = Int64(i)
}

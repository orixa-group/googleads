package googleads

type CampaignBudgetQueryBuilder struct {
	*QueryBuilder[CampaignBudgetFilter]
}

func NewCampaignBudgetQueryBuilder() *CampaignBudgetQueryBuilder {
	return &CampaignBudgetQueryBuilder{NewQueryBuilder[CampaignBudgetFilter]().
		Select(
			"campaign_budget.type",
			"campaign_budget.total_amount_micros",
			"campaign_budget.status",
			"campaign_budget.resource_name",
			"campaign_budget.reference_count",
			"campaign_budget.recommended_budget_estimated_change_weekly_interactions",
			"campaign_budget.recommended_budget_estimated_change_weekly_views",
			"campaign_budget.recommended_budget_estimated_change_weekly_cost_micros",
			"campaign_budget.recommended_budget_estimated_change_weekly_clicks",
			"campaign_budget.recommended_budget_amount_micros",
			"campaign_budget.period",
			"campaign_budget.id",
			"campaign_budget.name",
			"campaign_budget.has_recommended_budget",
			"campaign_budget.explicitly_shared",
			"campaign_budget.delivery_method",
			"campaign_budget.amount_micros",
			"campaign_budget.aligned_bidding_strategy_id",
		).
		From("campaign_budget"),
	}
}

func (b *CampaignBudgetQueryBuilder) Where(clauses ...CampaignBudgetFilter) *CampaignBudgetQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

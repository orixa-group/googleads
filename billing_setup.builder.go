package googleads

type BillingSetupQueryBuilder struct {
	*QueryBuilder[BillingSetupFilter]
}

func NewBillingSetupQueryBuilder() *BillingSetupQueryBuilder {
	return &BillingSetupQueryBuilder{NewQueryBuilder[BillingSetupFilter]().
		Select(
			"billing_setup.end_date_time",
			"billing_setup.end_time_type",
			"billing_setup.id",
			"billing_setup.payments_account",
			"billing_setup.payments_account_info.payments_account_id",
			"billing_setup.status",
			"billing_setup.start_date_time",
			"billing_setup.resource_name",
			"billing_setup.payments_account_info.secondary_payments_profile_id",
			"billing_setup.payments_account_info.payments_profile_name",
			"billing_setup.payments_account_info.payments_profile_id",
			"billing_setup.payments_account_info.payments_account_name",
		).
		From("billing_setup")}
}

func (b *BillingSetupQueryBuilder) Where(clauses ...BillingSetupFilter) *BillingSetupQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

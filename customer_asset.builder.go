package googleads

type CustomerAssetQueryBuilder struct {
	*QueryBuilder[CustomerAssetFilter]
}

func NewCustomerAssetQueryBuilder() *CustomerAssetQueryBuilder {
	asset := NewAssetQueryBuilder()

	return &CustomerAssetQueryBuilder{NewQueryBuilder[CustomerAssetFilter]().
		Select(
			append(asset.fields,
				"customer_asset.asset",
				"customer_asset.field_type",
				"customer_asset.primary_status",
				"customer_asset.status",
				"customer_asset.source",
				"customer_asset.resource_name",
				"customer_asset.primary_status_reasons",
				"customer_asset.primary_status_details",
			)...,
		).
		From("customer_asset")}
}

func (b *CustomerAssetQueryBuilder) Where(clauses ...CustomerAssetFilter) *CustomerAssetQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

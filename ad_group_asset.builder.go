package googleads

type AdGroupAssetQueryBuilder struct {
	*QueryBuilder[AdGroupAssetFilter]
}

func NewAdGroupAssetQueryBuilder() *AdGroupAssetQueryBuilder {
	asset := NewAssetQueryBuilder()

	return &AdGroupAssetQueryBuilder{NewQueryBuilder[AdGroupAssetFilter]().
		Select(
			append(asset.fields,
				"ad_group_asset.status",
				"ad_group_asset.source",
				"ad_group_asset.resource_name",
				"ad_group_asset.primary_status_reasons",
				"ad_group_asset.primary_status_details",
				"ad_group_asset.primary_status",
				"ad_group_asset.field_type",
				"ad_group_asset.asset",
				"ad_group_asset.ad_group",
			)...,
		).
		From("ad_group_asset")}
}

func (b *AdGroupAssetQueryBuilder) Where(clauses ...AdGroupAssetFilter) *AdGroupAssetQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

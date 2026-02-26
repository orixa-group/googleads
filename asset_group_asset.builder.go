package googleads

type AssetGroupAssetQueryBuilder struct {
	*QueryBuilder[AssetGroupAssetFilter]
}

func NewAssetGroupAssetQueryBuilder() *AssetGroupAssetQueryBuilder {
	asset := NewAssetQueryBuilder()

	return &AssetGroupAssetQueryBuilder{NewQueryBuilder[AssetGroupAssetFilter]().
		Select(
			append(asset.fields,
				"asset_group_asset.asset",
				"asset_group_asset.asset_group",
				"asset_group_asset.field_type",
				"asset_group_asset.policy_summary.approval_status",
				"asset_group_asset.policy_summary.policy_topic_entries",
				"asset_group_asset.policy_summary.review_status",
				"asset_group_asset.primary_status",
				"asset_group_asset.primary_status_details",
				"asset_group_asset.primary_status_reasons",
				"asset_group_asset.resource_name",
				"asset_group_asset.source",
				"asset_group_asset.status",
			)...,
		).
		From("asset_group_asset"),
	}
}

func (b *AssetGroupAssetQueryBuilder) Where(clauses ...AssetGroupAssetFilter) *AssetGroupAssetQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

package googleads

type AssetGroupQueryBuilder struct {
	*QueryBuilder[AssetGroupFilter]
}

func NewAssetGroupQueryBuilder() *AssetGroupQueryBuilder {
	campaign := NewCampaignQueryBuilder()
	customer := NewCustomerQueryBuilder()

	return &AssetGroupQueryBuilder{NewQueryBuilder[AssetGroupFilter]().
		Select(
			append(append(campaign.fields, customer.fields...),
				"asset_group.status",
				"asset_group.resource_name",
				"asset_group.primary_status_reasons",
				"asset_group.primary_status",
				"asset_group.path2",
				"asset_group.path1",
				"asset_group.name",
				"asset_group.id",
				"asset_group.final_urls",
				"asset_group.final_mobile_urls",
				"asset_group.campaign",
				"asset_group.asset_coverage.ad_strength_action_items",
				"asset_group.ad_strength",
			)...,
		).
		From("asset_group"),
	}
}

func (b *AssetGroupQueryBuilder) Where(clauses ...AssetGroupFilter) *AssetGroupQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

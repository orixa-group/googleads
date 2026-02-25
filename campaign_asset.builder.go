package googleads

type CampaignAssetQueryBuilder struct {
	*QueryBuilder[CampaignAssetFilter]
}

func NewCampaignAssetQueryBuilder() *CampaignAssetQueryBuilder {
	asset := NewAssetQueryBuilder()

	return &CampaignAssetQueryBuilder{NewQueryBuilder[CampaignAssetFilter]().
		Select(
			append(asset.fields,
				"campaign_asset.status",
				"campaign_asset.source",
				"campaign_asset.resource_name",
				"campaign_asset.primary_status_reasons",
				"campaign_asset.primary_status",
				"campaign_asset.primary_status_details",
				"campaign_asset.field_type",
				"campaign_asset.campaign",
				"campaign_asset.asset",
			)...,
		).
		From("campaign_asset")}
}

func (b *CampaignAssetQueryBuilder) Where(clauses ...CampaignAssetFilter) *CampaignAssetQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

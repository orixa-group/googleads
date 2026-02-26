package googleads

type AdGroupAdQueryBuilder struct {
	*QueryBuilder[AdGroupAdFilter]
}

func NewAdGroupAdQueryBuilder() *AdGroupAdQueryBuilder {
	ad := NewAdQueryBuilder()

	return &AdGroupAdQueryBuilder{NewQueryBuilder[AdGroupAdFilter]().
		Select(
			append(Map(ad.fields, func(field string) string {
				return "ad_group_ad." + field
			}),
				"ad_group_ad.action_items",
				"ad_group_ad.ad_group",
				"ad_group_ad.ad_group_ad_asset_automation_settings",
				"ad_group_ad.ad_strength",
				"ad_group_ad.labels",
				"ad_group_ad.policy_summary.approval_status",
				"ad_group_ad.policy_summary.policy_topic_entries",
				"ad_group_ad.policy_summary.review_status",
				"ad_group_ad.primary_status",
				"ad_group_ad.primary_status_reasons",
				"ad_group_ad.resource_name",
				"ad_group_ad.status",
			)...,
		).
		From("ad_group_ad")}
}

func (b *AdGroupAdQueryBuilder) Where(clauses ...AdGroupAdFilter) *AdGroupAdQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

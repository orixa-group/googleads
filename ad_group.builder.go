package googleads

type AdGroupQueryBuilder struct {
	*QueryBuilder[AdGroupFilter]
}

func NewAdGroupQueryBuilder() *AdGroupQueryBuilder {
	campaign := NewCampaignQueryBuilder()
	customer := NewCustomerQueryBuilder()

	return &AdGroupQueryBuilder{NewQueryBuilder[AdGroupFilter]().
		Select(
			append(append(campaign.fields, customer.fields...),
				"ad_group.video_ad_group_settings.video_ad_sequence.step_id",
				"ad_group.url_custom_parameters",
				"ad_group.type",
				"ad_group.tracking_url_template",
				"ad_group.targeting_setting.target_restrictions",
				"ad_group.target_roas",
				"ad_group.target_cpv_micros",
				"ad_group.target_cpm_micros",
				"ad_group.target_cpc_micros",
				"ad_group.target_cpa_micros",
				"ad_group.status",
				"ad_group.resource_name",
				"ad_group.primary_status_reasons",
				"ad_group.primary_status",
				"ad_group.optimized_targeting_enabled",
				"ad_group.percent_cpc_bid_micros",
				"ad_group.name",
				"ad_group.labels",
				"ad_group.id",
				"ad_group.final_url_suffix",
				"ad_group.fixed_cpm_micros",
				"ad_group.excluded_parent_asset_set_types",
				"ad_group.excluded_parent_asset_field_types",
				"ad_group.exclude_demographic_expansion",
				"ad_group.effective_target_roas_source",
				"ad_group.effective_target_roas",
				"ad_group.effective_target_cpc_source",
				"ad_group.effective_target_cpc",
				"ad_group.effective_target_cpa_source",
				"ad_group.effective_target_cpa_micros",
				"ad_group.display_custom_bid_dimension",
				"ad_group.effective_cpc_bid_micros",
				"ad_group.demand_gen_ad_group_settings.channel_controls.selected_channels.youtube_shorts",
				"ad_group.demand_gen_ad_group_settings.channel_controls.selected_channels.youtube_in_stream",
				"ad_group.demand_gen_ad_group_settings.channel_controls.selected_channels.youtube_in_feed",
				"ad_group.demand_gen_ad_group_settings.channel_controls.selected_channels.gmail",
				"ad_group.demand_gen_ad_group_settings.channel_controls.selected_channels.display",
				"ad_group.demand_gen_ad_group_settings.channel_controls.selected_channels.discover",
				"ad_group.demand_gen_ad_group_settings.channel_controls.channel_strategy",
				"ad_group.demand_gen_ad_group_settings.channel_controls.channel_config",
				"ad_group.cpv_bid_micros",
				"ad_group.cpm_bid_micros",
				"ad_group.cpc_bid_micros",
				"ad_group.campaign",
				"ad_group.base_ad_group",
				"ad_group.audience_setting.use_audience_grouped",
				"ad_group.ai_max_ad_group_setting.disable_search_term_matching",
				"ad_group.ad_rotation_mode",
			)...,
		).
		From("ad_group"),
	}
}

func (b *AdGroupQueryBuilder) Where(clauses ...AdGroupFilter) *AdGroupQueryBuilder {
	b.QueryBuilder.Where(clauses...)
	return b
}

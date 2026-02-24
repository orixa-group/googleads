package googleads

type CustomerQueryBuilder struct {
	*QueryBuilder[Filter]
}

func NewCustomerQueryBuilder() *CustomerQueryBuilder {
	return &CustomerQueryBuilder{NewQueryBuilder[Filter]().
		Select(
			"customer.video_customer.third_party_integration_partners.viewability_integration_partners",
			"customer.video_customer.third_party_integration_partners.reach_integration_partners",
			"customer.video_customer.third_party_integration_partners.brand_safety_integration_partners",
			"customer.video_customer.third_party_integration_partners.brand_lift_integration_partners",
			"customer.video_brand_safety_suitability",
			"customer.tracking_url_template",
			"customer.time_zone",
			"customer.test_account",
			"customer.status",
			"customer.resource_name",
			"customer.remarketing_setting.google_global_site_tag",
			"customer.pay_per_conversion_eligibility_failure_reasons",
			"customer.optimization_score_weight",
			"customer.optimization_score",
			"customer.manager",
			"customer.location_asset_auto_migration_done_date_time",
			"customer.location_asset_auto_migration_done",
			"customer.local_services_settings.granular_license_statuses",
			"customer.local_services_settings.granular_insurance_statuses",
			"customer.image_asset_auto_migration_done_date_time",
			"customer.image_asset_auto_migration_done",
			"customer.id",
			"customer.has_partners_badge",
			"customer.final_url_suffix",
			"customer.descriptive_name",
			"customer.currency_code",
			"customer.conversion_tracking_setting.google_ads_conversion_customer",
			"customer.customer_agreement_setting.accepted_lead_form_terms",
			"customer.conversion_tracking_setting.enhanced_conversions_for_leads_enabled",
			"customer.conversion_tracking_setting.conversion_tracking_status",
			"customer.conversion_tracking_setting.cross_account_conversion_tracking_id",
			"customer.conversion_tracking_setting.conversion_tracking_id",
			"customer.conversion_tracking_setting.accepted_customer_data_terms",
			"customer.call_reporting_setting.call_reporting_enabled",
			"customer.call_reporting_setting.call_conversion_reporting_enabled",
			"customer.call_reporting_setting.call_conversion_action",
			"customer.auto_tagging_enabled",
		).
		From("customer")}
}

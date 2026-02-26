package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type CampaignKeywordMatchType func(campaign *resources.Campaign)

func CampaignKeywordMatchBroad(campaign *resources.Campaign) {
	campaign.KeywordMatchType = enums.CampaignKeywordMatchTypeEnum_BROAD
}

type GeoTargetTypeSetting func(setting *resources.Campaign_GeoTargetTypeSetting)

func GeoTargetPositivePresenceOrInterest(setting *resources.Campaign_GeoTargetTypeSetting) {
	setting.PositiveGeoTargetType = enums.PositiveGeoTargetTypeEnum_PRESENCE_OR_INTEREST
}

func GeoTargetPositivePresence(setting *resources.Campaign_GeoTargetTypeSetting) {
	setting.PositiveGeoTargetType = enums.PositiveGeoTargetTypeEnum_PRESENCE
}

func GeoTargetNegativePresence(setting *resources.Campaign_GeoTargetTypeSetting) {
	setting.NegativeGeoTargetType = enums.NegativeGeoTargetTypeEnum_PRESENCE
}

func GeoTargetNegativePresenceOrInterest(setting *resources.Campaign_GeoTargetTypeSetting) {
	setting.NegativeGeoTargetType = enums.NegativeGeoTargetTypeEnum_PRESENCE_OR_INTEREST
}

type AdServingOptimizationStatus func(campaign *resources.Campaign)

func AdServingOptimize(campaign *resources.Campaign) {
	campaign.AdServingOptimizationStatus = enums.AdServingOptimizationStatusEnum_OPTIMIZE
}

func AdServingRotateIndefinitely(campaign *resources.Campaign) {
	campaign.AdServingOptimizationStatus = enums.AdServingOptimizationStatusEnum_ROTATE_INDEFINITELY
}

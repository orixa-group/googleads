package googleads

import (
	"strings"

	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

const (
	ChannelTypeSearch         ChannelType = "SEARCH"
	ChannelTypePerformanceMax ChannelType = "PERFORMANCE_MAX"
)

type ChannelType string

func (ct ChannelType) is(channelType ChannelType) bool {
	return strings.EqualFold(ct.String(), channelType.String())
}

func (ct ChannelType) String() string {
	return string(ct)
}

var channelTypeToEnum = map[ChannelType]enums.AdvertisingChannelTypeEnum_AdvertisingChannelType{
	ChannelTypeSearch:         enums.AdvertisingChannelTypeEnum_SEARCH,
	ChannelTypePerformanceMax: enums.AdvertisingChannelTypeEnum_PERFORMANCE_MAX,
}

var enumToChannelType = map[enums.AdvertisingChannelTypeEnum_AdvertisingChannelType]ChannelType{
	enums.AdvertisingChannelTypeEnum_SEARCH:          ChannelTypeSearch,
	enums.AdvertisingChannelTypeEnum_PERFORMANCE_MAX: ChannelTypePerformanceMax,
}

type Objective string

const (
	ObjectiveConversions Objective = "CONVERSIONS"
	ObjectiveClicks      Objective = "CLICKS"
)

func (o Objective) is(objective Objective) bool {
	return strings.EqualFold(o.String(), objective.String())
}

func (o Objective) String() string {
	return string(o)
}

var objectiveToEnum = map[Objective]func(c *Campaign){
	ObjectiveConversions: func(c *Campaign) {
		c.BiddingStrategyType = enums.BiddingStrategyTypeEnum_MAXIMIZE_CONVERSIONS
		c.CampaignBiddingStrategy = &resources.Campaign_MaximizeConversions{
			MaximizeConversions: &common.MaximizeConversions{},
		}
	},
	ObjectiveClicks: func(c *Campaign) {
		c.BiddingStrategyType = enums.BiddingStrategyTypeEnum_TARGET_SPEND
		c.CampaignBiddingStrategy = &resources.Campaign_TargetSpend{
			TargetSpend: &common.TargetSpend{},
		}
	},
}

var enumToObjective = map[enums.BiddingStrategyTypeEnum_BiddingStrategyType]Objective{
	enums.BiddingStrategyTypeEnum_MAXIMIZE_CONVERSIONS: ObjectiveConversions,
	enums.BiddingStrategyTypeEnum_TARGET_SPEND:         ObjectiveClicks,
}

type GeoTargetType string

const (
	GeoTargetTypeInterest          GeoTargetType = "INTEREST"
	GeoTargetTypePresenceOrInterest GeoTargetType = "PRESENCE_OR_INTEREST"
)

func (g GeoTargetType) is(geoTargetType GeoTargetType) bool {
	return strings.EqualFold(g.String(), geoTargetType.String())
}

func (g GeoTargetType) String() string {
	return string(g)
}

var geoTargetTypeToEnum = map[GeoTargetType]enums.PositiveGeoTargetTypeEnum_PositiveGeoTargetType{
	GeoTargetTypeInterest:           enums.PositiveGeoTargetTypeEnum_SEARCH_INTEREST,
	GeoTargetTypePresenceOrInterest: enums.PositiveGeoTargetTypeEnum_PRESENCE_OR_INTEREST,
}

var enumToGeoTargetType = map[enums.PositiveGeoTargetTypeEnum_PositiveGeoTargetType]GeoTargetType{
	enums.PositiveGeoTargetTypeEnum_SEARCH_INTEREST:      GeoTargetTypeInterest,
	enums.PositiveGeoTargetTypeEnum_PRESENCE_OR_INTEREST: GeoTargetTypePresenceOrInterest,
}

type AssetAutomationType string

const (
	AssetAutomationTypeTextAssets              AssetAutomationType = "TEXT_ASSET_AUTOMATION"
	AssetAutomationTypeVerticalYouTubeVideos   AssetAutomationType = "GENERATE_VERTICAL_YOUTUBE_VIDEOS"
	AssetAutomationTypeShorterYouTubeVideos    AssetAutomationType = "GENERATE_SHORTER_YOUTUBE_VIDEOS"
	AssetAutomationTypeLandingPagePreview      AssetAutomationType = "GENERATE_LANDING_PAGE_PREVIEW"
	AssetAutomationTypeEnhancedYouTubeVideos   AssetAutomationType = "GENERATE_ENHANCED_YOUTUBE_VIDEOS"
	AssetAutomationTypeImageEnhancement        AssetAutomationType = "GENERATE_IMAGE_ENHANCEMENT"
	AssetAutomationTypeImageExtraction         AssetAutomationType = "GENERATE_IMAGE_EXTRACTION"
	AssetAutomationTypeDesignVersionsForImages AssetAutomationType = "GENERATE_DESIGN_VERSIONS_FOR_IMAGES"
	AssetAutomationTypeFinalURLExpansion       AssetAutomationType = "FINAL_URL_EXPANSION_TEXT_ASSET_AUTOMATION"
	AssetAutomationTypeVideosFromOtherAssets   AssetAutomationType = "GENERATE_VIDEOS_FROM_OTHER_ASSETS"
)

func (a AssetAutomationType) String() string {
	return string(a)
}

var assetAutomationTypeToEnum = map[AssetAutomationType]enums.AssetAutomationTypeEnum_AssetAutomationType{
	AssetAutomationTypeTextAssets:              enums.AssetAutomationTypeEnum_TEXT_ASSET_AUTOMATION,
	AssetAutomationTypeVerticalYouTubeVideos:   enums.AssetAutomationTypeEnum_GENERATE_VERTICAL_YOUTUBE_VIDEOS,
	AssetAutomationTypeShorterYouTubeVideos:    enums.AssetAutomationTypeEnum_GENERATE_SHORTER_YOUTUBE_VIDEOS,
	AssetAutomationTypeLandingPagePreview:      enums.AssetAutomationTypeEnum_GENERATE_LANDING_PAGE_PREVIEW,
	AssetAutomationTypeEnhancedYouTubeVideos:   enums.AssetAutomationTypeEnum_GENERATE_ENHANCED_YOUTUBE_VIDEOS,
	AssetAutomationTypeImageEnhancement:        enums.AssetAutomationTypeEnum_GENERATE_IMAGE_ENHANCEMENT,
	AssetAutomationTypeImageExtraction:         enums.AssetAutomationTypeEnum_GENERATE_IMAGE_EXTRACTION,
	AssetAutomationTypeDesignVersionsForImages: enums.AssetAutomationTypeEnum_GENERATE_DESIGN_VERSIONS_FOR_IMAGES,
	AssetAutomationTypeFinalURLExpansion:       enums.AssetAutomationTypeEnum_FINAL_URL_EXPANSION_TEXT_ASSET_AUTOMATION,
	AssetAutomationTypeVideosFromOtherAssets:   enums.AssetAutomationTypeEnum_GENERATE_VIDEOS_FROM_OTHER_ASSETS,
}

var enumToAssetAutomationType = map[enums.AssetAutomationTypeEnum_AssetAutomationType]AssetAutomationType{
	enums.AssetAutomationTypeEnum_TEXT_ASSET_AUTOMATION:                    AssetAutomationTypeTextAssets,
	enums.AssetAutomationTypeEnum_GENERATE_VERTICAL_YOUTUBE_VIDEOS:         AssetAutomationTypeVerticalYouTubeVideos,
	enums.AssetAutomationTypeEnum_GENERATE_SHORTER_YOUTUBE_VIDEOS:          AssetAutomationTypeShorterYouTubeVideos,
	enums.AssetAutomationTypeEnum_GENERATE_LANDING_PAGE_PREVIEW:            AssetAutomationTypeLandingPagePreview,
	enums.AssetAutomationTypeEnum_GENERATE_ENHANCED_YOUTUBE_VIDEOS:         AssetAutomationTypeEnhancedYouTubeVideos,
	enums.AssetAutomationTypeEnum_GENERATE_IMAGE_ENHANCEMENT:               AssetAutomationTypeImageEnhancement,
	enums.AssetAutomationTypeEnum_GENERATE_IMAGE_EXTRACTION:                AssetAutomationTypeImageExtraction,
	enums.AssetAutomationTypeEnum_GENERATE_DESIGN_VERSIONS_FOR_IMAGES:      AssetAutomationTypeDesignVersionsForImages,
	enums.AssetAutomationTypeEnum_FINAL_URL_EXPANSION_TEXT_ASSET_AUTOMATION: AssetAutomationTypeFinalURLExpansion,
	enums.AssetAutomationTypeEnum_GENERATE_VIDEOS_FROM_OTHER_ASSETS:        AssetAutomationTypeVideosFromOtherAssets,
}

type AssetAutomationStatus string

const (
	AssetAutomationStatusOptedIn  AssetAutomationStatus = "OPTED_IN"
	AssetAutomationStatusOptedOut AssetAutomationStatus = "OPTED_OUT"
)

func (a AssetAutomationStatus) String() string {
	return string(a)
}

var assetAutomationStatusToEnum = map[AssetAutomationStatus]enums.AssetAutomationStatusEnum_AssetAutomationStatus{
	AssetAutomationStatusOptedIn:  enums.AssetAutomationStatusEnum_OPTED_IN,
	AssetAutomationStatusOptedOut: enums.AssetAutomationStatusEnum_OPTED_OUT,
}

var enumToAssetAutomationStatus = map[enums.AssetAutomationStatusEnum_AssetAutomationStatus]AssetAutomationStatus{
	enums.AssetAutomationStatusEnum_OPTED_IN:  AssetAutomationStatusOptedIn,
	enums.AssetAutomationStatusEnum_OPTED_OUT: AssetAutomationStatusOptedOut,
}

type AssetAutomationSetting struct {
	Type   AssetAutomationType
	Status AssetAutomationStatus
}

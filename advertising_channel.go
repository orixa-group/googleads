package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type AdvertisingChannel func(campaign *resources.Campaign)

func AdvertisingChannelSearch(campaign *resources.Campaign) {
	campaign.AdvertisingChannelType = enums.AdvertisingChannelTypeEnum_SEARCH
}

func AdvertisingChannelDisplay(campaign *resources.Campaign) {
	campaign.AdvertisingChannelType = enums.AdvertisingChannelTypeEnum_DISPLAY
}

func AdvertisingChannelVideo(campaign *resources.Campaign) {
	campaign.AdvertisingChannelType = enums.AdvertisingChannelTypeEnum_VIDEO
}

func AdvertisingChannelPerformanceMax(campaign *resources.Campaign) {
	campaign.AdvertisingChannelType = enums.AdvertisingChannelTypeEnum_PERFORMANCE_MAX
}

func AdvertisingChannelShopping(campaign *resources.Campaign) {
	campaign.AdvertisingChannelType = enums.AdvertisingChannelTypeEnum_SHOPPING
}

func AdvertisingChannelDemandGen(campaign *resources.Campaign) {
	campaign.AdvertisingChannelType = enums.AdvertisingChannelTypeEnum_DEMAND_GEN
}

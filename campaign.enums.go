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

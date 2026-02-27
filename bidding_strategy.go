package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type BiddingStrategy func(campaign *resources.Campaign)

func BiddingManualCpc(enhanced bool) BiddingStrategy {
	return func(campaign *resources.Campaign) {
		campaign.CampaignBiddingStrategy = &resources.Campaign_ManualCpc{
			ManualCpc: &common.ManualCpc{
				EnhancedCpcEnabled: Bool(enhanced),
			},
		}
	}
}

func BiddingMaximizeConversions(targetCpaCents int) BiddingStrategy {
	return func(campaign *resources.Campaign) {
		m := &common.MaximizeConversions{}
		if targetCpaCents > 0 {
			m.TargetCpaMicros = int64(targetCpaCents * 10_000)
		}
		campaign.CampaignBiddingStrategy = &resources.Campaign_MaximizeConversions{
			MaximizeConversions: m,
		}
	}
}

func BiddingMaximizeConversionValue(targetRoas float64) BiddingStrategy {
	return func(campaign *resources.Campaign) {
		m := &common.MaximizeConversionValue{}
		if targetRoas > 0 {
			m.TargetRoas = targetRoas
		}
		campaign.CampaignBiddingStrategy = &resources.Campaign_MaximizeConversionValue{
			MaximizeConversionValue: m,
		}
	}
}

func BiddingTargetCpa(targetCpaCents int) BiddingStrategy {
	return func(campaign *resources.Campaign) {
		campaign.CampaignBiddingStrategy = &resources.Campaign_TargetCpa{
			TargetCpa: &common.TargetCpa{
				TargetCpaMicros: Int64(int64(targetCpaCents * 10_000)),
			},
		}
	}
}

func BiddingTargetRoas(targetRoas float64) BiddingStrategy {
	return func(campaign *resources.Campaign) {
		campaign.CampaignBiddingStrategy = &resources.Campaign_TargetRoas{
			TargetRoas: &common.TargetRoas{
				TargetRoas: Float64(targetRoas),
			},
		}
	}
}

func BiddingTargetSpend() BiddingStrategy {
	return func(campaign *resources.Campaign) {
		campaign.CampaignBiddingStrategy = &resources.Campaign_TargetSpend{
			TargetSpend: &common.TargetSpend{},
		}
	}
}

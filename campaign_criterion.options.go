package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type CampaignCriterionOption func(*resources.CampaignCriterion)

func ChangeCampaignCriterionProximityRadius(address string, km float64) CampaignCriterionOption {
	return func(cc *resources.CampaignCriterion) {
		if criterion, ok := cc.GetCriterion().(*resources.CampaignCriterion_Proximity); ok {
			criterion.Proximity = &common.ProximityInfo{
				Radius:      &km,
				RadiusUnits: enums.ProximityRadiusUnitsEnum_KILOMETERS,
				Address: &common.AddressInfo{
					StreetAddress: String(address),
				},
			}
		}
	}
}

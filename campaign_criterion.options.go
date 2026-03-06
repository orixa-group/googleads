package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type CampaignCriterionOption func(*resources.CampaignCriterion)

func ChangeCampaignCriterionProximityRadius(lat, long int32, km float64) CampaignCriterionOption {
	return func(cc *resources.CampaignCriterion) {
		if criterion, ok := cc.GetCriterion().(*resources.CampaignCriterion_Proximity); ok {
			criterion.Proximity = &common.ProximityInfo{
				Radius:      &km,
				RadiusUnits: enums.ProximityRadiusUnitsEnum_KILOMETERS,
				GeoPoint: &common.GeoPointInfo{
					LongitudeInMicroDegrees: &long,
					LatitudeInMicroDegrees:  &lat,
				},
			}
		}
	}
}

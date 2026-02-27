package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type CampaignCriterion struct {
	*resources.CampaignCriterion
}

func NewCampaignCriterion() *CampaignCriterion {
	return &CampaignCriterion{
		CampaignCriterion: &resources.CampaignCriterion{},
	}
}

func (c *CampaignCriterion) createOperation(campaign *Campaign) *services.MutateOperation {
	c.Campaign = String(campaign.GetResourceName())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_CampaignCriterionOperation{
			CampaignCriterionOperation: &services.CampaignCriterionOperation{
				Operation: &services.CampaignCriterionOperation_Create{
					Create: c.CampaignCriterion,
				},
			},
		},
	}
}

type CampaignCriteria []*CampaignCriterion

func NewCampaignCriteria() CampaignCriteria {
	return make(CampaignCriteria, 0)
}

func (c *CampaignCriteria) Add(criterion *CampaignCriterion) {
	*c = append(*c, &CampaignCriterion{&resources.CampaignCriterion{
		Criterion: criterion.GetCriterion(),
	}})
}

func (c *CampaignCriteria) AddKeyword(keyword string, matchType KeywordMatchType) {
	k := &common.KeywordInfo{
		Text: String(keyword),
	}
	matchType(k)

	c.Add(&CampaignCriterion{&resources.CampaignCriterion{
		Criterion: &resources.CampaignCriterion_Keyword{
			Keyword: k,
		},
	}})
}

func (c CampaignCriteria) createOperations(campaign *Campaign) []*services.MutateOperation {
	return Map(c, func(item *CampaignCriterion) *services.MutateOperation {
		return item.createOperation(campaign)
	})
}

func (c *CampaignCriteria) AddDevice(deviceType DeviceType) {
	d := &common.DeviceInfo{}
	deviceType(d)

	c.Add(&CampaignCriterion{&resources.CampaignCriterion{
		Criterion: &resources.CampaignCriterion_Device{
			Device: d,
		},
	}})
}

func (c *CampaignCriteria) AddProximityByCoordinates(latitude, longitude float64, radius float64) {
	c.Add(&CampaignCriterion{&resources.CampaignCriterion{
		Criterion: &resources.CampaignCriterion_Proximity{
			Proximity: &common.ProximityInfo{
				GeoPoint: &common.GeoPointInfo{
					LatitudeInMicroDegrees:  Int32(int32(latitude * 1000000)),
					LongitudeInMicroDegrees: Int32(int32(longitude * 1000000)),
				},
				Radius:      Float64(radius),
				RadiusUnits: enums.ProximityRadiusUnitsEnum_KILOMETERS,
			},
		},
	}})
}

func (c *CampaignCriteria) AddProximityByAddress(address string, radius float64) {
	c.Add(&CampaignCriterion{&resources.CampaignCriterion{
		Criterion: &resources.CampaignCriterion_Proximity{
			Proximity: &common.ProximityInfo{
				Address: &common.AddressInfo{
					StreetAddress: String(address),
				},
				Radius:      Float64(radius),
				RadiusUnits: enums.ProximityRadiusUnitsEnum_KILOMETERS,
			},
		},
	}})
}

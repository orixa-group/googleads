package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
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

func (c CampaignCriteria) Add(criterion *CampaignCriterion) {
	c = append(c, &CampaignCriterion{&resources.CampaignCriterion{
		Criterion: criterion.GetCriterion(),
	}})
}

func (c CampaignCriteria) AddKeyword(keyword string, matchType KeywordMatchType) {
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

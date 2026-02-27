package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AdGroupCriterion struct {
	*resources.AdGroupCriterion
}

func NewAdGroupCriterion() *AdGroupCriterion {
	return &AdGroupCriterion{&resources.AdGroupCriterion{}}
}

func (agc *AdGroupCriterion) createOperation(adGroup *AdGroup) *services.MutateOperation {
	agc.AdGroup = String(adGroup.GetResourceName())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_AdGroupCriterionOperation{
			AdGroupCriterionOperation: &services.AdGroupCriterionOperation{
				Operation: &services.AdGroupCriterionOperation_Create{
					Create: agc.AdGroupCriterion,
				},
			},
		},
	}
}

type AdGroupCriteria []*AdGroupCriterion

func NewAdGroupCriteria() AdGroupCriteria {
	return make(AdGroupCriteria, 0)
}

func (agcs *AdGroupCriteria) Add(criterion *AdGroupCriterion) {
	*agcs = append(*agcs, &AdGroupCriterion{&resources.AdGroupCriterion{
		Criterion: criterion.GetCriterion(),
	}})
}

func (agcs *AdGroupCriteria) AddKeyword(keyword string, matchType KeywordMatchType) {
	k := &common.KeywordInfo{
		Text: String(keyword),
	}
	matchType(k)

	agcs.Add(&AdGroupCriterion{&resources.AdGroupCriterion{
		Criterion: &resources.AdGroupCriterion_Keyword{
			Keyword: k,
		},
	}})
}

func (agcs AdGroupCriteria) createOperations(adGroup *AdGroup) []*services.MutateOperation {
	return Map(agcs, func(item *AdGroupCriterion) *services.MutateOperation {
		return item.createOperation(adGroup)
	})
}

package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/resources"
)

type AdGroupCriterion struct {
	*resources.AdGroupCriterion
}

func NewAdGroupCriterion() *AdGroupCriterion {
	return &AdGroupCriterion{&resources.AdGroupCriterion{}}
}

type AdGroupCriteria []*AdGroupCriterion

func NewAdGroupCriteria() AdGroupCriteria {
	return make(AdGroupCriteria, 0)
}

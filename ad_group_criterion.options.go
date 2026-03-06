package googleads

import (
	"strings"

	"github.com/shenzhencenter/google-ads-pb/resources"
)

type AdGroupCriterionOption func(*resources.AdGroupCriterion)

func ReplaceWordInKeyword(old, new string) AdGroupCriterionOption {
	return func(c *resources.AdGroupCriterion) {
		c.Criterion.(*resources.AdGroupCriterion_Keyword).Keyword.Text = String(strings.ReplaceAll(c.Criterion.(*resources.AdGroupCriterion_Keyword).Keyword.GetText(), old, new))
	}
}

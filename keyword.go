package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
)

type KeywordMatchType func(keyword *common.KeywordInfo)

func KeywordMatchExact(keyword *common.KeywordInfo) {
	keyword.MatchType = enums.KeywordMatchTypeEnum_EXACT
}

func KeywordMatchPhrase(keyword *common.KeywordInfo) {
	keyword.MatchType = enums.KeywordMatchTypeEnum_PHRASE
}

func KeywordMatchBroad(keyword *common.KeywordInfo) {
	keyword.MatchType = enums.KeywordMatchTypeEnum_BROAD
}

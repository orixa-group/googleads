package googleads

import (
	"strings"

	"github.com/shenzhencenter/google-ads-pb/enums"
)

const (
	KeywordMatchTypePhrase KeywordMatchType = "PHRASE"
	KeywordMatchTypeExact  KeywordMatchType = "EXACT"
	KeywordMatchTypeBroad  KeywordMatchType = "BROAD"
)

type KeywordMatchType string

func (k KeywordMatchType) is(matchType KeywordMatchType) bool {
	return strings.EqualFold(k.String(), matchType.String())
}

func (k KeywordMatchType) String() string {
	return string(k)
}

var keywordMatchTypeToEnum = map[KeywordMatchType]enums.KeywordMatchTypeEnum_KeywordMatchType{
	KeywordMatchTypePhrase: enums.KeywordMatchTypeEnum_PHRASE,
	KeywordMatchTypeExact:  enums.KeywordMatchTypeEnum_EXACT,
	KeywordMatchTypeBroad:  enums.KeywordMatchTypeEnum_BROAD,
}

var enumToKeywordMatchType = map[enums.KeywordMatchTypeEnum_KeywordMatchType]KeywordMatchType{
	enums.KeywordMatchTypeEnum_PHRASE: KeywordMatchTypePhrase,
	enums.KeywordMatchTypeEnum_EXACT:  KeywordMatchTypeExact,
	enums.KeywordMatchTypeEnum_BROAD:  KeywordMatchTypeBroad,
}

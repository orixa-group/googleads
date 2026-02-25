package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type Asset struct {
	*resources.Asset
}

func NewAsset() *Asset {
	return &Asset{&resources.Asset{}}
}

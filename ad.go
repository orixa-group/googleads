package googleads

import "github.com/shenzhencenter/google-ads-pb/resources"

type Ad struct {
	*resources.Ad
}

func NewAd() *Ad {
	return &Ad{&resources.Ad{}}
}

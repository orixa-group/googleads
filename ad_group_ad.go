package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AdGroupAd struct {
	*resources.AdGroupAd
}

func NewAdGroupAd() *AdGroupAd {
	return &AdGroupAd{&resources.AdGroupAd{}}
}

func (aga *AdGroupAd) createOperation(adGroup *AdGroup) *services.MutateOperation {
	aga.AdGroup = String(adGroup.GetResourceName())

	return &services.MutateOperation{
		Operation: &services.MutateOperation_AdGroupAdOperation{
			AdGroupAdOperation: &services.AdGroupAdOperation{
				Operation: &services.AdGroupAdOperation_Create{
					Create: aga.AdGroupAd,
				},
			},
		},
	}
}

type AdGroupAds []*AdGroupAd

func NewAdGroupAds() AdGroupAds {
	return make(AdGroupAds, 0)
}

func (agas *AdGroupAds) Add(ad *resources.Ad) {
	*agas = append(*agas, &AdGroupAd{
		AdGroupAd: &resources.AdGroupAd{
			Ad: ad,
		},
	})
}

func (agas *AdGroupAds) AddResponsiveSearchAd(headlines []string, descriptions []string, finalUrls ...string) {
	agas.Add(&resources.Ad{
		FinalUrls: finalUrls,
		AdData: &resources.Ad_ResponsiveSearchAd{
			ResponsiveSearchAd: &common.ResponsiveSearchAdInfo{
				Headlines: Map(headlines, func(text string) *common.AdTextAsset {
					return &common.AdTextAsset{Text: String(text)}
				}),
				Descriptions: Map(descriptions, func(text string) *common.AdTextAsset {
					return &common.AdTextAsset{Text: String(text)}
				}),
			},
		},
	})
}

func (agas *AdGroupAds) AddResponsiveDisplayAd(headlines []string, descriptions []string, businessName string, finalUrls ...string) {
	agas.Add(&resources.Ad{
		FinalUrls: finalUrls,
		AdData: &resources.Ad_ResponsiveDisplayAd{
			ResponsiveDisplayAd: &common.ResponsiveDisplayAdInfo{
				Headlines: Map(headlines, func(text string) *common.AdTextAsset {
					return &common.AdTextAsset{Text: String(text)}
				}),
				Descriptions: Map(descriptions, func(text string) *common.AdTextAsset {
					return &common.AdTextAsset{Text: String(text)}
				}),
				BusinessName: String(businessName),
			},
		},
	})
}

func (agas *AdGroupAds) AddAppAd(headlines []string, descriptions []string, finalUrls ...string) {
	agas.Add(&resources.Ad{
		FinalUrls: finalUrls,
		AdData: &resources.Ad_AppAd{
			AppAd: &common.AppAdInfo{
				Headlines: Map(headlines, func(text string) *common.AdTextAsset {
					return &common.AdTextAsset{Text: String(text)}
				}),
				Descriptions: Map(descriptions, func(text string) *common.AdTextAsset {
					return &common.AdTextAsset{Text: String(text)}
				}),
			},
		},
	})
}

func (agas AdGroupAds) createOperations(adGroup *AdGroup) []*services.MutateOperation {
	return Map(agas, func(item *AdGroupAd) *services.MutateOperation {
		return item.createOperation(adGroup)
	})
}

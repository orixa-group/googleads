package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AdGroupAsset struct {
	*resources.AdGroupAsset
	Asset *Asset
}

func (aga *AdGroupAsset) createOperations(adGroup *AdGroup, tempId tempIdGenerator) []*services.MutateOperation {
	aop := aga.Asset.createOperation(adGroup.Campaign.Customer, tempId)

	aga.AdGroupAsset.AdGroup = adGroup.GetResourceName()
	aga.AdGroupAsset.Asset = aga.Asset.GetResourceName()

	return []*services.MutateOperation{
		aop,
		{
			Operation: &services.MutateOperation_AdGroupAssetOperation{
				AdGroupAssetOperation: &services.AdGroupAssetOperation{
					Operation: &services.AdGroupAssetOperation_Create{
						Create: aga.AdGroupAsset,
					},
				},
			},
		},
	}
}

type AdGroupAssets []*AdGroupAsset

func (agas *AdGroupAssets) Add(asset *AdGroupAsset, options ...AssetOption) {
	a := &resources.Asset{
		AssetData: asset.Asset.GetAssetData(),
		FinalUrls: asset.Asset.GetFinalUrls(),
		Name:      String(asset.Asset.GetName()),
	}
	for _, opt := range options {
		opt(a)
	}

	*agas = append(*agas, &AdGroupAsset{&resources.AdGroupAsset{
		FieldType: asset.GetFieldType(),
	}, &Asset{a}})
}

func (agas *AdGroupAssets) AddSitelink(text, description1, description2 string, finalUrls ...string) {
	agas.Add(&AdGroupAsset{
		AdGroupAsset: &resources.AdGroupAsset{
			FieldType: enums.AssetFieldTypeEnum_SITELINK,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				FinalUrls: finalUrls,
				AssetData: &resources.Asset_SitelinkAsset{
					SitelinkAsset: &common.SitelinkAsset{
						LinkText:     text,
						Description1: description1,
						Description2: description2,
					},
				},
			},
		},
	})
}

func (agas *AdGroupAssets) AddCallout(text string) {
	agas.Add(&AdGroupAsset{
		AdGroupAsset: &resources.AdGroupAsset{
			FieldType: enums.AssetFieldTypeEnum_CALLOUT,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_CalloutAsset{
					CalloutAsset: &common.CalloutAsset{
						CalloutText: text,
					},
				},
			},
		},
	})
}

func (agas *AdGroupAssets) AddCall(countryCode, phoneNumber string) {
	agas.Add(&AdGroupAsset{
		AdGroupAsset: &resources.AdGroupAsset{
			FieldType: enums.AssetFieldTypeEnum_CALL,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_CallAsset{
					CallAsset: &common.CallAsset{
						CountryCode: countryCode,
						PhoneNumber: phoneNumber,
					},
				},
			},
		},
	})
}

func (agas *AdGroupAssets) AddImage(source AssetImageSource) error {
	ia := &common.ImageAsset{}
	if err := source(ia); err != nil {
		return err
	}

	agas.Add(&AdGroupAsset{
		AdGroupAsset: &resources.AdGroupAsset{
			FieldType: enums.AssetFieldTypeEnum_AD_IMAGE,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_ImageAsset{
					ImageAsset: ia,
				},
			},
		},
	})

	return nil
}

func (agas AdGroupAssets) createOperations(adGroup *AdGroup, tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(agas, func(item *AdGroupAsset) []*services.MutateOperation {
		return item.createOperations(adGroup, tempId)
	}))
}

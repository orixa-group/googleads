package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type AssetGroupAsset struct {
	*resources.AssetGroupAsset
	Asset *Asset
}

func NewAssetGroupAsset() *AssetGroupAsset {
	return &AssetGroupAsset{
		AssetGroupAsset: &resources.AssetGroupAsset{},
		Asset:           NewAsset(),
	}
}

func (aga *AssetGroupAsset) createOperations(assetGroup *AssetGroup, tempId tempIdGenerator) []*services.MutateOperation {
	aop := aga.Asset.createOperation(assetGroup.Campaign.Customer, tempId)

	aga.AssetGroupAsset.AssetGroup = assetGroup.GetResourceName()
	aga.AssetGroupAsset.Asset = aga.Asset.GetResourceName()

	return []*services.MutateOperation{
		aop,
		{
			Operation: &services.MutateOperation_AssetGroupAssetOperation{
				AssetGroupAssetOperation: &services.AssetGroupAssetOperation{
					Operation: &services.AssetGroupAssetOperation_Create{
						Create: aga.AssetGroupAsset,
					},
				},
			},
		},
	}
}

type AssetGroupAssets []*AssetGroupAsset

func NewAssetGroupAssets() AssetGroupAssets {
	return make(AssetGroupAssets, 0)
}

func (aga *AssetGroupAssets) Add(asset *AssetGroupAsset) {
	*aga = append(*aga, &AssetGroupAsset{&resources.AssetGroupAsset{
		FieldType: asset.GetFieldType(),
	}, &Asset{&resources.Asset{
		AssetData: asset.Asset.GetAssetData(),
		FinalUrls: asset.Asset.GetFinalUrls(),
	}}})
}

func (aga *AssetGroupAssets) AddSitelink(text, description1, description2 string, finalUrls ...string) {
	aga.Add(&AssetGroupAsset{
		AssetGroupAsset: &resources.AssetGroupAsset{
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

func (aga *AssetGroupAssets) AddCallout(text string) {
	aga.Add(&AssetGroupAsset{
		AssetGroupAsset: &resources.AssetGroupAsset{
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

func (aga *AssetGroupAssets) AddCall(countryCode, phoneNumber string) {
	aga.Add(&AssetGroupAsset{
		AssetGroupAsset: &resources.AssetGroupAsset{
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

func (aga *AssetGroupAssets) AddStructuredSnippet(header string, values ...string) {
	aga.Add(&AssetGroupAsset{
		AssetGroupAsset: &resources.AssetGroupAsset{
			FieldType: enums.AssetFieldTypeEnum_STRUCTURED_SNIPPET,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_StructuredSnippetAsset{
					StructuredSnippetAsset: &common.StructuredSnippetAsset{
						Header: header,
						Values: values,
					},
				},
			},
		},
	})
}

func (aga AssetGroupAssets) createOperations(assetGroup *AssetGroup, tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(aga, func(item *AssetGroupAsset) []*services.MutateOperation {
		return item.createOperations(assetGroup, tempId)
	}))
}

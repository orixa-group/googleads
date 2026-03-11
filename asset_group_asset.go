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

func (aga *AssetGroupAssets) Add(asset *AssetGroupAsset) {
	a := &resources.Asset{
		AssetData: asset.Asset.GetAssetData(),
	}

	if asset.Asset.Name != nil {
		a.Name = String(asset.Asset.GetName())
	}

	*aga = append(*aga, &AssetGroupAsset{&resources.AssetGroupAsset{
		FieldType: asset.GetFieldType(),
	}, &Asset{a}})
}

func (aga *AssetGroupAssets) AddHeadline(text string) {
	aga.addTextAsset(text, enums.AssetFieldTypeEnum_HEADLINE)
}

func (aga *AssetGroupAssets) AddLongHeadline(text string) {
	aga.addTextAsset(text, enums.AssetFieldTypeEnum_LONG_HEADLINE)
}

func (aga *AssetGroupAssets) AddDescription(text string) {
	aga.addTextAsset(text, enums.AssetFieldTypeEnum_DESCRIPTION)
}

func (aga *AssetGroupAssets) AddBusinessName(text string) {
	aga.addTextAsset(text, enums.AssetFieldTypeEnum_BUSINESS_NAME)
}

func (aga *AssetGroupAssets) AddMarketingImage(source AssetImageSource) error {
	return aga.addImageAsset(source, enums.AssetFieldTypeEnum_MARKETING_IMAGE)
}

func (aga *AssetGroupAssets) AddSquareMarketingImage(source AssetImageSource) error {
	return aga.addImageAsset(source, enums.AssetFieldTypeEnum_SQUARE_MARKETING_IMAGE)
}

func (aga *AssetGroupAssets) AddLogo(source AssetImageSource) error {
	return aga.addImageAsset(source, enums.AssetFieldTypeEnum_LOGO)
}

func (aga *AssetGroupAssets) addTextAsset(text string, fieldType enums.AssetFieldTypeEnum_AssetFieldType) {
	aga.Add(&AssetGroupAsset{
		AssetGroupAsset: &resources.AssetGroupAsset{
			FieldType: fieldType,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_TextAsset{
					TextAsset: &common.TextAsset{
						Text: String(text),
					},
				},
			},
		},
	})
}

func (aga *AssetGroupAssets) addImageAsset(source AssetImageSource, fieldType enums.AssetFieldTypeEnum_AssetFieldType) error {
	ia := &common.ImageAsset{}
	if err := source(ia); err != nil {
		return err
	}

	aga.Add(&AssetGroupAsset{
		AssetGroupAsset: &resources.AssetGroupAsset{
			FieldType: fieldType,
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

func (aga AssetGroupAssets) createOperations(assetGroup *AssetGroup, tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(aga, func(item *AssetGroupAsset) []*services.MutateOperation {
		return item.createOperations(assetGroup, tempId)
	}))
}

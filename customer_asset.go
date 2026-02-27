package googleads

import (
	"github.com/shenzhencenter/google-ads-pb/common"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

type CustomerAsset struct {
	*resources.CustomerAsset
	Asset *Asset
}

func NewCustomerAsset() *CustomerAsset {
	return &CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{},
		Asset:         NewAsset(),
	}
}

func (c *CustomerAsset) createOperations(customer *Customer, tempId tempIdGenerator) []*services.MutateOperation {
	aop := c.Asset.createOperation(customer, tempId)

	c.CustomerAsset.Asset = c.Asset.GetResourceName()

	return []*services.MutateOperation{
		aop,
		{
			Operation: &services.MutateOperation_CustomerAssetOperation{
				CustomerAssetOperation: &services.CustomerAssetOperation{
					Operation: &services.CustomerAssetOperation_Create{
						Create: c.CustomerAsset,
					},
				},
			},
		},
	}
}

type CustomerAssets []*CustomerAsset

func NewCustomerAssets() CustomerAssets {
	return make(CustomerAssets, 0)
}

func (c *CustomerAssets) Add(asset *CustomerAsset) {
	*c = append(*c, &CustomerAsset{&resources.CustomerAsset{
		FieldType: asset.GetFieldType(),
	}, &Asset{&resources.Asset{
		AssetData: asset.Asset.GetAssetData(),
		FinalUrls: asset.Asset.GetFinalUrls(),
	}}})
}

func (c *CustomerAssets) AddSitelink(text, finalUrl string, description string) {
	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
			FieldType: enums.AssetFieldTypeEnum_SITELINK,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				FinalUrls: []string{finalUrl},
				AssetData: &resources.Asset_SitelinkAsset{
					SitelinkAsset: &common.SitelinkAsset{
						LinkText:     text,
						Description1: description,
					},
				},
			},
		},
	})
}

func (c *CustomerAssets) AddCallout(text string) {
	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
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

func (c *CustomerAssets) AddStructuredSnippet(header string, values ...string) {
	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
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

func (c *CustomerAssets) AddBusinessName(name string) {
	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
			FieldType: enums.AssetFieldTypeEnum_BUSINESS_NAME,
		},
		Asset: &Asset{
			Asset: &resources.Asset{
				AssetData: &resources.Asset_TextAsset{
					TextAsset: &common.TextAsset{
						Text: String(name),
					},
				},
			},
		},
	})
}

func (c *CustomerAssets) AddBusinessLogo(source AssetImageSource) error {
	return c.addImageAsset(source, enums.AssetFieldTypeEnum_BUSINESS_LOGO)
}

func (c *CustomerAssets) AddImage(source AssetImageSource) error {
	return c.addImageAsset(source, enums.AssetFieldTypeEnum_AD_IMAGE)
}

func (c *CustomerAssets) addImageAsset(source AssetImageSource, fieldType enums.AssetFieldTypeEnum_AssetFieldType) error {
	ia := &common.ImageAsset{}
	if err := source(ia); err != nil {
		return err
	}

	c.Add(&CustomerAsset{
		CustomerAsset: &resources.CustomerAsset{
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

func (c CustomerAssets) createOperations(customer *Customer, tempId tempIdGenerator) []*services.MutateOperation {
	return Flatten(Map(c, func(item *CustomerAsset) []*services.MutateOperation {
		return item.createOperations(customer, tempId)
	}))
}
